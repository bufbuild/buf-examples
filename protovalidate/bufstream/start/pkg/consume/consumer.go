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

// Package consume implements a toy consumer.
package consume

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
)

// Consumer is an example consumer of a given topic using given Protobuf message type.
//
// A Consume takes a Kafka client and a topic, and expects to recieve Protobuf messages
// of the given type. Upon every received message, a handler is invoked. If malformed
// data is recieved (data that can not be deserialized into the given Protobuf message type),
// a malformed data handler is invoked.
//
// This is a toy example, but shows the basics you need to receive Protobuf messages
// from Kafka using franz-go. You can likely use this as a base to build out your own demo.
type Consumer[M proto.Message] struct {
	client               *kgo.Client
	deserializer         serde.Deserializer
	topic                string
	messageHandler       func(M) error
	malformedDataHandler func([]byte, error) error
}

// NewConsumer returns a new Consumer.
//
// Always use this constructor to construct Consumers.
func NewConsumer[M proto.Message](
	client *kgo.Client,
	deserializer serde.Deserializer,
	topic string,
	options ...ConsumerOption[M],
) *Consumer[M] {
	consumer := &Consumer[M]{
		client:               client,
		deserializer:         deserializer,
		topic:                topic,
		messageHandler:       defaultMessageHandler[M],
		malformedDataHandler: defaultMalformedDataHandler,
	}
	for _, option := range options {
		option(consumer)
	}
	return consumer
}

// ConsumerOption is an option when constructing a new Consumer.
//
// All parameters except options are required. ConsumerOptions allow
// for optional parameters.
type ConsumerOption[M proto.Message] func(*Consumer[M])

// WithMessageHandler returns a new ConsumerOption that overrides the default
// handler of received messages.
//
// The default handler uses slog to log incoming messages.
func WithMessageHandler[M proto.Message](messageHandler func(M) error) ConsumerOption[M] {
	return func(consumer *Consumer[M]) {
		consumer.messageHandler = messageHandler
	}
}

// WithMalformedDataHandler returns a new Consumer Option that overrides the default
// handler of malformed received data.
//
// The default handler uses slog to log the error returned on malformed data.
func WithMalformedDataHandler[M proto.Message](
	malformedDataHandler func([]byte, error) error,
) ConsumerOption[M] {
	return func(consumer *Consumer[M]) {
		consumer.malformedDataHandler = malformedDataHandler
	}
}

// Consume consumes as many record as it can from the topic, deserializing them into
// a message of type M if it can, and then invoking the message handler. It invokes the
// malformed data handler if the record's payload cannot be deserialized into type M.
func (c *Consumer[M]) Consume(ctx context.Context) error {
	fetches := c.client.PollFetches(ctx)
	if errs := fetches.Errors(); len(errs) > 0 {
		return fmt.Errorf("failed to fetch records: %v", errs)
	}
	for _, record := range fetches.Records() {
		data, err := c.deserializer.Deserialize(record.Topic, record.Value)
		if err != nil {
			if err := c.malformedDataHandler(record.Value, err); err != nil {
				return err
			}
			continue
		}
		message, ok := data.(M)
		if !ok {
			if err := c.malformedDataHandler(
				record.Value,
				fmt.Errorf("received unexpected message type: %T", data),
			); err != nil {
				return err
			}
		}
		if err := c.messageHandler(message); err != nil {
			return err
		}
	}
	return nil
}

func defaultMessageHandler[M proto.Message](message M) error {
	slog.Info("consumed message", "message", message)
	return nil
}

func defaultMalformedDataHandler(payload []byte, err error) error {
	slog.Info("consumed malformed data", "error", err)
	return nil
}
