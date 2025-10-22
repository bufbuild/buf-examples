// Package main implements the producer of the demo.

// After producing these three records, the producer sleeps for one second, and then loops.
package main

import (
	"context"
	"errors"
	"log/slog"
	"math/rand/v2"
	"sync"
	"sync/atomic"

	shoppingv1 "github.com/bufbuild/buf-examples/bufstream/quickstart/gen/shopping/v1"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/app"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/kafka"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/produce"
	"github.com/bufbuild/buf-examples/bufstream/quickstart/internal/shopping"
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

	producer := produce.NewProducer[*shoppingv1.Cart](
		client,
		config.Kafka.Topic,
	)

	numWorkers := 50
	var attempts atomic.Int64
	var produced atomic.Int64
	var rejected atomic.Int64

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

					var inv *shoppingv1.Cart
					n := rand.IntN(100)
					if n < 3 {
						inv = shopping.NewCart(nil)
					} else {
						inv = shopping.NewValidCart()
					}

					currentAttempts := attempts.Add(1)
					if err := producer.ProduceMessage(ctx, id, inv); err != nil {
						if errors.Is(err, context.Canceled) {
							return
						}
						slog.ErrorContext(ctx, "error producing message", "id", id, "err", err)
						//json, _ := protojson.Marshal(inv)
						//slog.InfoContext(ctx, fmt.Sprintf("invoice rejected %s", string(json)))

						rejected.Add(1)
					} else {
						produced.Add(1)
					}

					if currentAttempts%100 == 0 {
						slog.InfoContext(ctx, "Producer running",
							"attempts", attempts.Load(),
							"produced", produced.Load(),
							"rejected", rejected.Load(), "n", n, "li", len(inv.GetLineItems()))

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
