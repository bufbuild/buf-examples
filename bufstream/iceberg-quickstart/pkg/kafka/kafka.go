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
	"context"
	"errors"
	"fmt"
	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kerr"
	"github.com/twmb/franz-go/pkg/kgo"
	"log/slog"
)

// Config describes all Kafka configuration needed to run this example.
type Config struct {
	// BootstrapServers are the bootstrap servers to call.
	BootstrapServers []string
	ClientID         string
	Topic            string
	ArchiveKind      string
	Catalog          string
	Table            string
	TopicPartitions  int32
}

// NewKafkaClient returns a new franz-go Kafka Client for the given Config.
func NewKafkaClient(config Config) (*kgo.Client, error) {
	opts := []kgo.Opt{
		kgo.SeedBrokers(config.BootstrapServers...),
		kgo.ClientID(config.ClientID),
	}
	return kgo.NewClient(opts...)
}

// NewAdminClient returns an franz-go admin client for the given Config.
func NewAdminClient(config Config) (*kadm.Client, error) {
	client, err := NewKafkaClient(config)
	if err != nil {
		return nil, err
	}

	return kadm.NewClient(client), nil
}

// TopicExists uses an admin client and a Config to determine if the Config's
// Topic exists.
func TopicExists(ctx context.Context, client *kadm.Client, config Config) (bool, error) {
	topicConfig, err := topicConfig(ctx, client, config)
	if err != nil {
		return false, err
	}
	return len(topicConfig.Name) > 0, nil
}

// CreateTopic uses an admin client to create the Topic requested in the
// provided Config, using the Config.TopicConfig and Config.TopicPartitions. If
// the topic already exists, an error is returned.
func CreateTopic(ctx context.Context, client *kadm.Client, config Config) error {
	slog.Info("Creating topic", "topic", config.Topic)

	// Create the topic.
	_, err := client.CreateTopic(ctx, config.TopicPartitions, 1, nil, config.Topic)
	if err != nil {
		return err
	}
	slog.Info("Created topic", "topic", config.Topic)

	return nil
}

// ConfigureTopic sets topic parameters.
func ConfigureTopic(ctx context.Context, client *kadm.Client, config Config) error {
	alterConfigs := make([]kadm.AlterConfig, 3)
	alterConfigs[0] = kadm.AlterConfig{
		Name:  "bufstream.archive.kind",
		Value: &config.ArchiveKind,
	}
	alterConfigs[1] = kadm.AlterConfig{
		Name:  "bufstream.archive.iceberg.catalog",
		Value: &config.Catalog,
	}
	alterConfigs[2] = kadm.AlterConfig{
		Name:  "bufstream.archive.iceberg.table",
		Value: &config.Table,
	}
	_, err := client.AlterTopicConfigs(ctx, alterConfigs, config.Topic)

	return err
}

// VerifyTopicConfig verifies that a given topic exists and has the
// correct parameters set for Iceberg integration.
func VerifyTopicConfig(ctx context.Context, client *kadm.Client, config Config) error {
	topicConfig, err := topicConfig(ctx, client, config)
	if err != nil {
		return err
	}
	if len(topicConfig.Name) == 0 {
		return fmt.Errorf("kafka topic %s does not exist, see this example's documentation for how to use the `configure` command to create it", config.Topic)
	}

	configSpecs := map[string]func(value string) error{
		"bufstream.archive.iceberg.table": func(value string) error {
			if len(value) == 0 {
				return fmt.Errorf("bufstream.archive.iceberg.table must be set for topic %s, run 'configure' to configure this topic for Iceberg", config.Topic)
			}
			return nil
		},
		"bufstream.archive.iceberg.catalog": func(value string) error {
			if len(value) == 0 {
				return fmt.Errorf("bufstream.archive.iceberg.catalog must be set for topic %s, run 'configure' to configure this topic for Iceberg", config.Topic)
			}
			return nil
		},
		"bufstream.archive.kind": func(value string) error {
			if value != "ICEBERG" {
				return fmt.Errorf("the value of bufstream.archive.kind must be `ICEBERG`, not %s", value)
			}
			return nil
		},
	}

	for name, expect := range configSpecs {
		parameterFound := false
		for _, r := range topicConfig.Configs {
			if r.Key == name {
				if err := expect(*r.Value); err != nil {
					return err
				}
				parameterFound = true
			}
		}
		if !parameterFound {
			return fmt.Errorf("topic parameter %s wasn't set for topic %s, run 'configure' to configure this topic for Iceberg", name, config.Topic)
		}
	}
	return nil
}

func topicConfig(ctx context.Context, client *kadm.Client, config Config) (kadm.ResourceConfig, error) {
	// Check to see if the topic exists.
	resp, err := client.DescribeTopicConfigs(ctx, config.Topic)
	if err == nil {
		if len(resp) != 1 {
			return kadm.ResourceConfig{}, fmt.Errorf("expected 1 topic config, got %d", len(resp))
		}
		err = resp[0].Err
	}
	// The topic exists.
	if err == nil {
		return resp[0], nil
	}
	// There's an error other than the expected kerr.UnknownTopicOrPartition.
	if !isUnknownTopic(err) {
		return kadm.ResourceConfig{}, err // something went wrong
	}
	// The topic doesn't exist.
	return kadm.ResourceConfig{}, nil
}

func isUnknownTopic(err error) bool {
	var kError *kerr.Error
	return errors.As(err, &kError) && kError.Code == kerr.UnknownTopicOrPartition.Code
}
