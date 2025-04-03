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

package invoice

import (
	"context"

	invoicev1 "github.com/bufbuild/buf-examples/protovalidate/grpc-go/finish/gen/invoice/v1"
)

// Service is a gRPC handler for the RPC services defined
// in invoice_service.proto.
type Service struct {
	invoicev1.UnimplementedInvoiceServiceServer
}

// NewService creates a new Service.
func NewService() *Service {
	return &Service{}
}

func (s *Service) CreateInvoice(_ context.Context, req *invoicev1.CreateInvoiceRequest) (*invoicev1.CreateInvoiceResponse, error) {
	invoice := req.GetInvoice()

	// Handle the request: persist the invoice, or maybe place it on a stream
	// or queue where workers would handle its creation.

	return &invoicev1.CreateInvoiceResponse{
		InvoiceId: invoice.GetInvoiceId(),
		Version:   1,
	}, nil
}
