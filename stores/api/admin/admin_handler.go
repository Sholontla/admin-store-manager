package admin

import (
	"fmt"
	"log"
	"service/stores/case1/config"
	"service/stores/case1/internal/domain/inventory"
	"service/stores/case1/internal/inventory/models"
	client "service/stores/case1/internal/messaging/grpc"
	"service/stores/case1/internal/proto/admin_orders"
	"sync"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var wg sync.WaitGroup

func RequestOrderStoreInventoryHandler(ctx *fiber.Ctx) error {

	var orderRequest models.AdminOrderRequest

	if err := ctx.BodyParser(&orderRequest); err != nil {
		log.Println("Invalid JSON BDOY", err)
	}

	grpc_host, grpc_port, _ := config.GRPCConfig1()
	add := fmt.Sprintf("%s:%s", grpc_host, grpc_port)

	// cert_file := config.GRPCPathsConfig()

	// opts := []grpc.DialOption{}

	// // if grpc_tls {
	// // 	creds, err := credentials.NewClientTLSFromFile(cert_file, add)
	// // 	if err != nil {
	// // 		log.Fatalf("Error While loading client cert file: %v\n", err)
	// // 	}
	// // 	opts = append(opts, grpc.WithTransportCredentials(creds))
	// // }

	//conn, err := grpc.Dial(add, opts...)

	conn, err := grpc.Dial(add, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect Client: %v\n", err)
	}
	defer conn.Close()

	log.Printf("GRPC Client Listening on %s\n: ", "50051")

	c := admin_orders.NewRequestOrderStoreInventoryServiceClient(conn)
	requestChan := make(chan models.AdminOrderRequest)

	wg.Add(1)
	go func() {
		defer wg.Done()
		client.StreamRequestStoreInventory(c, requestChan)
	}()
	requestChan <- orderRequest
	close(requestChan)

	wg.Wait()
	return ctx.JSON(fiber.Map{"store": orderRequest})
}

func GetInventory(ctx *fiber.Ctx) error {

	inv := <-inventory.GetAllInventoryDAO()

	return ctx.JSON(fiber.Map{"store": inv})
}
