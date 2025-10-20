// Package main implements the consumer of the demo.
// The consumer will read as many records it can at once, print what it received,
package main

import (
	"context"
	"log/slog"

	invoicev1 "github.com/bufbuild/buf-examples/bufstream/quickstart/finish/gen/invoice/v1"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/finish/internal/app"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/finish/internal/consume"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/finish/internal/kafka"
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

	// NewSerde creates a CSR-based deserializer if there is a CSR URL,
	// otherwise it creates a single-type deserializer for demov1.EmailUpdated.
	// serde, err := csr.NewSerde[*demov1.EmailUpdated](ctx, config.CSR, config.Kafka.Topic)
	// if err != nil {
	// 	return err
	// }

	consumer := consume.NewConsumer(
		client,
		// serde,
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

func handleInvoice(ctx context.Context, invoice *invoicev1.Invoice) error {
	// json, err := protojson.Marshal(invoice)
	// if err != nil {
	// 	return err
	// }
	lineItems := len(invoice.GetLineItems())

	if lineItems == 0 {
		slog.Error("Hey, why are you sending me invoices with no line items???")
	} else {
		slog.Info("got an invoice with %d line items", "line items", lineItems)
	}
	// slog.InfoContext(ctx, fmt.Sprintf("consumed invoice %s", string(json)))
	return nil
}
