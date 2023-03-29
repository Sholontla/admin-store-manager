package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	InventoryProduct []Product          `bson:"inventory_Product" json:"inventory_Product"`
	CreatedAt        string             `bson:"created_at" json:"created_at"`
	UpdatedAt        string             `bson:"updated_at" json:"updated_at"`
}
