package inventory

import (
	"context"
	"log"
	db "service/stores/case1/datasource/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllInventoryDAO() chan InventoryInformation {
	invInfoChan := make(chan InventoryInformation)
	mongoCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn := db.MongoConn.Database("admin_db")
	collInv := conn.Collection("inventory")

	go func() {
		defer close(invInfoChan)
		cursor, err := collInv.Find(mongoCtx, bson.M{})
		if err != nil {
			log.Println("Error while interst inventory order ...", err.Error())
			return
		}
		defer cursor.Close(mongoCtx)
		for cursor.Next(mongoCtx) {
			var invInfo InventoryInformation
			if err := cursor.Decode(&invInfo); err != nil {
				log.Println("Error decoding inventory order ...", err.Error())
				continue
			}
			invInfoChan <- invInfo
		}
		if err := cursor.Err(); err != nil {
			log.Println("Error while iterating inventory order cursor ...", err.Error())
		}
	}()

	return invInfoChan
}
