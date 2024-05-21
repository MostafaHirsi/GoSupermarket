package test

import (
	"errors"
	"supermarket_test/checkout"
	"testing"
)

// TestCheckoutScanSuccess calls checkout.scan with a sku string that is valid, should not throw an error
func TestCheckoutScanSuccess(t *testing.T) {
	checkout := checkout.Checkout{}
	err := checkout.Scan("U")
	if err != nil {
		t.Errorf("Received error: %v", err)
	}
}

// TestCheckoutScanFail calls checkout.scan with a sku string that is invalid, should return specific error
func TestCheckoutScanFail(t *testing.T) {
	checkout := checkout.Checkout{}
	expectedErr := errors.New("item does not exist!")
	actualErr := checkout.Scan("A")
	if actualErr != expectedErr {
		t.Errorf("Error actual = %v, and Expected = %v.", actualErr, expectedErr)
	}
}

// TestCheckoutTotalPriceWith0TotalPrice calls checkout.GetTotalPrice without scanning any items, should return 0
func TestCheckoutTotalPriceWith0TotalPrice(t *testing.T) {
	checkout := checkout.Checkout{}
	expectedTotalPrice := 0
	actualTotalPrice, err := checkout.GetTotalPrice()
	if err != nil {
		t.Errorf("Error: %v", err)
	} else if actualTotalPrice != &expectedTotalPrice {
		t.Errorf("Actual: %d, Expected %d", actualTotalPrice, expectedTotalPrice)
	}
}

// TestCheckoutTotalPriceWith100TotalPrice calls checkout.GetTotalPrice without scanning any items, should return 100
func TestCheckoutTotalPriceWith100TotalPrice(t *testing.T) {
	checkout := checkout.Checkout{}
	expectedTotalPrice := 100
	actualTotalPrice, err := checkout.GetTotalPrice()
	if err != nil {
		t.Errorf("Error: %v", err)
	} else if actualTotalPrice != &expectedTotalPrice {
		t.Errorf("Actual: %d, Expected %d", actualTotalPrice, expectedTotalPrice)
	}
}
