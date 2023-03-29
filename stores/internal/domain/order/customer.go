package order

import (
	inv "service/stores/case1/internal/inventory/models"
	"service/stores/case1/internal/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerOrderCompleted struct {
	ID              primitive.ObjectID                 `bson:"_id" json:"id"`
	OrderReference  string                             `bson:"order_reference" json:"order_reference"`
	Inventory       inv.Inventory                      `bson:"inventory" json:"inventory"`
	StoreInfomation models.StoreInfomation             `bson:"store_information" json:"store_information"`
	Employee        models.EmployeeInternalInformation `bson:"employee" json:"employee"`
	CreatedAt       string                             `bson:"created_at" json:"created_at"`
	TicketHash      uuid.UUID                          `bson:"ticket_hash" json:"ticket_hash"`
}

type TicketCustomerOrder struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	TicketReference string             `bson:"ticket_reference" json:"ticket_reference"`
	Products        CustomerOrderCompleted
	Employee        CustomerOrderCompleted
	Store           CustomerOrderCompleted
	CreatedAt       string `bson:"created_at" json:"created_at"`
}

type Finance struct {
	ID              primitive.ObjectID     `bson:"_id" json:"id"`
	Inventory       inv.Inventory          `bson:"inventory" json:"inventory"`
	StoreInfomation models.StoreInfomation `bson:"store_information" json:"store_information"`
	CreatedAt       string                 `bson:"created_at" json:"created_at"`
}

type Marketing struct {
	ID              primitive.ObjectID                 `bson:"_id" json:"id"`
	Inventory       inv.Inventory                      `bson:"inventory" json:"inventory"`
	StoreInfomation models.StoreInfomation             `bson:"store_information" json:"store_information"`
	Employee        models.EmployeeInternalInformation `bson:"employee" json:"employee"`
	CreatedAt       string                             `bson:"created_at" json:"created_at"`
}

type Admin struct {
	ID              primitive.ObjectID     `bson:"_id" json:"id"`
	Inventory       inv.Inventory          `bson:"inventory" json:"inventory"`
	StoreInfomation models.StoreInfomation `bson:"store_information" json:"store_information"`
	CreatedAt       string                 `bson:"created_at" json:"created_at"`
}

type Accounting struct {
	ID              primitive.ObjectID                 `bson:"_id" json:"id"`
	Inventory       inv.Inventory                      `bson:"inventory" json:"inventory"`
	StoreInfomation models.StoreInfomation             `bson:"store_information" json:"store_information"`
	Employee        models.EmployeeInternalInformation `bson:"employee" json:"employee"`
	CreatedAt       string                             `bson:"created_at" json:"created_at"`
}
