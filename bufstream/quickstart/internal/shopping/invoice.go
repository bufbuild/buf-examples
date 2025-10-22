package shopping

import (
	"math/rand/v2"
	"time"

	shoppingv1 "github.com/bufbuild/buf-examples/bufstream/quickstart/gen/shopping/v1"
	"github.com/google/uuid"
	"google.golang.org/genproto/googleapis/type/date"
)

// NewValidInvoice creates a valid invoice with a random set of line items.
func NewValidInvoice() *shoppingv1.Invoice {
	return NewInvoice(
		NewRandomLineItems(),
	)
}

// NewInvoice is a helper function for creating shoppingv1.Invoice based on
// provided line items.
func NewInvoice(lineItems []*shoppingv1.LineItem) *shoppingv1.Invoice {
	now := time.Now()

	invoice := &shoppingv1.Invoice{
		InvoiceId: uuid.New().String(),
		AccountId: uuid.New().String(),
		InvoiceDate: &date.Date{
			Year:  int32(now.Year()),
			Month: int32(now.Month()),
			Day:   int32(now.Day()),
		},
		LineItems: lineItems,
	}

	return invoice
}

// NewRandomLineItems generates between 1 and 10 line items with random products
// and sensible quantities for each product type.
func NewRandomLineItems() []*shoppingv1.LineItem {
	maxItems := min(10, len(catalog))
	numItems := rand.IntN(maxItems) + 1
	lineItems := make([]*shoppingv1.LineItem, 0, numItems)

	// Track product_ids to ensure uniqueness.
	usedProductIds := make(map[string]bool)

	for len(lineItems) < numItems {
		product := RandomProduct()

		// Skip if we've already used this product.
		if usedProductIds[product.ProductId] {
			continue
		}

		usedProductIds[product.ProductId] = true

		lineItems = append(lineItems, &shoppingv1.LineItem{
			LineItemId:     uuid.New().String(),
			Product:        product,
			Quantity:       randomQuantityForProduct(product),
			UnitPriceCents: product.UnitPriceCents,
		})
	}

	return lineItems
}

// randomQuantityForProduct returns a quantity based on product category.
func randomQuantityForProduct(product *shoppingv1.Product) uint64 {
	switch product.Category.Id {
	case CategoryBooksStationery.Name:
		// Books and stationery: people buy multiple pens, notebooks, etc.
		return uint64(rand.IntN(5) + 1) // 1-5
	case CategoryPersonalCare.Name:
		// Personal care: typically buy 1-3 (stocking up on toothbrushes, etc.)
		return uint64(rand.IntN(3) + 1) // 1-3
	case CategoryKitchenDining.Name:
		// Kitchen items: usually 1-2 (mixing bowls, cutting boards)
		return uint64(rand.IntN(2) + 1) // 1-2
	case CategoryHomeGarden.Name:
		// Home & garden: typically 1-3 (plant pots, decorative items)
		return uint64(rand.IntN(3) + 1) // 1-3
	case CategorySportsOutdoors.Name:
		// Sports equipment: usually just 1 (yoga mat, backpack, etc.)
		return uint64(rand.IntN(2) + 1) // 1-2
	case CategoryElectronicsAccessories.Name:
		// Electronics accessories: 1-4 (cables, chargers, cases)
		return uint64(rand.IntN(4) + 1) // 1-4
	default:
		// Default: 1-3 items
		return uint64(rand.IntN(3) + 1)
	}
}
