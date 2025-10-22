// Package produce implements a toy producer.
package produce

import (
	"context"
	"fmt"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
)

// Producer is an example producer to a given topic using a Protobuf message
// type.
//
// This is a toy example, but shows the basics you need to send Protobuf
// messages to Kafka using franz-go.
type Producer[M proto.Message] struct {
	client *kgo.Client
	topic  string
}

// NewProducer returns a new Producer.
//
// Always use this constructor to construct Producers.
func NewProducer[M proto.Message](
	client *kgo.Client,
	topic string,
) *Producer[M] {
	return &Producer[M]{
		client: client,
		topic:  topic,
	}
}

// ProduceMessage serializes a Protobuf message and synchronously sends it to
// the Producer's topic with the given key.
func (p *Producer[M]) ProduceMessage(ctx context.Context, key string, message M) error {
	payload, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	return p.produce(ctx, key, payload)
}

func (p *Producer[M]) produce(ctx context.Context, key string, payload []byte) error {
	produceResults := p.client.ProduceSync(
		ctx,
		&kgo.Record{
			Key:   []byte(key),
			Value: payload,
			Topic: p.topic,
		},
	)
	if err := produceResults.FirstErr(); err != nil {
		return fmt.Errorf("failed to produce: %w", err)
	}
	return nil
}
