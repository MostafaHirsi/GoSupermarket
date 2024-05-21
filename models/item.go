package models

type Item struct {
	SKU       string
	UnitPrice int
	OfferRule *OfferRule
}

type OfferRule struct {
	NoOfItems   int
	OfferAmount int
}
