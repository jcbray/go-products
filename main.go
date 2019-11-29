package main

import (
	"encoding/json"
	"github.com/jcbray/go-products/database"
	"github.com/jcbray/go-products/products"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const SERVER_PORT = ":8080"
const INSERT_ERROR_CODE = "INSERT_ERROR"
const GET_ERROR_CODE = "GET_ERROR"

type App struct {
	Router *httprouter.Router
}

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func initializeRoutes(a *App) {
	a.Router.PUT("/products/v1/", InsertProductRequest)
	a.Router.GET("/products/v1/:id", GetProductRequest)
	a.Router.GET("/products/health", health)
	//a.Router.GET("/products/v1/api-spec", apiSpec)
}

//ToDo: Add api spec with swagger endpoint to view it
/*func apiSpec(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte(SWAGGER_JSON))
}*/

var iProducts products.IProducts = products.Request{}

func health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	json.NewEncoder(w).Encode("Up")
}

func isValidProductRequest(product products.Products) bool {
	return isValidID(product.ID)
}

func isValidID(id string) bool {
	_, err := strconv.ParseFloat(id, 64)
	return err == nil
}

func InsertProductRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var product products.Products
	jsonDecodeErr := json.NewDecoder(r.Body).Decode(&product)

	if jsonDecodeErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: jsonDecodeErr.Error(), Code: INSERT_ERROR_CODE})
		return
	}
	if !isValidProductRequest(product) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "invalid input", Code: INSERT_ERROR_CODE})
		return
	}
	err := iProducts.Put(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error(), Code: INSERT_ERROR_CODE})
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetProductRequest(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	if !isValidID(id) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: "invalid input", Code: GET_ERROR_CODE})
		return
	}
	product, err := iProducts.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Message: err.Error(), Code: GET_ERROR_CODE})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func main() {
	// Establish connection to database, if unable to connect, fail to start
	var iMongoDb database.IMongoDb = database.Request{}
	err := iMongoDb.Connect()
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	app := initializeRouter()
	s := &http.Server{
		Addr:    SERVER_PORT,
		Handler: app.Router,
	}
	log.Fatal(s.ListenAndServe())
}

func initializeRouter() *App {
	app := &App{Router: httprouter.New()}
	if app.Router != nil {
		initializeRoutes(app)
	}
	return app
}
