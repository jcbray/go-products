package api

import "errors"

type Mock struct{}

const FAKE_ITEM_DESCRIPTION = "TEST ITEM"

func (Mock) GetItemInfo(id string, item chan Response) {
	var response Response
	switch id {
	case "123456", "123459", "123460": // success
		response.Product.Item.Tcin = id
		response.Product.Item.ProductDescription.Title = FAKE_ITEM_DESCRIPTION
	case "123457", "123461": // no content
		response.Product.Item.Tcin = ""
		response.Product.Item.ProductDescription.Title = ""
	case "123458", "123462": // fatal error
		response.Error = errors.New("fatal item error")
	}
	item <- response
}
