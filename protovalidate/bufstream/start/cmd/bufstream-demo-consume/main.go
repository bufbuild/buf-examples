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

// Package main implements the consumer of the demo.
//
// This is run as part of docker compose.
//
// The consumer will read as many records it can at once, print what it received,
// sleep for one second, and then loop.
package main

import (
	"context"
	"fmt"
	demov1 "github.com/bufbuild/buf-examples/protovalidate/bufstream-data-quality/start/gen/bufstream/demo/v1"
	"log/slog"
	"time"

	"github.com/bufbuild/buf-examples/protovalidate/bufstream-data-quality/start/pkg/app"
	"github.com/bufbuild/buf-examples/protovalidate/bufstream-data-quality/start/pkg/consume"
	"github.com/bufbuild/buf-examples/protovalidate/bufstream-data-quality/start/pkg/csr"
	"github.com/bufbuild/buf-examples/protovalidate/bufstream-data-quality/start/pkg/kafka"
)

func main() {
	// See the app package for the boilerplate we use to set up the producer and
	// consumer, including bound flags.
	app.Main(run)
}

func run(ctx context.Context, config app.Config) error {
	client, err := kafka.NewKafkaClient(config.Kafka, true)
	if err != nil {
		return err
	}
	defer client.Close()

	// NewDeserializer creates a CSR-based Deserializer if there is a CSR URL,
	// otherwise it creates a single-type Deserializer for demov1.EmailUpdated.
	deserializer, err := csr.NewDeserializer[*demov1.EmailUpdated](config.CSR)
	if err != nil {
		return err
	}
	defer func() { _ = deserializer.Close() }()

	consumer := consume.NewConsumer(
		client,
		deserializer,
		config.Kafka.Topic,
		consume.WithMessageHandler(handleEmailUpdated),
	)

	slog.Info("starting consume")
	for {
		// Read as many messages as we can.
		//
		// Only return error if there is an unexpected system error. Of note, an error is not
		// returned if the data that the consumer receives is malformed.
		if err := consumer.Consume(ctx); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
}

func handleEmailUpdated(message *demov1.EmailUpdated) error {
	var suffix string
	if old := message.GetOldEmailAddress(); old == "" {
		suffix = "redacted old email"
	} else {
		suffix = fmt.Sprintf("old email %s", old)
	}
	slog.Info(fmt.Sprintf("consumed message with new email %s and %s", message.GetNewEmailAddress(), suffix))
	return nil
}
