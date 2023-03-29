package grpc

import (
	"fmt"
	"io"
	"log"
	"service/admin/case1/internal/proto/admin_orders"
)

func (s *Server) StreamOrderStoreInventory(stream admin_orders.RequestOrderStoreInventoryService_StreamOrderStoreInventoryServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		fmt.Println("Comming from StreamInventory req: ", req)

		err = stream.Send(&admin_orders.ResponseOrderStoreInventory{
			OrderId:   "request.OrderId.String()",
			OrderHash: "request.OrderHash",
			CreatedAt: "request.CreatedAt",
			Inventory: []*admin_orders.ResponseInventory{
				{
					Sku:      " request.RequestInventory.Sku",
					Quantity: 12,
				},
				{
					Sku:      " request.RequestInventory.Sku",
					Quantity: 12,
				},
			},
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}

}
