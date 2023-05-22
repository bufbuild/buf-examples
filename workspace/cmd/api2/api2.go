package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"

	apiv2 "github.com/bufbuild/buf-examples/workspace/gen/proto/go/api/v2"
	"github.com/bufbuild/buf-examples/workspace/gen/proto/go/api/v2/apiv2connect"
	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var ErrNeedsImplementation = connect.NewError(connect.CodeUnimplemented, errors.New("needs implementation"))

func main() {
	addr := flag.String("listen", "127.0.0.1:8080", "listen to address on")
	flag.Parse()

	log.Printf("Observability service starting")

	path, handler := apiv2connect.NewObservabilityServiceHandler(&service{})
	http.Handle(path, handler)
	log.Printf("Handling connect service at prefix: %v", path)

	log.Printf("Listening on: %v", *addr)
	err := http.ListenAndServe(
		*addr,
		h2c.NewHandler(http.DefaultServeMux, &http2.Server{}),
	)

	if err != http.ErrServerClosed {
		log.Printf("Error running or stopping: %v", err)
	}
}

type service struct{}

func (s *service) GetLogs(
	ctx context.Context,
	req *connect.Request[apiv2.GetLogsRequest],
) (
	*connect.Response[apiv2.GetLogsResponse],
	error,
) {
	return nil, ErrNeedsImplementation
}

func (s *service) PutLogs(
	ctx context.Context,
	req *connect.Request[apiv2.PutLogsRequest],
) (
	*connect.Response[emptypb.Empty],
	error,
) {
	return nil, ErrNeedsImplementation
}

func (s *service) GetMetrics(
	ctx context.Context,
	req *connect.Request[apiv2.GetMetricsRequest],
) (
	*connect.Response[apiv2.GetMetricsResponse],
	error,
) {
	return nil, ErrNeedsImplementation
}
