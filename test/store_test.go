package test

import (
	"supermarket_test/models"
	"supermarket_test/store"
	"testing"
)

// TestStoreSearchStoreSuccess calls store.searchstore with a sku string that is valid, should return Item
func TestStoreSearchStoreSuccess(t *testing.T) {
	store := store.Store{
		Items: []models.Item{
			{SKU: "A",
				UnitPrice: 10,
				OfferRule: &models.OfferRule{
					NoOfItems:   0,
					OfferAmount: 0,
				},
			},
			{
				SKU:       "B",
				UnitPrice: 20,
				OfferRule: &models.OfferRule{
					NoOfItems:   0,
					OfferAmount: 0,
				},
			},
		}}
	expectedItem := models.Item{SKU: "B", UnitPrice: 20}
	actualItem, err := store.SearchStore("B")

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if &expectedItem != actualItem {
		t.Errorf("Error actual = %v, and Expected = %v.", actualItem, expectedItem)
	}
}
