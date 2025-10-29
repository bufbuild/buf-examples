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
	"reflect"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
)

// Consumer is an example consumer of a topic using a Protobuf message.
//
// A Consumer takes a Kafka client and a topic and expects to receive Protobuf
// messages of the given type. Upon every received message, a handler is
// invoked.
//
// This is a toy example, but shows the basics you need to receive Protobuf
// messages from Kafka using franz-go.
type Consumer[M proto.Message] struct {
	client         *kgo.Client
	topic          string
	messageHandler func(context.Context, M) error
}

// NewConsumer returns a new Consumer.
//
// Always use this constructor to construct Consumers.
func NewConsumer[M proto.Message](
	client *kgo.Client,
	topic string,
	options ...ConsumerOption[M],
) *Consumer[M] {
	consumer := &Consumer[M]{
		client:         client,
		topic:          topic,
		messageHandler: defaultMessageHandler[M],
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
func WithMessageHandler[M proto.Message](messageHandler func(context.Context, M) error) ConsumerOption[M] {
	return func(consumer *Consumer[M]) {
		consumer.messageHandler = messageHandler
	}
}

// Consume consumes as many records as it can from the topic, deserializing them
// into a message of type M if it can, and then passing each to the message
// handler.
func (c *Consumer[M]) Consume(ctx context.Context) error {
	fetches := c.client.PollFetches(ctx)
	if errs := fetches.Errors(); len(errs) > 0 {
		return fmt.Errorf("failed to fetch records: %v", errs)
	}
	for _, record := range fetches.Records() {
		message, err := c.toMessage(record)
		if err != nil {
			return err
		} else {
			if err := c.messageHandler(ctx, message); err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *Consumer[M]) toMessage(record *kgo.Record) (M, error) {
	var message M
	msgType := reflect.TypeOf(message).Elem()
	message = reflect.New(msgType).Interface().(M)
	err := proto.Unmarshal(record.Value, message)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal record value onto %s: %w", msgType.Name(), err)
	}
	return message, err
}

func defaultMessageHandler[M proto.Message](ctx context.Context, message M) error {
	slog.InfoContext(ctx, "consumed message", "message", message)
	return nil
}
