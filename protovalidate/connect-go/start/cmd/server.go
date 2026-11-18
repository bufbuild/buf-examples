// Copyright 2020-2026 Buf Technologies, Inc.
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

	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"
	"github.com/bufbuild/buf-examples/protovalidate/connect-go/start/gen/invoice/v1/invoicev1connect"
	"github.com/bufbuild/buf-examples/protovalidate/connect-go/start/internal/invoice"
)

func main() {
	addr := "localhost:8080"
	invoiceServer := invoice.NewService()
	mux := http.NewServeMux()

	connectServer := connect.NewServer()
	invoicev1connect.RegisterInvoiceServiceHandler(connectServer, invoiceServer)
	connecthttp.Mount(mux, connectServer)

	slog.Info("starting invoice server", slog.String("addr", addr))
	// Serve HTTP/2 without TLS (h2c) alongside HTTP/1.1.
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		Protocols:         protocols,
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		slog.Error("error running application",
			slog.String("error", err.Error()),
		)
	}
}
