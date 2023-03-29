package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConn = ConnMongoDB()
var clientOptions = options.Client().ApplyURI("mongodb://servicedemo:servicedemo@servadmindb:27017/").SetMaxPoolSize(200)

func ConnMongoDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("MongoDB Connected ...")
	return client
}

func CheckConection() int {
	err := MongoConn.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
