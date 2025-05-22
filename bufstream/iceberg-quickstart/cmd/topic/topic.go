package main

import (
	"context"
	"log/slog"

	"github.com/bufbuild/buf-examples/bufstream/iceberg-quickstart/pkg/app"
	"github.com/bufbuild/buf-examples/bufstream/iceberg-quickstart/pkg/kafka"
)

func main() {
	// See the app package for the boilerplate we use to set up the topic,
	// including bound flags.
	app.Main(run)
}

func run(ctx context.Context, config app.Config) error {
	return createTopic(ctx, config)
}

func createTopic(ctx context.Context, config app.Config) error {
	// Create an admin client.
	client, err := kafka.NewAdminClient(config.Kafka)
	if err != nil {
		return err
	}
	defer client.Close()

	// Check for existence of the topic.
	exists, err := kafka.TopicExists(ctx, client, config.Kafka)
	if err != nil {
		return err
	}
	// If it already exists, continue.
	if exists {
		slog.Info("Topic already exists", "topic", config.Kafka.Topic)
		return nil
	}
	// TODO: is it properly configured?

	// Topic does not exist, so we fall through to create it.
	if err := kafka.CreateTopic(ctx, client, config.Kafka); err != nil {
		return err
	}
	return err
}
