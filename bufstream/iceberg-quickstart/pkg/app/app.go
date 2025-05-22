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

// Package app implements boilerplate code used by the producer.
//
// It implements Main, which the producer uses within its main function.
// It also binds all relevant flags.
package app

import (
	"context"
	"github.com/bufbuild/buf-examples/bufstream/iceberg-quickstart/pkg/csr"
	"github.com/bufbuild/buf-examples/bufstream/iceberg-quickstart/pkg/kafka"
	"github.com/spf13/pflag"
	"log/slog"
	"os"
	"os/signal"
)

const (
	defaultKafkaClientID = "bufstream-iceberg-quickstart"
)

var (
	defaultKafkaBootstrapServers = []string{"localhost:9092"}
)

// Config contains all application configuration needed by the producer.
type Config struct {
	Kafka kafka.Config
	CSR   csr.Config
}

// Main reads any configuration and does minimal setup for any binary.
//
// Note that in a real production workload, applications should not create
// topics. This should be considered an infrastructure concern, and the topics
// should be provisioned with correct configuration before a producer ever tries
// to send messages to it. If the topic does not exist, this should be a failure
// in the producer since it means a likely misconfiguration.
//
// This example creates the topic, despite it not being a typical good practice,
// just for simplicity, so there are fewer steps to get the example running.
func Main(action func(context.Context, Config) error) {
	// Set up slog. We use the global logger throughout this demo.
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))
	// Cancel the context on interrupt, i.e. ctrl+c for our purposes.
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := run(ctx, action); err != nil {
		slog.Error("program error", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, do func(context.Context, Config) error) error {
	config, err := parseConfig()
	if err != nil {
		return err
	}
	return do(ctx, config)
}

func parseConfig() (Config, error) {
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
