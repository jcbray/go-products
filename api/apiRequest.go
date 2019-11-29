package api

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type IApiRequest interface {
	GetItemInfo(id string, item chan Response)
}

type Request struct{}

// Configure http client.  Will need increase max idle connections when the tps gets past 1000
var tr = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, MaxIdleConns: 1000, MaxIdleConnsPerHost: 1000}
var client = &http.Client{Transport: tr, Timeout: time.Millisecond * 2000}

var baseUrl = "https://redsky.target.com/v2/pdp/tcin/"
var optionalExcludes = "?excludes=taxonomy,price,promotion,bulk_ship,rating_and_review_reviews,rating_and_review_statistics,available_to_promise_network,deep_red_labels,question_answer_statistics"

const FATEL_ITEM_ERROR = "Fatal item error "

func generateUrl(id string) string {
	return baseUrl + id + optionalExcludes
}

func isValidStatusCode(status int) bool {
	return status == http.StatusOK || status == http.StatusCreated || status == http.StatusNoContent
}

// Get item information from Redsky api
func (Request) GetItemInfo(id string, item chan Response) {
	var response Response
	httpResponse, responseErr := client.Get(generateUrl(id))
	if responseErr != nil {
		response.Error = errors.New(FATEL_ITEM_ERROR + responseErr.Error())
		item <- response
		return
	}
	if !isValidStatusCode(httpResponse.StatusCode) {
		response.Error = errors.New(FATEL_ITEM_ERROR + "api request failed")
		item <- response
		return
	}
	body, readError := ioutil.ReadAll(httpResponse.Body)
	if readError != nil {
		response.Error = errors.New(FATEL_ITEM_ERROR + readError.Error())
		item <- response
		return
	}
	io.Copy(ioutil.Discard, httpResponse.Body)
	httpResponse.Body.Close()
	json.Unmarshal(body, &response)
	item <- response
	return
}
