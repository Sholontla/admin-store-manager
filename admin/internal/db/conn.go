package db

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConn = ConnMongoDB()
var clientOptions = options.Client().ApplyURI("mongodb://servicedemo:servicedemo@servadmindb:27017/")

var mongoClientPool = &sync.Pool{
	New: func() interface{} {
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("MongoDB Connected Service Admin ...")
		return client
	},
}

func ConnMongoDB() *mongo.Client {
	client := mongoClientPool.Get().(*mongo.Client)
	return client
}

func ReleaseMongoDB(client *mongo.Client) {
	mongoClientPool.Put(client)
}
