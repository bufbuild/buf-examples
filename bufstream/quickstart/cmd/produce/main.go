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

// Package main implements the producer of the quickstart.

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
					var inv *shoppingv1.Cart
					n := rand.IntN(100)
					currentAttempts := attempts.Add(1)
					if n < 1 {
						inv = shopping.NewInvalidCart()
					} else {
						inv = shopping.NewValidCart()
					}
					if err := producer.ProduceMessage(ctx, newID(), inv); err != nil {
						if errors.Is(err, context.Canceled) {
							return
						}
						slog.ErrorContext(ctx, "error producing message", "err", err)
					}

					if currentAttempts%100 == 0 {
						slog.InfoContext(ctx, "Producer running",
							"orders produced", attempts.Load())
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
