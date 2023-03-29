package admin_order

import (
	"service/stores/case1/internal/inventory/models"
	"time"

	"github.com/google/uuid"
)

func RequestAdminInventoryOrder(orderRequest models.AdminOrderRequest) (models.AdminOrderRequest, error) {

	request := models.AdminOrderRequest{
		OrderId:          uuid.New(),
		OrderHash:        uuid.New().String(),
		CreatedAt:        time.Now().Local().String(),
		RequestInventory: orderRequest.RequestInventory,
	}

	return request, nil
}
