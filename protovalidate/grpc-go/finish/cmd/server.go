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

package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"buf.build/go/protovalidate"
	invoicev1 "github.com/bufbuild/buf-examples/protovalidate/grpc-go/finish/gen/invoice/v1"
	invoice "github.com/bufbuild/buf-examples/protovalidate/grpc-go/finish/internal"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("error running application",
			slog.String("error", err.Error()),
		)
	}

	slog.Info("closing server")
}

func run(ctx context.Context) error {
	validator, err := protovalidate.New()
	if err != nil {
		return errors.New("failed to create protovalidate validator")
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(protovalidate_middleware.UnaryServerInterceptor(validator)),
	)
	reflection.Register(grpcServer)

	invoicev1.RegisterInvoiceServiceServer(grpcServer, invoice.NewService())

	errorGroup, ctx := errgroup.WithContext(ctx)

	errorGroup.Go(func() error {
		address := "localhost:50051"

		config := net.ListenConfig{}

		lis, err := config.Listen(ctx, "tcp", address)
		if err != nil {
			return fmt.Errorf("failed to listen on address %q: %w", address, err)
		}

		slog.Info("starting invoice service", slog.String("address", address))

		if err := grpcServer.Serve(lis); err != nil {
			return fmt.Errorf("failed to serve grpc service: %w", err)
		}

		return nil
	})

	errorGroup.Go(func() error {
		<-ctx.Done()

		grpcServer.GracefulStop()

		return ctx.Err()
	})

	return errorGroup.Wait()
}
