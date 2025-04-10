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
	"net"
	"testing"
	"time"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	invoicev1 "github.com/bufbuild/buf-examples/protovalidate/grpc-go/finish/gen/invoice/v1"
	"github.com/bufbuild/protovalidate-go"
	"github.com/google/uuid"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// violationSpec is a simple representation of fields tested when inspecting
// a connect.Error that we expect to contain Protovalidate validate.Violations
// messages.
type violationSpec struct {
	constraintID string
	fieldPath    string
	message      string
}

// Test the CreateInvoice service method and verify that Protovalidate
// rules are honored.
func TestCreateInvoice(t *testing.T) {
	t.Parallel()

	validator, err := protovalidate.New()
	require.NoError(t, err)

	interceptor := protovalidate_middleware.UnaryServerInterceptor(validator)

	// Start a server.
	listener := bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor),
	)
	invoiceService := NewService()
	invoicev1.RegisterInvoiceServiceServer(grpcServer, invoiceService)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			t.Errorf("failed to serve: %v", err)
		}
	}()

	// Create a client shared by all of our test cases.
	resolver.SetDefaultScheme("passthrough")
	conn, err := grpc.NewClient("bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	invoiceServiceClient := invoicev1.NewInvoiceServiceClient(conn)

	testCases := map[string]struct {
		producer   func(invoice *invoicev1.Invoice) *invoicev1.Invoice
		violations []violationSpec
	}{
		"A valid invoice passes validation": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				return invoice
			},
		},
		"A request without an invoice fails validation": {
			producer: func(_ *invoicev1.Invoice) *invoicev1.Invoice {
				return nil
			},
			violations: []violationSpec{
				{
					constraintID: "required",
					fieldPath:    "invoice",
					message:      "value is required",
				},
			},
		},
		"InvoiceId is required": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.InvoiceId = ""
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "string.uuid_empty",
					fieldPath:    "invoice.invoice_id",
					message:      "value is empty, which is not a valid UUID",
				},
			},
		},
		"InvoiceId must be a UUID": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.InvoiceId = "this not a uuid"
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "string.uuid",
					fieldPath:    "invoice.invoice_id",
					message:      "value must be a valid UUID",
				},
			},
		},
		"CustomerId is required": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.AccountId = ""
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "string.uuid_empty",
					fieldPath:    "invoice.account_id",
					message:      "value is empty, which is not a valid UUID",
				},
			},
		},
		"CustomerId must be a UUID": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.AccountId = "this not a uuid"
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "string.uuid",
					fieldPath:    "invoice.account_id",
					message:      "value must be a valid UUID",
				},
			},
		},
		"InvoiceDate may not have a time component": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.InvoiceDate = timestamppb.Now()
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "invoice_date.no_time",
					fieldPath:    "invoice.invoice_date",
					message:      "time component should be zero",
				},
			},
		},
		"At least one line item must be provided": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.LineItems = make([]*invoicev1.LineItem, 0)
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "repeated.min_items",
					fieldPath:    "invoice.line_items",
					message:      "value must contain at least 1 item(s)",
				},
			},
		},
		"Two line items cannot have the same identifier": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.GetLineItems()[0].LineItemId = invoice.GetLineItems()[1].GetLineItemId()
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "line_items.unique_line_item_id",
					fieldPath:    "invoice.line_items",
					message:      "all line_item_id values must be unique",
				},
			},
		},
		"Two line items cannot have the same product_id and unit price": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.GetLineItems()[0].ProductId = invoice.GetLineItems()[1].GetProductId()
				invoice.GetLineItems()[0].UnitPrice = invoice.GetLineItems()[1].GetUnitPrice()
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "line_items.logically_unique",
					fieldPath:    "invoice.line_items",
					message:      "line items must be unique combinations of product_id and unit_price",
				},
			},
		},
		"A line item cannot have the same line_item_id and product_id": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.GetLineItems()[0].ProductId = invoice.GetLineItems()[0].GetLineItemId()
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "line_item_id.not.product_id",
					fieldPath:    "invoice.line_items[0]",
					message:      "line_item_id can't also be a product_id",
				},
			},
		},
		"LineItem LineItemId is required": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.GetLineItems()[0].LineItemId = ""
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "string.uuid_empty",
					fieldPath:    "invoice.line_items[0].line_item_id",
					message:      "value is empty, which is not a valid UUID",
				},
			},
		},
		"LineItem LineItemId must be a UUID": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.GetLineItems()[0].LineItemId = "this is not a uuid"
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "string.uuid",
					fieldPath:    "invoice.line_items[0].line_item_id",
					message:      "value must be a valid UUID",
				},
			},
		},
		"LineItem ProductId is required": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.GetLineItems()[0].ProductId = ""
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "string.uuid_empty",
					fieldPath:    "invoice.line_items[0].product_id",
					message:      "value is empty, which is not a valid UUID",
				},
			},
		},
		"LineItem ProductId must be a UUID": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.GetLineItems()[0].ProductId = "this is not a uuid"
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "string.uuid",
					fieldPath:    "invoice.line_items[0].product_id",
					message:      "value must be a valid UUID",
				},
			},
		},
		"LineItem Quantity must be greater than zero": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.GetLineItems()[0].Quantity = 0
				return invoice
			},
			violations: []violationSpec{
				{
					constraintID: "uint64.gt",
					fieldPath:    "invoice.line_items[0].quantity",
					message:      "value must be greater than 0",
				},
			},
		},
		"LineItem UnitPrice can be zero": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.GetLineItems()[0].UnitPrice = 0
				return invoice
			},
		},
		"LineItem UnitPrice can be nonzero": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.GetLineItems()[0].UnitPrice = 1
				return invoice
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			invoice := testCase.producer(newValidInvoice())

			// Make our request.
			res, err := invoiceServiceClient.CreateInvoice(
				context.Background(),
				&invoicev1.CreateInvoiceRequest{
					Invoice: invoice,
				},
			)

			// If we expect a nonzero response code, check our error against
			// our expected violations.
			if len(testCase.violations) > 0 {
				require.Error(t, err)
				responseStatus, ok := status.FromError(err)
				require.True(t, ok)
				assert.Equal(t, codes.InvalidArgument, responseStatus.Code())
				checkStatusError(t, responseStatus, testCase.violations)
			} else {
				// Otherwise, ensure we have no error and receive expected results.
				require.NoError(t, err)
				assert.NotNil(t, res)
				assert.NotNil(t, res.GetInvoiceId())
				assert.Equal(t, uint64(1), res.GetVersion())
			}
		})
	}
}

func newValidInvoice() *invoicev1.Invoice {
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)

	invoice := &invoicev1.Invoice{
		InvoiceId:   uuid.New().String(),
		AccountId:   uuid.New().String(),
		InvoiceDate: timestamppb.New(today),
		LineItems: []*invoicev1.LineItem{
			{
				LineItemId: uuid.New().String(),
				ProductId:  uuid.New().String(),
				Quantity:   1,
				UnitPrice:  1,
			},
			{
				LineItemId: uuid.New().String(),
				ProductId:  uuid.New().String(),
				Quantity:   1,
				UnitPrice:  1,
			},
			{
				LineItemId: uuid.New().String(),
				ProductId:  uuid.New().String(),
				Quantity:   1,
				UnitPrice:  1,
			},
		},
	}

	return invoice
}

// checkStatusError receives a status.Error that's expected to be returned
// by the grpc-ecosystem validation intercepter. It's expected to contain one
// Detail, and for the value of that Detail to be a validate.Violations message.
//
// It uses a slice of violationSpec as an expectation for the contents of
// validate.Violations.GetViolations().
func checkStatusError(t *testing.T, responseStatus *status.Status, specs []violationSpec) {
	t.Helper()

	details := responseStatus.Details()
	if len(details) != 1 {
		t.Errorf("Status error had %d details instead of one", len(details))
	}
	detail := details[0]
	switch violations := detail.(type) {
	case *validate.Violations:
		if specs == nil {
			t.Fatalf("Received a connect error with Violations but no spec was provided")
		}

		allViolations := violations.GetViolations()
		if len(allViolations) != len(specs) {
			t.Fatalf("violations returned %d violations instead of %d", len(allViolations), len(specs))
		}
		for i, spec := range specs {
			violation := allViolations[i]

			// Validate that it meets the expectations in the violationSpec.
			if violation.GetConstraintId() != spec.constraintID {
				t.Fatalf("Wrong constraintID. Expected \"%v\", not \"%v\"", spec.constraintID, violation.GetConstraintId())
			}
			fieldPath := protovalidate.FieldPathString(violation.GetField())
			if fieldPath != spec.fieldPath {
				t.Fatalf("Wrong fieldPath. Expected \"%v\", not \"%v\"", spec.fieldPath, fieldPath)
			}
			if violation.GetMessage() != spec.message {
				t.Fatalf("Wrong message. Expected \"%v\", not \"%v\"", spec.message, violation.GetMessage())
			}
		}
	default:
		t.Fatalf("Received unexpected type of detail: %v", detail)
	}
}
