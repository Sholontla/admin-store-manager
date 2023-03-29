package inventory

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryInformation struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	ProviderBusniess string             `bson:"provider_busniess" json:"provider_busniess"`
	ProviderProducts []ProviderProducts `bson:"provider" json:"provider"`
	CreatedAt        string             `bson:"created_at" json:"created_at"`
	UserAdmin        string             `bson:"user_admin" json:"user_admin"`
}

type InventorySearch struct {
	BusinessName string `bson:"business_name" json:"business_name"`
	ProductName  string `bson:"product_name" json:"product_name"`
}

type InventoryQuantity struct {
	ProductSku string `bson:"product_sku" json:"product_sku"`
	Quantity   int    `bson:"quantity" json:"quantity"`
}

type ProviderProducts struct {
	ProductName           string  `bson:"product_name" json:"product_name"`
	ProductCategory       string  `bson:"product_ctegory" json:"product_ctegory"`
	ProductPrice          float64 `bson:"product_price" json:"product_price"`
	ProductSKU            string  `bson:"product_sku" json:"product_sku"`
	ProductMaterial       string  `bson:"product_material" json:"product_material"`
	ProductClassification string  `bson:"product_classification" json:"product_classification"`
	ProductQuantity       int     `bson:"product_quantity" json:"product_quantity"`
}
