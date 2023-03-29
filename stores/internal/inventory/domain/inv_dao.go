package domain

import (
	"context"
	db "service/stores/case1/datasource/mongodb"
	"service/stores/case1/internal/inventory/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	dataBase   = "stores"
	collection = "stores-a"
)

func GetAllInventoryDAO() ([]models.Inventory, error) {
	var out []models.Inventory

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn := db.MongoConn.Database(dataBase)
	coll := conn.Collection(collection)

	cursor, err := coll.Find(ctx, out)
	if err != nil {
		return out, nil
	}

	for cursor.Next(ctx) {
		var result models.Inventory
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		out = append(out, result)
	}

	return out, nil

}

func FilterSingleObjctInventoryDAO(product string) (models.Inventory, error) {

	var inv models.Inventory
	condition := bson.M{
		"product": product,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	conn := db.MongoConn.Database(dataBase)
	coll := conn.Collection(collection)
	err := coll.FindOne(ctx, condition).Decode(&inv)
	if err != nil {
		return inv, err
	}

	return inv, nil
}
