package products

import (
	"errors"
	"github.com/jcbray/go-products/api"
	"github.com/jcbray/go-products/database"
)

type Mock struct{}

func (Mock) Get(id string) (Products, error) {
	var product Products
	var error error
	switch id {
	case "12345": // success
		product.ID = id
		product.ProductDescription = api.FAKE_ITEM_DESCRIPTION
		product.CurrentPrice.Value = database.FAKE_ITEM_PRICE
		error = nil
	case "12346": // error
		error = errors.New("unexpected error")
	}
	return product, error
}

func (Mock) Put(products Products) error {
	switch products.ID {
	case "12345": // success
		return nil
	case "12346": // error
		return errors.New("unexpected error")
	}
	return nil
}
