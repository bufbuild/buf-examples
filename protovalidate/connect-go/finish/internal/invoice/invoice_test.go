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
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"buf.build/go/protovalidate"
	"connectrpc.com/connect"
	connect_validate "connectrpc.com/validate"
	invoicev1 "github.com/bufbuild/buf-examples/protovalidate/connect-go/finish/gen/invoice/v1"
	"github.com/bufbuild/buf-examples/protovalidate/connect-go/finish/gen/invoice/v1/invoicev1connect"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// violationSpec is a simple representation of fields tested when inspecting
// a connect.Error that we expect to contain Protovalidate validate.Violations
// messages.
type violationSpec struct {
	ruleID    string
	fieldPath string
	message   string
}

// Test the CreateInvoice service method and verify that Protovalidate
// rules are honored.
func TestCreateInvoice(t *testing.T) {
	t.Parallel()

	// Create a handler with Connect's Protovalidate interceptor wrapping
	// its functions.
	interceptor, err := connect_validate.NewInterceptor()
	require.NoError(t, err)

	invoiceService := NewService()
	path, handler := invoicev1connect.NewInvoiceServiceHandler(
		invoiceService,
		connect.WithInterceptors(interceptor),
	)

	// Start a server.
	mux := http.NewServeMux()
	mux.Handle(path, handler)
	srv := startHTTPServer(t, mux)

	// Create a client shared by all of our test cases.
	invoiceServiceClient := invoicev1connect.NewInvoiceServiceClient(srv.Client(), srv.URL)

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
					ruleID:    "required",
					fieldPath: "invoice",
					message:   "value is required",
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
					ruleID:    "string.uuid_empty",
					fieldPath: "invoice.invoice_id",
					message:   "value is empty, which is not a valid UUID",
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
					ruleID:    "string.uuid",
					fieldPath: "invoice.invoice_id",
					message:   "value must be a valid UUID",
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
					ruleID:    "string.uuid_empty",
					fieldPath: "invoice.account_id",
					message:   "value is empty, which is not a valid UUID",
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
					ruleID:    "string.uuid",
					fieldPath: "invoice.account_id",
					message:   "value must be a valid UUID",
				},
			},
		},
		"InvoiceDate must not have a time component": {
			producer: func(invoice *invoicev1.Invoice) *invoicev1.Invoice {
				invoice.InvoiceDate = timestamppb.Now()
				return invoice
			},
			violations: []violationSpec{
				{
					ruleID:    "invoice_date.no_time",
					fieldPath: "invoice.invoice_date",
					message:   "time component should be zero",
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
					ruleID:    "repeated.min_items",
					fieldPath: "invoice.line_items",
					message:   "value must contain at least 1 item(s)",
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
					ruleID:    "line_items.unique_line_item_id",
					fieldPath: "invoice.line_items",
					message:   "all line_item_id values must be unique",
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
					ruleID:    "line_items.logically_unique",
					fieldPath: "invoice.line_items",
					message:   "line items must be unique combinations of product_id and unit_price",
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
					ruleID:    "line_item_id.not.product_id",
					fieldPath: "invoice.line_items[0]",
					message:   "line_item_id can't also be a product_id",
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
					ruleID:    "string.uuid_empty",
					fieldPath: "invoice.line_items[0].line_item_id",
					message:   "value is empty, which is not a valid UUID",
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
					ruleID:    "string.uuid",
					fieldPath: "invoice.line_items[0].line_item_id",
					message:   "value must be a valid UUID",
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
					ruleID:    "string.uuid_empty",
					fieldPath: "invoice.line_items[0].product_id",
					message:   "value is empty, which is not a valid UUID",
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
					ruleID:    "string.uuid",
					fieldPath: "invoice.line_items[0].product_id",
					message:   "value must be a valid UUID",
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
					ruleID:    "uint64.gt",
					fieldPath: "invoice.line_items[0].quantity",
					message:   "value must be greater than 0",
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
				var connectError *connect.Error
				require.ErrorAs(t, err, &connectError)
				assert.Equal(t, connect.CodeInvalidArgument, connectError.Code())
				checkConnectError(t, connectError, testCase.violations)
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

// checkConnectError receives a connect.Error that's expected to be returned
// by the connectrpc.com/validate validation intercepter. It's expected to
// contain one ErrorDetail, and for the value of that ErrorDetail to be a
// validate.Violations message.
//
// It uses a slice of violationSpec as an expectation for the contents of
// validate.Violations.GetViolations().
func checkConnectError(t *testing.T, connectError *connect.Error, specs []violationSpec) {
	t.Helper()

	// The response should contain one entry in Details and it should be accessible.
	details := connectError.Details()
	if len(details) != 1 {
		t.Errorf("Connect error had %d details instead of one", len(details))
	}
	detail, err := details[0].Value()
	if err != nil {
		t.Errorf("Couldn't get value of first error detail: %v", err)
	}

	// The value of the first item in Details should be a validate.Violations.
	switch violations := detail.(type) {
	case *validate.Violations:
		if specs == nil {
			t.Fatalf("Received a connect error with Violations but no spec was provided")
		}

		// Get the slice of individual violations.
		allViolations := violations.GetViolations()

		// It should be the same length as the provided spec.
		if len(allViolations) != len(specs) {
			t.Fatalf("violations returned %d violations instead of %d", len(allViolations), len(specs))
		}
		for i, spec := range specs {
			violation := allViolations[i]

			// Validate that it meets the expectations in the violationSpec.
			if violation.GetRuleId() != spec.ruleID {
				t.Fatalf("Wrong ruleID. Expected \"%v\", not \"%v\"", spec.ruleID, violation.GetRuleId())
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

// startHTTPServer starts an HTTP2 server for use within our test and
// registers it for closing within cleanup.
func startHTTPServer(tb testing.TB, h http.Handler) *httptest.Server {
	tb.Helper()
	srv := httptest.NewUnstartedServer(h)
	srv.EnableHTTP2 = true
	srv.Start()
	tb.Cleanup(srv.Close)
	return srv
}
