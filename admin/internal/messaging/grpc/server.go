package grpc

import "service/admin/case1/internal/proto/admin_orders"

type Server struct {
	admin_orders.RequestOrderStoreInventoryServiceServer
}
