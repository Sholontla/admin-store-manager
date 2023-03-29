package grpc

import (
	"context"
	"io"
	"log"
	"service/stores/case1/internal/inventory/models"
	"service/stores/case1/internal/proto/admin_orders"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func StreamRequestStoreInventory(c admin_orders.RequestOrderStoreInventoryServiceClient, requestChan <-chan models.AdminOrderRequest) {

	log.Println("Service stream request order store inventory ....")

	stream, err := c.StreamOrderStoreInventory(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}
	waitc := make(chan struct{})
	go func() {
		for request := range requestChan {

			inventory := make([]*admin_orders.RequestInventory, len(request.RequestInventory))
			for i, item := range request.RequestInventory {
				inventory[i] = &admin_orders.RequestInventory{
					Sku:      item.Sku,
					Quantity: item.Quantity,
				}
			}

			req := &admin_orders.RequestOrderStoreInventory{
				OrderId:   uuid.New().String(),
				OrderHash: uuid.New().String(),
				CreatedAt: timestamppb.Now().String(),
				Inventory: inventory,
			}

			log.Printf("Send Request: %v\n", req)
			err := stream.Send(req)
			if err != nil {
				log.Fatalf("Error while sending request: %v\n", err)
			}
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while reading stream: %v\n", err)
			}
			log.Println("Recived: ", res)
		}
		close(waitc)
	}()

	<-waitc
}
