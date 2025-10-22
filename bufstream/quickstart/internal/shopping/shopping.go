package shopping

import (
	"math/rand/v2"

	shoppingv1 "github.com/bufbuild/buf-examples/bufstream/quickstart/gen/shopping/v1"
	"github.com/google/uuid"
)

// NewValidCart creates a valid cart with a random set of line items.
func NewValidCart() *shoppingv1.Cart {
	return NewCart(
		newRandomLineItems(),
	)
}

// NewCart is a helper function for creating shoppingv1.Cart based on
// provided line items.
func NewCart(lineItems []*shoppingv1.LineItem) *shoppingv1.Cart {
	cart := &shoppingv1.Cart{
		CartId:    uuid.New().String(),
		LineItems: lineItems,
	}

	return cart
}

func newRandomLineItems() []*shoppingv1.LineItem {
	maxItems := min(10, len(catalog))
	numItems := rand.IntN(maxItems) + 1
	lineItems := make([]*shoppingv1.LineItem, 0, numItems)

	// Track product_ids to ensure uniqueness.
	usedProductIDs := make(map[string]bool)

	for len(lineItems) < numItems {
		product := randomProduct()

		// Skip if we've already used this product.
		if usedProductIDs[product.GetProductId()] {
			continue
		}

		usedProductIDs[product.GetProductId()] = true

		lineItems = append(lineItems, &shoppingv1.LineItem{
			LineItemId:     uuid.New().String(),
			Product:        product,
			Quantity:       uint64(rand.IntN(5) + 1),
			UnitPriceCents: product.GetUnitPriceCents(),
		})
	}

	return lineItems
}
