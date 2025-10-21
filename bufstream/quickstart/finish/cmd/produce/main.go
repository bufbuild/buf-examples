// Package main implements the producer of the demo.

// After producing these three records, the producer sleeps for one second, and then loops.
package main

import (
	"context"
	"errors"
	"log/slog"
	"sync"

	shoppingv1 "github.com/bufbuild/buf-examples/bufstream/quickstart/finish/gen/shopping/v1"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/finish/internal/app"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/finish/internal/kafka"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/finish/internal/produce"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/finish/internal/shopping"
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

	producer := produce.NewProducer[*shoppingv1.Invoice](
		client,
		config.Kafka.Topic,
	)

	numWorkers := 20
	produced := 0
	slog.InfoContext(ctx, "starting produce", "workers", numWorkers)

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					id := newID()
					inv := shopping.NewValidInvoice()

					if err := producer.ProduceMessage(ctx, id, inv); err != nil {
						if errors.Is(err, context.Canceled) {
							return
						}
						slog.ErrorContext(ctx, "Error producing invoice", "error", err)
					}

					produced++
					if produced%10 == 0 {
						slog.InfoContext(ctx, "Produced invoice", "count", produced)
					}
				}
			}
		}()
	}

	wg.Wait()
	return nil
}

// newID returns a new UUID.
//
// This is also used as the record key.
func newID() string {
	return uuid.New().String()
}
