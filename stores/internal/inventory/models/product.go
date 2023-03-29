package models

import (
	"github.com/google/uuid"
)

type ClassificationProductsGroups struct {
	PriceStartA float64 `json:"price_start_a"`
	PriceEndA   float64 `json:"price_end_a"`
	PriceStartB float64 `json:"price_start_b"`
	PriceEncql  float64 `json:"price_end_b"`
	PriceStartC float64 `json:"price_start_c"`
	PriceEndC   float64 `json:"price_end_c"`
	PriceStartD float64 `json:"price_start_d"`
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

type ProductGrpcLog struct {
	ProductId             uuid.UUID `json:"product_id"`
	ProductTitle          string    `json:"product_title"`
	ProductClassification string    `json:"product_classification"`
	ProductCategory       string    `json:"product_category"`
	ProductBrand          string    `json:"product_brand"`
	ProductModel          string    `json:"product_model"`
	ProductMaterial       string    `json:"product_material"`
	ProductDescription    string    `json:"product_description"`
	ProductImage          string    `json:"product_image"`
	ProductPrice          float64   `json:"product_price"`
	ProductQuantity       int64     `json:"product_quantity"`
	ProductSerialNumber   string    `json:"product_serial_number"`
	ProductCreatedAt      string    `json:"product_created_at"`
	ProductUpdatedAt      string    `json:"product_updated_at"`
	SupplierName          string    `json:"supplier_name"`
	Sku                   string    `json:"sku"`
}
