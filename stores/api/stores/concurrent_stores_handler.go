package stores

import (
	"fmt"
	"log"
	"service/stores/case1/pkg/pool"
	"time"

	"service/stores/case1/internal/stores"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ConcurrentCashierHnadlerService struct {
	Chashier stores.CashierOrderMainService
}

func (s ConcurrentCashierHnadlerService) ConcurrentCashierTotalOrderRegistrationHandler(ctx *fiber.Ctx) error {
	orderRequest := make(map[string]interface{})
	chanOrderRequest := make(chan map[string]interface{})
	if err := ctx.BodyParser(&orderRequest); err != nil {
		ctx.JSON(fiber.Map{"error": "invalid JSON"})
		return err
	}

	orderRequest["ticket_hash"] = uuid.New()
	orderRequest["created_at"] = time.Now().String()
	go func() {
		response, err := pool.WorkerPool(2, chanOrderRequest, s.Chashier.CashierOrderCompletedDAO)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(response)
	}()
	chanOrderRequest <- orderRequest
	close(chanOrderRequest)
	return ctx.JSON(fiber.Map{"ticket_created": "somthing"})
}

func (s ConcurrentCashierHnadlerService) ConcurrentGetAllInventoryHandler(ctx *fiber.Ctx) error {
	res, err := s.Chashier.GetAllInventoryDAO()
	if err != nil {
		ctx.JSON(fiber.Map{"error": err.Error()})
		return err
	}
	return ctx.JSON(fiber.Map{"ticket_created": res})
}
