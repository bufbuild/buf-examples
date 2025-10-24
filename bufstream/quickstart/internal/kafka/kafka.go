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

package kafka

import (
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

// Config is all configuration we need to build a new Kafka Client.
//
// franz-go uses functional options for the same purpose, but we're simplifying
// this to just the values in this config struct for the purposes of this demo.
// If you use franz-go in production code, we recommend using the functional
// options directly.
type Config struct {
	BootstrapServers []string
	ClientID         string
	Topic            string
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
