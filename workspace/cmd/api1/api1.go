package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	apiv1 "github.com/bufbuild/buf-examples/workspace/gen/proto/go/api/v1"
	"github.com/bufbuild/buf-examples/workspace/gen/proto/go/api/v1/apiv1connect"
	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	const addr = "127.0.0.1:8080"

	log.Printf("Observability service starting")

	mux := http.NewServeMux()

	path, handler := apiv1connect.NewObservabilityServiceHandler(&service{})
	mux.Handle(path, handler)
	log.Printf("Handling connect service at prefix: %v", path)

	log.Printf("Listening on: %v", addr)
	err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	)

	if err != http.ErrServerClosed {
		log.Printf("Error running or stopping: %v", err)
	}
}

type service struct {
}

func (s *service) GetLogs(
	ctx context.Context,
	req *connect.Request[apiv1.GetLogsRequest],
) (
	*connect.Response[apiv1.GetLogsResponse],
	error,
) {
	return nil, errors.New("Needs impementation")
}

func (s *service) PutLogs(
	ctx context.Context,
	req *connect.Request[apiv1.PutLogsRequest],
) (
	*connect.Response[emptypb.Empty],
	error,
) {
	return nil, errors.New("Needs implementation")
}

func (s *service) GetMetrics(
	ctx context.Context,
	req *connect.Request[apiv1.GetMetricsRequest],
) (
	*connect.Response[apiv1.GetMetricsResponse],
	error,
) {
	return nil, errors.New("Needs implementation")
}
