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
)

var (
	CategoryHomeGarden = &shoppingv1.Category{
		Id:   "home_garden",
		Name: "Home & Garden",
	}
	CategoryKitchenDining = &shoppingv1.Category{
		Id:   "kitchen_dining",
		Name: "Kitchen & Dining",
	}
	CategorySportsOutdoors = &shoppingv1.Category{
		Id:   "sports_outdoors",
		Name: "Sports & Outdoors",
	}
	CategoryBooksStationery = &shoppingv1.Category{
		Id:   "books_stationery",
		Name: "Books & Stationery",
	}
	CategoryElectronicsAccessories = &shoppingv1.Category{
		Id:   "electronics_accessories",
		Name: "Electronics & Accessories",
	}
	CategoryPersonalCare = &shoppingv1.Category{
		Id:   "personal_care",
		Name: "Personal Care",
	}
)

// catalog holds all available products.
var catalog = []*shoppingv1.Product{
	// Home & Garden
	{ProductId: "550e8400-e29b-41d4-a716-446655440001", Sku: "home_garden_001", Name: "Ceramic Plant Pot Set", Category: CategoryHomeGarden, UnitPriceCents: 2999},
	{ProductId: "550e8400-e29b-41d4-a716-446655440002", Sku: "home_garden_002", Name: "LED String Lights", Category: CategoryHomeGarden, UnitPriceCents: 1899},
	{ProductId: "550e8400-e29b-41d4-a716-446655440003", Sku: "home_garden_003", Name: "Bamboo Garden Tool Kit", Category: CategoryHomeGarden, UnitPriceCents: 4599},
	{ProductId: "550e8400-e29b-41d4-a716-446655440004", Sku: "home_garden_004", Name: "Outdoor Wind Chimes", Category: CategoryHomeGarden, UnitPriceCents: 3499},
	{ProductId: "550e8400-e29b-41d4-a716-446655440005", Sku: "home_garden_005", Name: "Succulent Starter Collection", Category: CategoryHomeGarden, UnitPriceCents: 2499},
	{ProductId: "550e8400-e29b-41d4-a716-446655440006", Sku: "home_garden_006", Name: "Decorative Throw Pillows", Category: CategoryHomeGarden, UnitPriceCents: 3999},

	// Kitchen & Dining
	{ProductId: "550e8400-e29b-41d4-a716-446655440011", Sku: "kitchen_dining_001", Name: "Stainless Steel Mixing Bowls", Category: CategoryKitchenDining, UnitPriceCents: 3599},
	{ProductId: "550e8400-e29b-41d4-a716-446655440012", Sku: "kitchen_dining_002", Name: "Silicone Baking Mat Set", Category: CategoryKitchenDining, UnitPriceCents: 1999},
	{ProductId: "550e8400-e29b-41d4-a716-446655440013", Sku: "kitchen_dining_003", Name: "Glass Food Storage Containers", Category: CategoryKitchenDining, UnitPriceCents: 4299},
	{ProductId: "550e8400-e29b-41d4-a716-446655440014", Sku: "kitchen_dining_004", Name: "Bamboo Cutting Board", Category: CategoryKitchenDining, UnitPriceCents: 2899},
	{ProductId: "550e8400-e29b-41d4-a716-446655440015", Sku: "kitchen_dining_005", Name: "Coffee Bean Grinder", Category: CategoryKitchenDining, UnitPriceCents: 5999},

	// Sports & Outdoors
	{ProductId: "550e8400-e29b-41d4-a716-446655440021", Sku: "sports_outdoors_001", Name: "Yoga Mat with Strap", Category: CategorySportsOutdoors, UnitPriceCents: 3999},
	{ProductId: "550e8400-e29b-41d4-a716-446655440022", Sku: "sports_outdoors_002", Name: "Insulated Water Bottle", Category: CategorySportsOutdoors, UnitPriceCents: 2499},
	{ProductId: "550e8400-e29b-41d4-a716-446655440023", Sku: "sports_outdoors_003", Name: "Resistance Band Set", Category: CategorySportsOutdoors, UnitPriceCents: 1999},
	{ProductId: "550e8400-e29b-41d4-a716-446655440024", Sku: "sports_outdoors_004", Name: "Camping Lantern", Category: CategorySportsOutdoors, UnitPriceCents: 3299},
	{ProductId: "550e8400-e29b-41d4-a716-446655440025", Sku: "sports_outdoors_005", Name: "Trail Running Backpack", Category: CategorySportsOutdoors, UnitPriceCents: 6999},
	{ProductId: "550e8400-e29b-41d4-a716-446655440026", Sku: "sports_outdoors_006", Name: "Portable Hammock", Category: CategorySportsOutdoors, UnitPriceCents: 4599},

	// Books & Stationery
	{ProductId: "550e8400-e29b-41d4-a716-446655440031", Sku: "books_stationery_001", Name: "Leather Journal Notebook", Category: CategoryBooksStationery, UnitPriceCents: 2799},
	{ProductId: "550e8400-e29b-41d4-a716-446655440032", Sku: "books_stationery_002", Name: "Gel Pen Collection", Category: CategoryBooksStationery, UnitPriceCents: 1299},
	{ProductId: "550e8400-e29b-41d4-a716-446655440033", Sku: "books_stationery_003", Name: "Desk Organizer Set", Category: CategoryBooksStationery, UnitPriceCents: 3499},
	{ProductId: "550e8400-e29b-41d4-a716-446655440034", Sku: "books_stationery_004", Name: "Bookmark Set", Category: CategoryBooksStationery, UnitPriceCents: 899},
	{ProductId: "550e8400-e29b-41d4-a716-446655440035", Sku: "books_stationery_005", Name: "Watercolor Painting Kit", Category: CategoryBooksStationery, UnitPriceCents: 4999},

	// Electronics & Accessories
	{ProductId: "550e8400-e29b-41d4-a716-446655440041", Sku: "electronics_acc_001", Name: "Wireless Phone Charger", Category: CategoryElectronicsAccessories, UnitPriceCents: 2999},
	{ProductId: "550e8400-e29b-41d4-a716-446655440042", Sku: "electronics_acc_002", Name: "USB-C Hub Adapter", Category: CategoryElectronicsAccessories, UnitPriceCents: 4599},
	{ProductId: "550e8400-e29b-41d4-a716-446655440043", Sku: "electronics_acc_003", Name: "Bluetooth Earbuds Case", Category: CategoryElectronicsAccessories, UnitPriceCents: 1599},
	{ProductId: "550e8400-e29b-41d4-a716-446655440044", Sku: "electronics_acc_004", Name: "Laptop Sleeve", Category: CategoryElectronicsAccessories, UnitPriceCents: 3299},
	{ProductId: "550e8400-e29b-41d4-a716-446655440045", Sku: "electronics_acc_005", Name: "Cable Management Clips", Category: CategoryElectronicsAccessories, UnitPriceCents: 999},
	{ProductId: "550e8400-e29b-41d4-a716-446655440046", Sku: "electronics_acc_006", Name: "Portable Power Bank", Category: CategoryElectronicsAccessories, UnitPriceCents: 3999},

	// Personal Care
	{ProductId: "550e8400-e29b-41d4-a716-446655440051", Sku: "personal_care_001", Name: "Essential Oil Diffuser", Category: CategoryPersonalCare, UnitPriceCents: 3599},
	{ProductId: "550e8400-e29b-41d4-a716-446655440052", Sku: "personal_care_002", Name: "Bamboo Toothbrush Set", Category: CategoryPersonalCare, UnitPriceCents: 1499},
	{ProductId: "550e8400-e29b-41d4-a716-446655440053", Sku: "personal_care_003", Name: "Reusable Cotton Rounds", Category: CategoryPersonalCare, UnitPriceCents: 1299},
	{ProductId: "550e8400-e29b-41d4-a716-446655440054", Sku: "personal_care_004", Name: "Travel Toiletry Bag", Category: CategoryPersonalCare, UnitPriceCents: 2799},
}

// RandomProduct returns a randomly selected product from the catalog.
func randomProduct() *shoppingv1.Product {
	return catalog[rand.IntN(len(catalog))]
}
