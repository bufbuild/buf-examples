package invoice

import (
	"time"

	invoicev1 "github.com/bufbuild/buf-examples/bufstream/quickstart/finish/gen/invoice/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewValidInvoice() *invoicev1.Invoice {
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
