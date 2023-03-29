package models

import "github.com/gocql/gocql"

type ProductRotation struct {
	ID                    gocql.UUID `bson:"_id" json:"id"`
	ProductTitle          string     `bson:"product_title" json:"product_title"`
	ProductClassification string     `bson:"product_classification" json:"product_classification"`
	ProductCategory       string     `bson:"product_category" json:"product_category"`
	ProductBrand          string     `bson:"product_brand" json:"product_brand"`
	ProductModel          string     `bson:"product_model" json:"product_model"`
	ProductMaterial       string     `bson:"product_material" json:"product_material"`
	ProductDescription    string     `bson:"product_description" json:"product_description"`
	ProductImage          string     `bson:"product_image" json:"product_image"`
	ProductPrice          float64    `bson:"product_price" json:"product_price"`
	ProductQuantity       int64      `bson:"product_quantity" json:"product_quantity"`
	ProductSerialNumber   string     `bson:"product_serial_number" json:"product_serial_number"`
	ProductCreatedAt      string     `bson:"product_created_at" json:"product_created_at"`
	ProductUpdatedAt      string     `bson:"product_updated_at" json:"product_updated_at"`
	SupplierName          string     `bson:"supplier_name" json:"supplier_name"`
	Sku                   string     `bson:"sku" json:"sku"`
}
