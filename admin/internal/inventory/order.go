package inventory

import (
	"service/admin/case1/internal/inventory/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryOrder struct {
	ID                primitive.ObjectID      `bson:"_id" json:"id"`
	ProviderBuissness string                  `bson:"provider_buissness" json:"provider_buissness"`
	ProviderContact   models.ProviderContact  `bson:"provider_contact" json:"provider_contact"`
	ProviderAddress   models.ProviderAddress  `bson:"provider_address" json:"provider_address"`
	ProviderProduct   []OrderProviderProducts `bson:"order_provider_products" json:"order_provider_products"`
	CreatedAt         string                  `bson:"created_at" json:"created_at"`
	UserAdmin         string                  `bson:"user_admin" json:"user_admin"`
	OrderPayment      OrderPayment            `bson:"ordr_payment" json:"ordr_payment"`
}

type OrderPayment struct {
	TypePayment   string `bson:"type_payment" json:"type_payment"`
	DaysOfPayment string `bson:"days_of_payment" json:"days_of_payment"`
}

type OrderProviderProducts struct {
	ProductName           string  `bson:"product_name" json:"product_name"`
	ProductCategory       string  `bson:"product_ctegory" json:"product_ctegory"`
	ProductPrice          float64 `bson:"product_price" json:"product_price"`
	ProductSKU            string  `bson:"product_sku" json:"product_sku"`
	ProductMaterial       string  `bson:"product_material" json:"product_material"`
	ProductClassification string  `bson:"product_classification" json:"product_classification"`
	ProductQuantity       int     `bson:"product_quantity" json:"product_quantity"`
}
