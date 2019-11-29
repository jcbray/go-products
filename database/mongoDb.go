package database

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoDb interface {
	Connect() error
	GetPriceInfo(id string, currentPrice chan CurrentPrice)
	PutPriceInfo(price CurrentPrice) error
}

type Request struct{}

var ctx context.Context
var db *mongo.Database

const FATEL_PRICE_ERROR = "Fatal price error "

func (Request) Connect() error {

	ctx = context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//Connect to locally running database, will need to update to a managed instance when one gets stood up
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
	if err != nil {
		return fmt.Errorf("todo: couldn't connect to mongo: %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
	}
	db = client.Database("local")
	return nil
}

func health() bool {
	// Check the connection
	err := db.Client().Ping(ctx, nil)

	if err != nil {
		fmt.Println("database failed health check " + err.Error())
		return false
	}
	return true
}

// Get price information for given item from database
func (Request) GetPriceInfo(id string, currentPrice chan CurrentPrice) {
	if !health() {
		var result CurrentPrice
		result.Error = FATEL_PRICE_ERROR + "unable to connect to database"
		currentPrice <- result
		return
	}
	result, _ := findOne(id)
	currentPrice <- result
	return
}

// Update or insert price information for given item
func (Request) PutPriceInfo(price CurrentPrice) error {
	if !health() {
		return errors.New(FATEL_PRICE_ERROR + "unable to connect to database")

	}
	_, found := findOne(price.ID)
	if found {
		return updatePriceInfo(price)
	}
	return insertPriceInfo(price)
}

func findOne(id string) (CurrentPrice, bool) {
	var result CurrentPrice
	filter := bson.D{{"id", id}}
	db.Collection("prices")
	collection := db.Collection("prices")
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		result.Error = FATEL_PRICE_ERROR + err.Error()
		return result, false
	}
	return result, true
}

func insertPriceInfo(price CurrentPrice) error {
	var bdoc interface{}
	json, marshalError := json.Marshal(price)
	if marshalError != nil {
		return errors.New("unable to parse request")
	}
	err := bson.UnmarshalExtJSON(json, true, &bdoc)
	if err != nil {
		return errors.New("unable to convert request to bson")
	}

	collection := db.Collection("prices")

	insertResult, err := collection.InsertOne(ctx, bdoc)
	if err != nil {
		return errors.New("failed to insert into database")
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

func updatePriceInfo(price CurrentPrice) error {
	filter := bson.D{{"id", price.ID}}

	update := bson.D{{"$set",
		bson.D{
			{"Value", price.Value},
			{"CurrencyCode", price.CurrencyCode},
		},
	}}

	collection := db.Collection("prices")

	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.New("failed to update database")
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}
