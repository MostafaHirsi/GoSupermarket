package checkout

import "errors"

type ICheckout interface {
	Scan(sku string) (err error)

	GetTotalPrice() (totalPrice int, err error)
}

type Checkout struct {
}

func (checkout Checkout) Scan(sku string) error {
	return errors.New("not yet implemented")
}

func (checkout Checkout) GetTotalPrice() (*int, error) {
	return nil, errors.New("not yet implemented")
}
