package db

import (
	"context"
	"fmt"
	"log"
	"service/admin/case1/internal/inventory/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dataBase = "admin_db"
)

func InsertOne(client *mongo.Client, request interface{}, collection string) *mongo.InsertOneResult {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	conn := client.Database(dataBase)
	col := conn.Collection(collection)
	response, errInsert := col.InsertOne(ctx, request)

	if errInsert != nil {
		log.Println(errInsert)
	}

	return response
}

func UpdateOne(client *mongo.Client, provider, update map[string]interface{}, collection string) (*mongo.UpdateResult, error) {
	filter := bson.M{"provider_buissness": provider, "productSKU": "123"}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
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

func UpdateInventory(client *mongo.Client, collection string) ([]map[string]interface{}, error) {
	var responseSlice []map[string]interface{}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	conn := client.Database(dataBase)
	coll := conn.Collection(collection)

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

func GetAll(client *mongo.Client, collection string) ([]models.ProviderInformation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var responseSlice []models.ProviderInformation

	conn := client.Database(dataBase)
	coll := conn.Collection(collection)

	cursor, err := coll.Find(ctx, bson.D{})
	//defer cancel()
	if err != nil {
		log.Println("..........", err)
	}

	for cursor.Next(ctx) {
		var result models.ProviderInformation
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		fmt.Println(result)
		responseSlice = append(responseSlice, result)
	}

	return responseSlice, nil
}

func UpdateMany(client *mongo.Client, providerBusiness string, productSKUs string, productQuantities []int, collection string) (*mongo.UpdateResult, error) {
	updates := []interface{}{}

	filter := bson.M{"provider_business": providerBusiness, "provider.product_sku": productSKUs}
	update := bson.M{
		"$addToSet": bson.M{
			"provider.$[elem].product_quantity": bson.M{"$each": productQuantities},
		},
	}
	opts := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.M{"elem.product_sku": productSKUs}},
	})
	updates = append(updates, filter, update, opts)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	conn := client.Database(dataBase)
	col := conn.Collection(collection)
	response, errUpdate := col.UpdateMany(ctx, bson.D{}, updates)
	if errUpdate != nil {
		log.Println(errUpdate)
		return nil, errUpdate
	}

	return response, nil
}
