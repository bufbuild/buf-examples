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
	"log/slog"
	"strings"

	"github.com/twmb/franz-go/pkg/kadm"
	"github.com/twmb/franz-go/pkg/kerr"
)

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
	// Check to see if the topic exists.
	resp, err := client.DescribeTopicConfigs(ctx, config.Topic)
	if err == nil {
		if len(resp) != 1 {
			return false, fmt.Errorf("expected 1 topic config, got %d", len(resp))
		}
		err = resp[0].Err
	}
	// The topic exists.
	if err == nil {
		return true, nil
	}
	// There's an error other than the expected kerr.UnknownTopicOrPartition.
	if !isUnknownTopic(err) {
		return false, err // something went wrong
	}
	// The topic doesn't exist.
	return false, nil
}

// CreateTopic uses an admin client to create the Topic requested in the
// provided Config, using the Config.TopicConfig and Config.TopicPartitions. If
// the topic already exists, an error is returned.
func CreateTopic(ctx context.Context, client *kadm.Client, config Config) error {
	slog.Info("Creating topic", "topic", config.Topic)
	configs := make(map[string]*string, len(config.TopicConfig))

	// Reduce config.TopicConfig to a map[string]*string for the admin client.
	for _, conf := range config.TopicConfig {
		param, value, _ := strings.Cut(conf, "=")
		if value == "" {
			configs[param] = nil
		} else {
			configs[param] = &value
		}
		slog.Info("Configuring topic", "topic", config.Topic, "parameter", param, "value", value)
	}

	// Create the topic.
	_, err := client.CreateTopic(ctx, config.TopicPartitions, 1, configs, config.Topic)
	if err != nil {
		return err
	}
	slog.Info("Created topic", "topic", config.Topic)

	return nil
}

func isUnknownTopic(err error) bool {
	var kError *kerr.Error
	return errors.As(err, &kError) && kError.Code == kerr.UnknownTopicOrPartition.Code
}
