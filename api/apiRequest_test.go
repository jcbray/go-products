package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const FAKE_ITEM = "13860427"

var request Request

func Test_GetItemInfoSuccessfulResponse(t *testing.T) {
	// Start a local HTTP server
	var server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		_, _ = rw.Write([]byte(`
		{ 
   		"product":{
      		"item":{ 
         		"tcin":"13860427",
         		"dpci":"058-34-0436",
         		"upc":"025192110306",
         		"product_description":{ 
            		"title":"The Big Lebowski (Blu-ray)",
            		"bullet_description":[ 
               			"<B>Movie Studio:</B> Universal Studios",
               			"<B>Movie Genre:</B> Comedy",
               			"<B>Software Format:</B> Blu-ray"
            		]
         		},
         		"recall_compliance":{ 
            		"is_product_recalled":false
         		},
         		"tax_category":{ 
            		"tax_class":"G",
            		"tax_code_id":99999,
            		"tax_code":"99999"
         		},
         		"fulfillment":{ 
            		"is_po_box_prohibited":true,
            		"po_box_prohibited_message":"We regret that this item cannot be shipped to PO Boxes.",
            		"box_percent_filled_by_volume":0.27,
            		"box_percent_filled_by_weight":0.43,
            		"box_percent_filled_display":0.43
         		},
         		"product_vendors":[ 
            		{ 
               			"id":"1984811",
               			"manufacturer_style":"025192110306",
               			"vendor_name":"Ingram Entertainment"
            		},
            		{ 
               			"id":"4667999",
               			"manufacturer_style":"61119422",
               			"vendor_name":"UNIVERSAL HOME VIDEO"
            		},
            		{ 
               			"id":"1979650",
               			"manufacturer_style":"61119422",
               			"vendor_name":"Universal Home Ent PFS"
            		}
         		],
         		"product_classification":{ 
            		"product_type":"542",
            		"product_type_name":"ELECTRONICS",
            		"item_type_name":"Movies",
            		"item_type":{ 
               		"category_type":"Item Type: MMBV",
               		"type":300752,
               		"name":"movies"
            		}
         		},
         		"product_brand":{ 
            		"brand":"Universal Home Video",
            		"manufacturer_brand":"Universal Home Video",
            		"facet_id":"55zki"
         		},
         		"item_state":"READY_FOR_LAUNCH",
         		"specifications":[],
         		"attributes":{ 
            		"gift_wrapable":"N",
            		"has_prop65":"N",
            		"is_hazmat":"N",
            		"manufacturing_brand":"Universal Home Video",
            		"max_order_qty":10,
            		"street_date":"2011-11-15",
            		"media_format":"Blu-ray",
           		 	"merch_class":"MOVIES",
            		"merch_classid":58,
            		"merch_subclass":34,
            		"return_method":"This item can be returned to any Target store or Target.com.",
            		"ship_to_restriction":"United States Minor Outlying Islands,American Samoa (see also separate entry under AS),Puerto Rico (see also separate entry under PR),Northern Mariana Islands,Virgin Islands, U.S.,APO/FPO,Guam (see also separate entry under GU)"
         		},
         		"country_of_origin":"US",
         		"relationship_type_code":"Stand Alone",
         		"subscription_eligible":false,
         		"ribbons":[],
         		"tags":[],
         		"ship_to_restriction":"This item cannot be shipped to the following locations: United States Minor Outlying Islands, American Samoa, Puerto Rico, Northern Mariana Islands, Virgin Islands, U.S., APO/FPO, Guam",
         		"estore_item_status_code":"A",
         		"is_proposition_65":false,
         		"return_policies":{ 
            		"user":"Regular Guest",
           		 "policyDays":"30",
            		"guestMessage":"This item must be returned within 30 days of the ship date. See return policy for details."
         		},
         		"gifting_enabled":false,
         		"packaging":{ 
            		"is_retail_ticketed":false
         		}
      		}
   		}
	}
		  `))
	})) // Close the server when test finishes
	defer server.Close()
	client = server.Client()
	baseUrl = server.URL + "/"
	optionalExcludes = ""
	itemOutputChannel := make(chan Response, 100)
	go request.GetItemInfo(FAKE_ITEM, itemOutputChannel)

	itemOutput := <-itemOutputChannel
	close(itemOutputChannel)

	assert.Equal(t, FAKE_ITEM, itemOutput.Product.Item.Tcin)
	assert.Equal(t, "The Big Lebowski (Blu-ray)", itemOutput.Product.Item.ProductDescription.Title)
	assert.Nil(t, itemOutput.Error)
}

func Test_GetItemInfoBadResponse(t *testing.T) {
	// Start a local HTTP server
	var server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.WriteHeader(http.StatusFailedDependency)
		rw.Write([]byte("something"))
	})) // Close the server when test finishes
	defer server.Close()
	client = server.Client()
	baseUrl = server.URL + "/"
	optionalExcludes = ""
	itemOutputChannel := make(chan Response, 100)
	go request.GetItemInfo(FAKE_ITEM, itemOutputChannel)

	itemOutput := <-itemOutputChannel
	close(itemOutputChannel)

	assert.Equal(t, "", itemOutput.Product.Item.Tcin)
	assert.Equal(t, "", itemOutput.Product.Item.ProductDescription.Title)
	assert.NotNil(t, itemOutput.Error)
}

func Test_GetItemInfoNotFound(t *testing.T) {
	// Start a local HTTP server
	var server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.WriteHeader(http.StatusNoContent)
	})) // Close the server when test finishes
	defer server.Close()
	client = server.Client()
	baseUrl = server.URL + "/"
	optionalExcludes = ""
	itemOutputChannel := make(chan Response, 100)
	go request.GetItemInfo(FAKE_ITEM, itemOutputChannel)

	itemOutput := <-itemOutputChannel
	close(itemOutputChannel)

	assert.Equal(t, "", itemOutput.Product.Item.Tcin)
	assert.Equal(t, "", itemOutput.Product.Item.ProductDescription.Title)
	assert.Nil(t, itemOutput.Error)
}

func Test_GetItemInfoServerFailure(t *testing.T) {
	// Start a local HTTP server
	var server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.WriteHeader(http.StatusInternalServerError)
	})) // Close the server when test finishes
	defer server.Close()
	client = server.Client()
	baseUrl = server.URL + "/"
	optionalExcludes = ""
	itemOutputChannel := make(chan Response, 100)
	go request.GetItemInfo(FAKE_ITEM, itemOutputChannel)

	itemOutput := <-itemOutputChannel
	close(itemOutputChannel)

	assert.Equal(t, "", itemOutput.Product.Item.Tcin)
	assert.Equal(t, "", itemOutput.Product.Item.ProductDescription.Title)
	assert.NotNil(t, itemOutput.Error)
}
