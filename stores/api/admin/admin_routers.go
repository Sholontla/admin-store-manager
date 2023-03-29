package admin

import (
	metrics "service/stores/case1/pkg/monitoring/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRouters(r *fiber.App) {

	ser := r.Group("service")
	store := ser.Group("store", metrics.RecordRequestLatency)

	store.Post("request/order", RequestOrderStoreInventoryHandler)
	store.Get("request/inventory", GetInventory)

}
