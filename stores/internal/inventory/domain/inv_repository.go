package domain

import "service/stores/case1/internal/inventory/models"

type InventoryRepository interface {
	GetAllInventoryDAO() ([]models.Inventory, error)
	FilterSingleObjctInventoryDAO(product string) (models.Inventory, error)
}
