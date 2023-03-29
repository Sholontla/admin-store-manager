package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dataBase            = "stores"
	collection          = "stores-conncurrent-a"
	collectionInventory = "stores-inventory-a"
)

var (
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
)

func InsertOne(client *mongo.Client, request map[string]interface{}) *mongo.InsertOneResult {

	conn := client.Database(dataBase)
	col := conn.Collection(collection)
	response, errInsert := col.InsertOne(ctx, request)
	defer cancel()
	if errInsert != nil {
		log.Println(errInsert)
	}

	return response
}

func UpdateOne(client *mongo.Client, filter map[string]interface{}, update map[string]interface{}) (*mongo.UpdateResult, error) {

	conn := client.Database(dataBase)
	col := conn.Collection(collection)
	response, errUpdate := col.UpdateOne(ctx, filter, update)
	defer cancel()
	if errUpdate != nil {
		log.Println(errUpdate)
		return nil, errUpdate
	}

	return response, nil
}

func UpdateInventory(client *mongo.Client) ([]map[string]interface{}, error) {
	var responseSlice []map[string]interface{}

	conn := client.Database(dataBase)
	coll := conn.Collection(collectionInventory)

	cursor, err := coll.Find(ctx, responseSlice)
	if err != nil {
		log.Println(err)
	}

	for cursor.Next(ctx) {
		result := make(map[string]interface{})
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		responseSlice = append(responseSlice, result)
	}
	defer cancel()

	return responseSlice, nil
}

func GetAll(client *mongo.Client) ([]map[string]interface{}, error) {
	var responseSlice []map[string]interface{}

	conn := client.Database(dataBase)
	coll := conn.Collection(collectionInventory)

	cursor, err := coll.Find(ctx, responseSlice)
	if err != nil {
		log.Println(err)
	}

	for cursor.Next(ctx) {
		result := make(map[string]interface{})
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		responseSlice = append(responseSlice, result)
	}
	defer cancel()

	return responseSlice, nil
}
