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
// It implements Main, which both the producer and consumer use within their
// main functions. It also binds all relevant flags.
package app

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"

	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/kafka"
	"github.com/spf13/pflag"
)

const defaultKafkaClientID = "bufstream-quickstart"

var (
	defaultKafkaBootstrapServers = []string{"localhost:9092"}
)

// Config contains all application configuration needed by the producer and
// consumer.
type Config struct {
	Kafka kafka.Config
}

// Main is used by the producer and consumers within their main functions.
//
// It sets up logging, interrupt handling, and binds and parses all flags.
// Afterwards, it calls action to invoke the application logic.
func Main(action func(context.Context, Config) error) {
	// Set up slog. We use the global logger throughout this demo.
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))
	// Cancel the context on interrupt, i.e. ctrl+c for our purposes.
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := run(ctx, action); err != nil {
		slog.ErrorContext(ctx, "program error", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, action func(context.Context, Config) error) error {
	config, err := parseConfig()
	if err != nil {
		return err
	}
	return action(ctx, config)
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
	flagSet.StringVar(
		&config.Kafka.Group,
		"group",
		"",
		"The Kafka consumer group ID.",
	)

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return Config{}, err
	}
	if config.Kafka.Topic == "" {
		return Config{}, errors.New("must specify --topic")
	}
	return config, nil
}
