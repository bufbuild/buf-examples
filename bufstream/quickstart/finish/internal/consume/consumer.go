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
// A Consume takes a Kafka client and a topic, and expects to receive Protobuf messages
// of the given type. Upon every received message, a handler is invoked. If malformed
// data is received (data that cannot be deserialized into the given Protobuf message type),
// a malformed data handler is invoked.
//
// This is a toy example, but shows the basics you need to receive Protobuf messages
// from Kafka using franz-go. You can likely use this as a base to build out your own demo.
type Consumer[M proto.Message] struct {
	client               *kgo.Client
	topic                string
	messageHandler       func(context.Context, M) error
	malformedDataHandler func(context.Context, []byte, error) error
}

// NewConsumer returns a new Consumer.
//
// Always use this constructor to construct Consumers.
func NewConsumer[M proto.Message](
	client *kgo.Client,
	// deserializer csr.Serde,
	topic string,
	options ...ConsumerOption[M],
) *Consumer[M] {

	consumer := &Consumer[M]{
		client: client,
		// deserializer:         deserializer,
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
func WithMessageHandler[M proto.Message](messageHandler func(context.Context, M) error) ConsumerOption[M] {
	return func(consumer *Consumer[M]) {
		consumer.messageHandler = messageHandler
	}
}

// Consume consumes as many records as it can from the topic, deserializing them into
// a message of type M if it can, and then invoking the message handler. It invokes the
// malformed data handler if the record's payload cannot be deserialized into type M.
func (c *Consumer[M]) Consume(ctx context.Context) error {
	fetches := c.client.PollFetches(ctx)
	if errs := fetches.Errors(); len(errs) > 0 {
		return fmt.Errorf("failed to fetch records: %v", errs)
	}
	for _, record := range fetches.Records() {
		var message M
		// Get the concrete type of M
		msgType := reflect.TypeOf(message).Elem()
		// Create a new instance of the concrete type
		message = reflect.New(msgType).Interface().(M)
		// Unmarshal bytes onto it
		if err := proto.Unmarshal(record.Value, message); err != nil {
			return err
		}

		if err := c.messageHandler(ctx, message); err != nil {
			return err
		}
	}
	return nil
}

func defaultMessageHandler[M proto.Message](ctx context.Context, message M) error {
	slog.InfoContext(ctx, "consumed message", "message", message)
	return nil
}

func defaultMalformedDataHandler(ctx context.Context, payload []byte, err error) error {
	slog.InfoContext(ctx, "consumed malformed data", "error", err, "length", len(payload))
	return nil
}
