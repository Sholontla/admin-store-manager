package stores

import (
	"fmt"
	"log"
	"reflect"
	"service/stores/case1/datasource/cassandra"
	"service/stores/case1/internal/domain"
	"service/stores/case1/internal/domain/order"
	inv "service/stores/case1/internal/inventory/domain"
	"service/stores/case1/internal/models"
	"service/stores/case1/pkg/concurrency"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var wg sync.WaitGroup

type StoreHnadlerService struct {
	Order order.IPoolService
}

func (s StoreHnadlerService) GetAllInventoryHandlers(ctx *fiber.Ctx) error {

	inv, err := inv.GetAllInventoryDAO()
	if err != nil {
		log.Println(err)
		return ctx.Status(500).SendString(err.Error())
	}

	return ctx.JSON(fiber.Map{"inventory": inv})
}

func (s StoreHnadlerService) CashierTotalOrderRegistrationHandler(ctx *fiber.Ctx) error {
	var orderRequest models.CustomerOrderCompleted

	orderPool, err := concurrency.GenericPool(reflect.TypeOf(orderRequest))
	if err != nil {
		log.Println(err)
	}

	orderRequestChan := orderPool.Get().(chan models.CustomerOrderCompleted)
	defer orderPool.Put(orderRequestChan)

	if err := ctx.BodyParser(&orderRequest); err != nil {
		ctx.JSON(fiber.Map{"error": "invalid JSON"})
		return err
	}
	var t []order.Ticket

	wg.Add(1)
	go func() {
		defer wg.Done()
		ti := s.Order.CreateOrderDAO(orderRequestChan)
		t = append(t, ti)
	}()

	orderRequestChan <- orderRequest
	close(orderRequestChan)

	wg.Wait()

	return ctx.JSON(fiber.Map{"ticket_created": t})
}

func StoreOrderCanceletaionProductHnadler(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{"message": "We Succeed !!!"})
}

func StoreOrderCanceletaionHnadler(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{"message": "We Succeed !!!"})
}

func StoreOrderUpdateHnadler(ctx *fiber.Ctx) error {
	db, err := cassandra.NewCassandraDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	r, _ := domain.GetAllStoreLogins(db)
	fmt.Println(r)
	email := "admin@example.com"
	resultChan := make(chan *models.StoreLogin)
	errorChan := make(chan error)

	go domain.GetStoreLogin(db.(*cassandra.CassandraDB).Session, email, resultChan, errorChan)

	select {
	case storeLogin := <-resultChan:
		if storeLogin != nil {
			fmt.Println(storeLogin)
		} else {
			fmt.Printf("Store login with email %s not found\n", email)
		}
	case err := <-errorChan:
		log.Fatal(err)
	}
	return ctx.JSON(fiber.Map{"message": "We Succeed !!!"})
}
