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

package shopping

import (
	"math/rand/v2"

	shoppingv1 "github.com/bufbuild/buf-examples/bufstream/quickstart/gen/shopping/v1"
	"github.com/google/uuid"
)

// NewValidCart creates a valid cart with a random set of line items.
func NewValidCart() *shoppingv1.Cart {
	return newCart(
		newRandomLineItems(),
	)
}

// NewInvalidCart creates a cart where at least one line item has an invalid
// quantity (0).
func NewInvalidCart() *shoppingv1.Cart {
	cart := NewValidCart()

	// Invalidate a random line item by setting quantity to 0
	if len(cart.LineItems) > 0 {
		invalidIndex := rand.IntN(len(cart.LineItems))
		cart.LineItems[invalidIndex].Quantity = 0
	}

	return cart
}

func newCart(lineItems []*shoppingv1.LineItem) *shoppingv1.Cart {
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
