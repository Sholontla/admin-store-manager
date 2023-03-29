package models

import "github.com/google/uuid"

type AdminOrderRequest struct {
	OrderId          uuid.UUID          `json:"order_id"`
	OrderHash        string             `json:"order_hash"`
	CreatedAt        string             `json:"created_at"`
	RequestInventory []RequestInventory `json:"inventory"`
}

type RequestInventory struct {
	Sku      string `json:"sku"`
	Quantity int64  `json:"quantity"`
}
