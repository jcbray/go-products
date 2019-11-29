package database

import "errors"

type Mock struct{}

const FAKE_ITEM_PRICE = 3.12
const FAKE_CURRENCY_CODE = "USD"

func (Mock) Connect() error {
	return nil
}

func (Mock) GetPriceInfo(id string, currentPrice chan CurrentPrice) {
	var price CurrentPrice
	switch id { //123457 //123458 //123459 123460
	case "123456", "123457": // success
		price.ID = id
		price.Value = FAKE_ITEM_PRICE
		price.CurrencyCode = FAKE_CURRENCY_CODE
	case "123459", "123461": // no content
		price.ID = ""
	case "123460", "123458", "123462": // fatal error
		price.Error = "fatal price error"
	}
	currentPrice <- price
	return
}

func (Mock) PutPriceInfo(price CurrentPrice) error {
	switch price.ID {
	case "123457", "123461": // no content
		price.ID = ""
		return errors.New("item not found")
	case "123458", "123462": // fatal error
		return errors.New("fatal price error")
	}
	return nil
}
