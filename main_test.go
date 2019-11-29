package main

import (
	"bytes"
	"encoding/json"
	"github.com/jcbray/go-products/products"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetSuccess(t *testing.T) {
	iProducts = products.Mock{}
	router := httprouter.New()
	router.GET("/products/:id", GetProductRequest)

	req, _ := http.NewRequest("GET", "/products/12345", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.NotNil(t, rr.Body)
}

func Test_GetFailure(t *testing.T) {
	iProducts = products.Mock{}
	router := httprouter.New()
	router.GET("/products/:id", GetProductRequest)

	req, _ := http.NewRequest("GET", "/products/12346", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusBadRequest)
	assert.NotNil(t, rr.Body)
}

func Test_PutSuccess(t *testing.T) {
	iProducts = products.Mock{}
	router := httprouter.New()
	router.PUT("/products/", InsertProductRequest)
	var product products.Products
	product.ID = "12345"
	request, _ := json.Marshal(product)
	req, _ := http.NewRequest("PUT", "/products/", bytes.NewBuffer(request))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.NotNil(t, rr.Body)
}

func Test_PutFailure(t *testing.T) {
	iProducts = products.Mock{}
	router := httprouter.New()
	router.PUT("/products/", InsertProductRequest)
	var product products.Products
	product.ID = "12346"
	request, _ := json.Marshal(product)
	req, _ := http.NewRequest("PUT", "/products/", bytes.NewBuffer(request))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusBadRequest)
	assert.NotNil(t, rr.Body)
}
