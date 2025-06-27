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
)

const (
	defaultKafkaClientID        = "bufstream-iceberg-quickstart"
	defaultKafkaBootstrapServer = "localhost:9092"
)

// Config contains all application configuration needed by the producer.
type Config struct {
	Kafka kafka.Config
	CSR   csr.Config
}

// Main reads any configuration and does minimal setup for any binary.
func Main(action func(context.Context, Config) error) {
	// Set up slog. We use the global logger throughout this example.
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))
	if err := run(context.Background(), action); err != nil {
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
		[]string{defaultKafkaBootstrapServer},
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
	flagSet.Int32Var(
		&config.Kafka.TopicPartitions,
		"topic-partitions",
		1,
		"The number of partitions to use when creating the topic.",
	)
	flagSet.StringVar(
		&config.Kafka.ArchiveKind,
		"bufstream.archive.kind",
		"",
		"The type of archival to use. Its value should be ICEBERG for this example.",
	)
	flagSet.StringVar(
		&config.Kafka.Catalog,
		"bufstream.archive.iceberg.catalog",
		"",
		"The name of the Iceberg catalog configured in bufstream.yaml.",
	)
	flagSet.StringVar(
		&config.Kafka.Table,
		"bufstream.archive.iceberg.table",
		"",
		"The namespace and table name of an Iceberg table to maintain.",
	)
	flagSet.StringVar(
		&config.CSR.URL,
		"csr-url",
		"",
		"The Confluent Schema Registry URL.",
	)
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return Config{}, err
	}
	return config, nil
}
