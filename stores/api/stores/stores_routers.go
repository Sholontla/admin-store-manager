package stores

import (
	db "service/stores/case1/datasource/mongodb"
	"service/stores/case1/internal/domain/order"
	"service/stores/case1/internal/stores"
	metrics "service/stores/case1/pkg/monitoring/middleware"

	"github.com/gofiber/fiber/v2"
)

func StoreRouters(r *fiber.App) {

	or := order.NewOrderTicket(db.ConnMongoDB())
	service := StoreHnadlerService{or}

	conOr := stores.NewCashierOrderCompeleted(db.ConnMongoDB())
	conService := ConcurrentCashierHnadlerService{conOr}

	ser := r.Group("service")
	store := ser.Group("store", metrics.RecordRequestLatency)

	store.Get("get/all", service.GetAllInventoryHandlers)
	store.Post("register/order", service.CashierTotalOrderRegistrationHandler)
	store.Post("concurrent/register/order", conService.ConcurrentCashierTotalOrderRegistrationHandler)
	store.Get("concurrent/get/all/inv", conService.ConcurrentGetAllInventoryHandler)
}
