package order

import (
	"service/stores/case1/internal/models"
)

type IPoolService interface {
	CreateOrderDAO(orderRequestChan <-chan models.CustomerOrderCompleted) Ticket
}
