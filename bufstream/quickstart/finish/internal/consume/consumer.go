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

// Consumer is an example consumer of a given topic using a given Protobuf message type.
//
// A Consumer takes a Kafka client and a topic and expects to receive Protobuf messages
// of the given type. Upon every received message, a handler is invoked.
//
// This is a toy example, but shows the basics you need to receive Protobuf messages
// from Kafka using franz-go.
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

// Consume consumes as many records as it can from the topic, deserializing them into
// a message of type M if it can, and then invoking the message handler.
func (c *Consumer[M]) Consume(ctx context.Context) error {
	fetches := c.client.PollFetches(ctx)
	if errs := fetches.Errors(); len(errs) > 0 {
		return fmt.Errorf("failed to fetch records: %v", errs)
	}
	for _, record := range fetches.Records() {
		message, err := c.toMessage(record)
		if err != nil {
			return err
		}
		if err := c.messageHandler(ctx, message); err != nil {
			return err
		}
	}
	return nil
}

func (c *Consumer[M]) toMessage(record *kgo.Record) (M, error) {
	var message M
	msgType := reflect.TypeOf(message).Elem()
	message = reflect.New(msgType).Interface().(M)
	err := proto.Unmarshal(record.Value, message)
	return message, err
}

func defaultMessageHandler[M proto.Message](ctx context.Context, message M) error {
	slog.InfoContext(ctx, "consumed message", "message", message)
	return nil
}
