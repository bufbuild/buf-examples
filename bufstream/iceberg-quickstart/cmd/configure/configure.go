package main

import (
	"context"
	"fmt"
	"github.com/bufbuild/buf-examples/bufstream/iceberg-quickstart/pkg/app"
	"github.com/bufbuild/buf-examples/bufstream/iceberg-quickstart/pkg/kafka"
	"github.com/twmb/franz-go/pkg/kadm"
	"log/slog"
)

func main() {
	// See the app package for the boilerplate we use to set up the topic,
	// including bound flags.
	app.Main(run)
}

func run(ctx context.Context, config app.Config) error {
	// Create an admin client.
	client, err := kafka.NewAdminClient(config.Kafka)
	if err != nil {
		return err
	}
	defer client.Close()

	// Ensure the topic exists, creating it if necessary.
	if err := ensureTopicExists(ctx, client, config.Kafka); err != nil {
		return err
	}

	// Verify that the topic is configured for this example.
	return configureTopic(ctx, client, config.Kafka)
}

func ensureTopicExists(ctx context.Context, client *kadm.Client, config kafka.Config) error {
	// Validate that we have a topic name.
	if len(config.Topic) == 0 {
		return fmt.Errorf("no topic name provided")
	}

	// Check for existence of the topic.
	exists, err := kafka.TopicExists(ctx, client, config)
	if err != nil {
		return err
	}

	// If it already exists, continue.
	if exists {
		slog.Info("Topic already existed", "topic", config.Topic)
		return nil
	}

	// Topic does not exist, so we fall through to create it.
	if err := kafka.CreateTopic(ctx, client, config); err != nil {
		return err
	}
	return err
}

func configureTopic(ctx context.Context, client *kadm.Client, config kafka.Config) error {
	// Validate provided configuration.
	if config.ArchiveKind != "ICEBERG" {
		return fmt.Errorf("bufstream.archive.kind must be ICEBERG")
	}
	if len(config.Catalog) == 0 {
		return fmt.Errorf("bufstream.archive.iceberg.catalog is required")
	}
	if len(config.Table) == 0 {
		return fmt.Errorf("bufstream.archive.iceberg.table is required")
	}

	// Configure the topic.
	if err := kafka.ConfigureTopic(ctx, client, config); err != nil {
		return err
	}

	slog.Info("Configured topic for iceberg", "topic", config.Topic, "catalog", config.Catalog, "table", config.Table)
	return nil
}
