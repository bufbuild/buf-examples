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

// Package main implements the dead-letter queue consumer of the quickstart.
// It will read as many records it can at once, printing why records were
// rejected. If it can't explain a rejection, it will log errors.
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	dlqv1beta1 "buf.build/gen/go/bufbuild/bufstream/protocolbuffers/go/buf/bufstream/dlq/v1beta1"
	"buf.build/go/protovalidate"
	shoppingv1 "github.com/bufbuild/buf-examples/bufstream/quickstart/gen/shopping/v1"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/app"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/consume"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/kafka"
	"google.golang.org/protobuf/proto"
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

	consumer := consume.NewConsumer(
		client,
		config.Kafka.Topic,
		consume.WithMessageHandler(handleDlqRecord),
	)

	slog.InfoContext(ctx, "consuming DLQ topic", "topic", config.Kafka.Topic)
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

func handleDlqRecord(_ context.Context, record *dlqv1beta1.Record) error {
	// Reconstruct the original message: we expect a Cart in this toy example.
	cart := &shoppingv1.Cart{}
	if err := proto.Unmarshal(record.GetValue(), cart); err != nil {
		return fmt.Errorf("failed to unmarshal dlq value into shoppingv1.Cart: %w", err)
	}

	// Try to use Protovalidate to determine what was wrong with the cart!
	if err := protovalidate.Validate(cart); err != nil {
		slog.Info("DLQ received a cart that failed due to validation errors:", "ID", cart.GetCartId(), "error", err)
		return nil
	}

	// If we can't explain why the Cart failed, that's an error.
	return errors.New("DLQ received a cart for an unknown reason")
}
