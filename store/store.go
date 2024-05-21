package store

import (
	"errors"
	"supermarket_test/models"
)

type IStore interface {
	SearchStore(sku string) (foundItem models.Item, err error)

	AddToStore(item models.Item)
}

type Store struct {
	Items []models.Item
}

func (store Store) AddToStore(item models.Item) {
}

func (store Store) SearchStore(sku string) (foundItem *models.Item, err error) {
	return nil, errors.New("not yet implemented")
}
