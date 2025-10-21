package kafka

import (
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

// Config is all configuration we need to build a new Kafka Client.
//
// franz-go uses functional options for the same purpose, but we're simplifying this
// to just the values in this config struct for the purposes of this demo. If you use
// franz-go in production code, we'd recommend using the functional options directly.
type Config struct {
	// BootstrapServers are the bootstrap servers to call.
	BootstrapServers []string
	ClientID         string
	Topic            string
	DlqTopic         string
	Group            string
}

// NewKafkaClient returns a new franz-go Kafka Client for the given Config.
func NewKafkaClient(config Config, consumer bool) (*kgo.Client, error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(config.BootstrapServers...),
		kgo.ClientID(config.ClientID),
	}

	if consumer {
		opts = append(opts,
			kgo.ConsumerGroup(config.Group),
			kgo.ConsumeTopics(config.Topic),
			kgo.FetchMaxWait(time.Second),
			kgo.FetchIsolationLevel(kgo.ReadCommitted()),
			kgo.RequireStableFetchOffsets(),
		)
	}

	return kgo.NewClient(opts...)
}
