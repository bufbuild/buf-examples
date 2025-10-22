// Package main implements the consumer of the demo.
// The consumer will read as many records it can at once, print what it received,
package main

import (
	"context"
	"log/slog"

	shoppingv1 "github.com/bufbuild/buf-examples/bufstream/quickstart/gen/shopping/v1"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/app"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/consume"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/kafka"
)

func main() {
	// See the app package for the boilerplate we use to set up the producer and
	// consumer, including bound flags.
	app.Main(run)
}

func run(ctx context.Context, config app.Config) error {
	client, err := kafka.NewKafkaClient(config.Kafka, true)
	if err != nil {
		return err
	}
	defer client.Close()

	consumer := consume.NewConsumer(
		client,
		config.Kafka.Topic,
		consume.WithMessageHandler(handleInvoice),
	)

	slog.InfoContext(ctx, "starting consume")
	for {
		// Read as many messages as we can.
		//
		// Only return error if there is an unexpected system error. Of note, an error is not
		// returned if the data that the consumer receives is malformed.
		if err := consumer.Consume(ctx); err != nil {
			return err
		}
		// time.Sleep(time.Second)
	}
}

func handleInvoice(ctx context.Context, invoice *shoppingv1.Cart) error {
	//json, err := protojson.Marshal(invoice)
	//if err != nil {
	//	return err
	//}
	lineItems := len(invoice.GetLineItems())
	if lineItems == 0 {
		slog.Error("Hey, why are you sending me invoices with no line items???")
	}
	//slog.InfoContext(ctx, fmt.Sprintf("consumed invoice %s", string(json)))
	return nil
}
