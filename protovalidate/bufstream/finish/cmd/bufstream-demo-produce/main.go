// Copyright 2025 Buf Technologies, Inc.
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

// Package main implements the producer of the demo.

// This is run as part of docker compose.
//
// The producer will produce three records at a time:
//
//   - A semantically-valid EmailUpdated message, where both email fields are valid email addresses.
//   - A semantically-invalid EmailUpdated message, where the new email field is not a valid email address.
//   - A record containing a payload that is not valid Protobuf.
//
// After producing these three records, the producer sleeps for one second, and then loops.
package main

import (
	"context"
	"errors"
	demov1 "github.com/bufbuild/buf-examples/protovalidate/bufstream-data-quality/finish/gen/bufstream/demo/v1"
	"log/slog"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/bufbuild/buf-examples/protovalidate/bufstream-data-quality/finish/pkg/app"
	"github.com/bufbuild/buf-examples/protovalidate/bufstream-data-quality/finish/pkg/csr"
	"github.com/bufbuild/buf-examples/protovalidate/bufstream-data-quality/finish/pkg/kafka"
	"github.com/bufbuild/buf-examples/protovalidate/bufstream-data-quality/finish/pkg/produce"
	"github.com/google/uuid"
)

func main() {
	// See the app package for the boilerplate we use to set up the producer and
	// consumer, including bound flags.
	app.Main(run)
}

func run(ctx context.Context, config app.Config) error {
	client, err := kafka.NewKafkaClient(config.Kafka, false)
	if err != nil {
		return err
	}
	defer client.Close()

	// NewSerializer creates a CSR-based Serializer if there is a CSR URL,
	// otherwise it creates a single-type Serializer for demov1.EmailUpdated.
	serializer, err := csr.NewSerializer[*demov1.EmailUpdated](config.CSR)
	if err != nil {
		return err
	}
	defer func() { _ = serializer.Close() }()

	producer := produce.NewProducer[*demov1.EmailUpdated](
		client,
		serializer,
		config.Kafka.Topic,
	)

	slog.Info("starting produce")
	for {
		id := newID()
		// Produces semantically-valid EmailUpdated message, where both email
		// fields are valid email addresses.
		if err := producer.ProduceProtobufMessage(ctx, id, newSemanticallyValidEmailUpdated(id)); err != nil {
			if errors.Is(err, context.Canceled) {
				return err
			}
			slog.Error("error on produce or semantically valid protobuf message", "error", err)
		} else {
			slog.Info("produced semantically valid protobuf message", "id", id)
		}
		id = newID()
		// Produces a semantically-invalid EmailUpdated message, where the new email field
		// is not a valid email address.
		if err := producer.ProduceProtobufMessage(ctx, id, newSemanticallyInvalidEmailUpdated(id)); err != nil {
			if errors.Is(err, context.Canceled) {
				return err
			}
			slog.Error("error on produce of semantically invalid protobuf message", "error", err)
		} else {
			slog.Info("produced semantically invalid protobuf message", "id", id)
		}
		id = newID()
		time.Sleep(time.Second)
	}
}

// newID returns a new UUID.
//
// This is also used as the record key.
func newID() string {
	return uuid.New().String()
}

func newSemanticallyValidEmailUpdated(id string) *demov1.EmailUpdated {
	return &demov1.EmailUpdated{
		Id:              id,
		OldEmailAddress: gofakeit.Email(),
		NewEmailAddress: gofakeit.Email(),
	}
}

func newSemanticallyInvalidEmailUpdated(id string) *demov1.EmailUpdated {
	return &demov1.EmailUpdated{
		Id:              id,
		OldEmailAddress: gofakeit.Email(),
		NewEmailAddress: "bad-email-" + gofakeit.Animal(),
	}
}
