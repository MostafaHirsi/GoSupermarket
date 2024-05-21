package main

type Item struct {
	sku       string
	unitPrice int
	offerRule *OfferRule
}

type OfferRule struct {
	noOfItems   int
	offerAmount int
}
