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
// The consumer will read as many records it can at once, print what it received,
package main

import (
	"context"
	"errors"
	"log/slog"

	shoppingv1 "github.com/bufbuild/buf-examples/bufstream/quickstart/gen/shopping/v1"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/app"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/consume"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/kafka"
)

func main() {
	// See the app package for the boilerplate we use to set up the producer and
	// consumer, including bound flags.
	app.Main(run)
}

func run(ctx context.Context, config app.Config) error {
	// consume-specific config validation.
	if config.Kafka.Group == "" {
		return errors.New("must specify --group")
	}

	client, err := kafka.NewKafkaClient(config.Kafka, true)
	if err != nil {
		return err
	}
	defer client.Close()

	consumer := consume.NewConsumer(
		client,
		config.Kafka.Topic,
		consume.WithMessageHandler(handleCart),
	)

	slog.InfoContext(ctx, "starting consume")
	for {
		// Read as many messages as we can.
		//
		// Only return error if there is an unexpected system error. Of note, an error is not
		// returned if the data that the consumer receives is malformed.
		if err := consumer.Consume(ctx); err != nil {
			return err
		}
	}
}

func handleCart(_ context.Context, invoice *shoppingv1.Cart) error {
	for _, lineItem := range invoice.GetLineItems() {
		if lineItem.GetQuantity() == 0 {
			slog.Error("received a Cart with a zero-quantity LineItem")
		}
	}
	return nil
}
