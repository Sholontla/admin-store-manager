package stores

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

type F struct {
	Total           float64                `bson:"total" json:"total"`
	StoreInfomation models.StoreInfomation `bson:"store_information" json:"store_information"`
}

type M struct {
	Total            float64                            `bson:"total" json:"total"`
	StoreInfomation  models.StoreInfomation             `bson:"store_information" json:"store_information"`
	Employee         models.EmployeeInternalInformation `bson:"employee" json:"employee"`
	InventoryProduct []Product                          `bson:"inventory_Product" json:"inventory_Product"`
}

type Product struct {
	ID                    uuid.UUID `bson:"_id" json:"id"`
	ProductTitle          string    `bson:"product_title" json:"product_title"`
	ProductClassification string    `bson:"product_classification" json:"product_classification"`
	ProductCategory       string    `bson:"product_category" json:"product_category"`
	ProductBrand          string    `bson:"product_brand" json:"product_brand"`
	ProductModel          string    `bson:"product_model" json:"product_model"`
	ProductMaterial       string    `bson:"product_material" json:"product_material"`
	ProductDescription    string    `bson:"product_description" json:"product_description"`
	ProductImage          string    `bson:"product_image" json:"product_image"`
	ProductPrice          float64   `bson:"product_price" json:"product_price"`
	ProductQuantity       int64     `bson:"product_quantity" json:"product_quantity"`
	ProductSerialNumber   string    `bson:"product_serial_number" json:"product_serial_number"`
	ProductCreatedAt      string    `bson:"product_created_at" json:"product_created_at"`
	ProductUpdatedAt      string    `bson:"product_updated_at" json:"product_updated_at"`
	SupplierName          string    `bson:"supplier_name" json:"supplier_name"`
	Sku                   string    `bson:"sku" json:"sku"`
}
