package products

import (
	"errors"
	"github.com/jcbray/go-products/api"
	"github.com/jcbray/go-products/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

var request Request

func Test_Get(t *testing.T) {
	tests := []struct {
		name                string
		id                  string
		returnedID          string
		returnedDescription string
		returnedPrice       float64
		returnedError       error
	}{
		{
			name:                "Success",
			id:                  "123456",
			returnedID:          "123456",
			returnedDescription: api.FAKE_ITEM_DESCRIPTION,
			returnedPrice:       database.FAKE_ITEM_PRICE,
			returnedError:       nil,
		},
		{
			name:                "Item No Content",
			id:                  "123457",
			returnedID:          "",
			returnedDescription: "",
			returnedPrice:       database.FAKE_ITEM_PRICE,
			returnedError:       nil,
		},
		{
			name:                "Item Fatal Error",
			id:                  "123458",
			returnedID:          "",
			returnedDescription: "",
			returnedPrice:       0,
			returnedError:       errors.New("fatal item error"),
		},
		{
			name:                "Price No Content",
			id:                  "123459",
			returnedID:          "123459",
			returnedDescription: api.FAKE_ITEM_DESCRIPTION,
			returnedPrice:       0,
			returnedError:       nil,
		},
		{
			name:                "Price Fatal Error",
			id:                  "123460",
			returnedID:          "123460",
			returnedDescription: api.FAKE_ITEM_DESCRIPTION,
			returnedPrice:       0,
			returnedError:       nil,
		},
		{
			name:                "Price and Item No Content",
			id:                  "123461",
			returnedID:          "",
			returnedDescription: "",
			returnedPrice:       0,
			returnedError:       nil,
		},
		{
			name:                "Price and Item Fatal Error",
			id:                  "123462",
			returnedID:          "",
			returnedDescription: "",
			returnedPrice:       0,
			returnedError:       errors.New("fatal item error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iApiRequest = api.Mock{}
			iMongoDb = database.Mock{}
			result, err := request.Get(tt.id)
			assert.Equal(t, tt.returnedID, result.ID)
			assert.Equal(t, tt.returnedDescription, result.ProductDescription)
			assert.Equal(t, tt.returnedPrice, result.CurrentPrice.Value)
			if tt.returnedError != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_Put(t *testing.T) {
	tests := []struct {
		name          string
		id            string
		returnedError error
	}{
		{
			name:          "Success",
			id:            "123456",
			returnedError: nil,
		},
		{
			name:          "Item Not Found",
			id:            "123457",
			returnedError: errors.New("item not found"),
		},
		{
			name:          "Update Fatal Error",
			id:            "123458",
			returnedError: errors.New("fatal item error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iApiRequest = api.Mock{}
			var products Products
			products.ID = tt.id
			err := request.Put(products)
			if tt.returnedError != nil {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
