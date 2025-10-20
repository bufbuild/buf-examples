// Package main implements the producer of the demo.

// After producing these three records, the producer sleeps for one second, and then loops.
package main

import (
	"context"
	"errors"
	"log/slog"

	invoicev1 "github.com/bufbuild/buf-examples/bufstream/quickstart/finish/gen/invoice/v1"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/finish/internal/app"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/finish/internal/invoice"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/finish/internal/kafka"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/finish/internal/produce"
	"github.com/google/uuid"
)

func main() {
	// See the app package for the boilerplate we use to set up the producer and
	// consumer, including bound flags.
	app.Main(run)
}

func run(ctx context.Context, config app.Config) error {
	client, err := kafka.NewKafkaClient(config.Kafka, false)
	if err != nil {
		return err
	}
	defer client.Close()

	producer := produce.NewProducer[*invoicev1.Invoice](
		client,
		config.Kafka.Topic,
	)

	slog.InfoContext(ctx, "starting produce")
	for {
		id := newID()
		inv := invoice.NewValidInvoice()
		// inv.LineItems = []*invoicev1.LineItem{}

		if err := producer.ProduceProtobufMessage(ctx, id, inv); err != nil {
			if errors.Is(err, context.Canceled) {
				return err
			}
			slog.ErrorContext(ctx, "Error producing invoice", "error", err)
		} else {
			slog.InfoContext(ctx, "Successfully produced invoice")
		}
	}
}

// newID returns a new UUID.
//
// This is also used as the record key.
func newID() string {
	return uuid.New().String()
}
