// Copyright 2025 Buf Technologies, Inc.
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
	"log/slog"
	"net/http"
	"time"

	"github.com/bufbuild/buf-examples/protovalidate/connect-go/start/gen/invoice/v1/invoicev1connect"
	"github.com/bufbuild/buf-examples/protovalidate/connect-go/start/internal/invoice"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	addr := "localhost:8080"
	invoiceServer := invoice.NewService()
	mux := http.NewServeMux()

	mux.Handle(invoicev1connect.NewInvoiceServiceHandler(
		invoiceServer,
	))

	slog.Info("starting invoice server", slog.String("addr", addr))
	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 3 * time.Second,
		// Use h2c so we can serve HTTP/2 without TLS.
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	if err := server.ListenAndServe(); err != nil {
		slog.Error("error running application",
			slog.String("error", err.Error()),
		)
	}
}
