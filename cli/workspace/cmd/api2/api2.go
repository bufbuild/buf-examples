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
	"context"
	"flag"
	"log"
	"net/http"

	"connectrpc.com/connect/v2"
	"connectrpc.com/connect/v2/connecthttp"
	apiv2 "github.com/bufbuild/buf-examples/workspace/gen/proto/go/api/v2"
	"github.com/bufbuild/buf-examples/workspace/gen/proto/go/api/v2/apiv2connect"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var ErrNeedsImplementation = connect.NewError(connect.CodeUnimplemented, "needs implementation")

func main() {
	addr := flag.String("listen", "127.0.0.1:8080", "listen to address on")
	flag.Parse()

	log.Printf("Observability service starting")

	connectServer := connect.NewServer()
	apiv2connect.RegisterObservabilityServiceHandler(connectServer, &service{})
	connecthttp.Mount(http.DefaultServeMux, connectServer)
	log.Printf("Handling connect service: %v", apiv2connect.ObservabilityServiceName)

	log.Printf("Listening on: %v", *addr)
	// Serve HTTP/2 without TLS (h2c) alongside HTTP/1.1.
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)
	server := &http.Server{
		Addr:      *addr,
		Handler:   http.DefaultServeMux,
		Protocols: protocols,
	}
	err := server.ListenAndServe()

	if err != http.ErrServerClosed {
		log.Printf("Error running or stopping: %v", err)
	}
}

type service struct{}

func (s *service) GetLogs(
	ctx context.Context,
	req *apiv2.GetLogsRequest,
) (
	*apiv2.GetLogsResponse,
	error,
) {
	return nil, ErrNeedsImplementation
}

func (s *service) PutLogs(
	ctx context.Context,
	req *apiv2.PutLogsRequest,
) (
	*emptypb.Empty,
	error,
) {
	return nil, ErrNeedsImplementation
}

func (s *service) GetMetrics(
	ctx context.Context,
	req *apiv2.GetMetricsRequest,
) (
	*apiv2.GetMetricsResponse,
	error,
) {
	return nil, ErrNeedsImplementation
}
