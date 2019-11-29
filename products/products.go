package products

import (
	"fmt"
	"github.com/jcbray/go-products/api"
	"github.com/jcbray/go-products/database"
)

type IProducts interface {
	Get(id string) (Products, error)
	Put(products Products) error
}

type Request struct{}

var iApiRequest api.IApiRequest = api.Request{}
var iMongoDb database.IMongoDb = database.Request{}

// Get product information based on the id inputted
func (Request) Get(id string) (Products, error) {
	var products Products
	itemOutputChannel := make(chan api.Response, 100)
	go iApiRequest.GetItemInfo(id, itemOutputChannel)

	priceOutputChannel := make(chan database.CurrentPrice, 100)
	go iMongoDb.GetPriceInfo(id, priceOutputChannel)

	itemOutput := <-itemOutputChannel
	close(itemOutputChannel)

	priceOutput := <-priceOutputChannel
	close(priceOutputChannel)

	if priceOutput.Error != "" {
		fmt.Println("fatal price error " + priceOutput.Error)
	}
	mapApiResponseToProducts(&products, itemOutput)

	if itemOutput.Error != nil {
		fmt.Println("fatal item error " + itemOutput.Error.Error())
		return products, itemOutput.Error
	}
	mapDatabaseResponseToPrice(&products, priceOutput)
	return products, nil
}

// Update price information
func (Request) Put(products Products) error {
	var price database.CurrentPrice
	price.ID = products.ID
	price.CurrencyCode = products.CurrentPrice.CurrencyCode
	price.Value = products.CurrentPrice.Value
	return iMongoDb.PutPriceInfo(price)
}

func mapDatabaseResponseToPrice(products *Products, currentPrice database.CurrentPrice) {
	products.CurrentPrice.Value = currentPrice.Value
	products.CurrentPrice.CurrencyCode = currentPrice.CurrencyCode
}

func mapApiResponseToProducts(products *Products, response api.Response) {
	products.ID = response.Product.Item.Tcin
	products.Dpci = response.Product.Item.Dpci
	products.ProductDescription = response.Product.Item.ProductDescription.Title
	products.Upc = response.Product.Item.Upc
}
