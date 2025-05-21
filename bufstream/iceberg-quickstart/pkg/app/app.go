// Copyright 2020-2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package app implements boilerplate code shared by the producer and consumer.
//
// It implements Main, which both the producer and consumer use within their main functions.
// It also binds all relevant flags.
package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"

	"github.com/bufbuild/buf-examples/bufstream/iceberg-quickstart/pkg/csr"
	"github.com/bufbuild/buf-examples/bufstream/iceberg-quickstart/pkg/kafka"
	"github.com/spf13/pflag"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kerr"
)

const (
	defaultKafkaClientID = "bufstream-demo"
)

var (
	defaultKafkaBootstrapServers = []string{"localhost:9092"}
)

// Config contains all application configuration needed by the producer and consumer.
type Config struct {
	Kafka kafka.Config
	CSR   csr.Config
}

// Main is used by the consumer's main function.
//
// It sets up logging, interrupt handling, and binds and parses all flags. Afterwards, it calls
// action to invoke the application logic.
func Main(action func(context.Context, Config) error) {
	doMain(false, action)
}

// MainAutoCreateTopic is used by the producer's main function. It is just like [Main] except
// that it will also create the topic if necessary. The producer defines the topic and provides
// the data, so the consumer should not be the one auto-creating it.
//
// Note that in a real production workload, neither producer nor consumer applications should
// ever create topics. This should be considered an infrastructure concern, and the topic
// should be provisioned with correct configuration before a producer ever tries to send
// messages to it. If the topic does not exist, this should be a failure in the producer
// since it means a likely misconfiguration.
//
// This demo workload creates the topic, despite it not being a typical good practice, just
// for simplicity, so there are fewer steps to get the demo running.
func MainAutoCreateTopic(action func(context.Context, Config) error) {
	doMain(true, action)
}

func doMain(autoCreateTopic bool, action func(context.Context, Config) error) { // Set up slog. We use the global logger throughout this demo.
	// Set up slog. We use the global logger throughout this demo.
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))
	// Cancel the context on interrupt, i.e. ctrl+c for our purposes.
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := run(ctx, autoCreateTopic, action); err != nil {
		slog.Error("program error", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, autoCreateTopic bool, action func(context.Context, Config) error) error {
	config, err := parseConfig(autoCreateTopic)
	if err != nil {
		return err
	}
	if autoCreateTopic {
		if err := maybeCreateTopic(ctx, config.Kafka); err != nil {
			return err
		}
	}
	return action(ctx, config)
}

func parseConfig(canCreateTopic bool) (Config, error) {
	flagSet := pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	config := Config{}
	flagSet.StringArrayVar(
		&config.Kafka.BootstrapServers,
		"bootstrap",
		defaultKafkaBootstrapServers,
		"The Bufstream bootstrap server addresses.",
	)
	flagSet.StringVar(
		&config.Kafka.ClientID,
		"client-id",
		defaultKafkaClientID,
		"The Kafka client ID.",
	)
	flagSet.StringVar(
		&config.Kafka.Topic,
		"topic",
		"",
		"The Kafka topic name to use.",
	)
	if canCreateTopic {
		flagSet.BoolVar(
			&config.Kafka.RecreateTopic,
			"recreate-topic",
			false,
			"If true, the topic will be recreated even if it already exists.",
		)
		flagSet.IntVar(
			&config.Kafka.TopicPartitions,
			"topic-partitions",
			1,
			"The number of partitions to use when creating the topic.",
		)
		flagSet.StringSliceVar(
			&config.Kafka.TopicConfig,
			"topic-config",
			nil,
			"Topic config parameters to use when creating the topic.",
		)
	}
	flagSet.StringVar(
		&config.Kafka.Group,
		"group",
		"",
		"The Kafka consumer group ID.",
	)
	flagSet.StringVar(
		&config.CSR.URL,
		"csr-url",
		"",
		"The Confluent Schema Registry URL.",
	)
	flagSet.StringVar(
		&config.CSR.Username,
		"csr-user",
		"",
		"The Confluent Schema Registry username, if authentication is needed.",
	)
	flagSet.StringVar(
		&config.CSR.Password,
		"csr-pass",
		"",
		"The Confluent Schema Registry password/token, if authentication is needed.",
	)
	flagSet.StringVar(
		&config.Kafka.RootCAPath,
		"tls-root-ca-path",
		"",
		"A path to root CA certificate for kafka TLS.",
	)
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return Config{}, err
	}
	return config, nil
}

func maybeCreateTopic(ctx context.Context, config kafka.Config) error {
	client, err := kafka.NewKafkaClient(config, false)
	if err != nil {
		return err
	}
	defer client.Close()

	admClient := kadm.NewClient(client)
	if config.RecreateTopic {
		resp, err := admClient.DeleteTopic(ctx, config.Topic)
		if err == nil {
			err = resp.Err
		}
		if !isUnknownTopic(err) {
			return err // something went wrong
		}
	} else {
		resp, err := admClient.DescribeTopicConfigs(ctx, config.Topic)
		if err == nil {
			if len(resp) != 1 {
				return fmt.Errorf("expected 1 topic config, got %d", len(resp))
			}
			err = resp[0].Err
		}
		if err == nil {
			return nil // topic exists; nothing to create
		}
		if !isUnknownTopic(err) {
			return err // something went wrong
		}
		// Else, topic does not exist, so we fall through to create it.
	}
	slog.Info("Creating topic", "name", config.Topic)
	configs := make(map[string]*string, len(config.TopicConfig))
	for _, conf := range config.TopicConfig {
		k, v, _ := strings.Cut(conf, "=")
		if v == "" {
			configs[k] = nil
		} else {
			configs[k] = &v
		}
		slog.Info("Configuring topic", "topic", config.Topic, "parameter", k, "value", v)
	}
	resp, err := admClient.CreateTopic(ctx, int32(config.TopicPartitions), 1, configs, config.Topic)
	if err == nil {
		err = resp.Err
	}
	return err
}

func isUnknownTopic(err error) bool {
	var kError *kerr.Error
	return errors.As(err, &kError) && kError.Code == kerr.UnknownTopicOrPartition.Code
}
