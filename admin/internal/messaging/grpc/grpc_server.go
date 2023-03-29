package grpc

import (
	"fmt"
	"log"
	"net"
	"service/admin/case1/internal/config"
	"service/admin/case1/internal/proto/admin_orders"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func GrpcProductServer() {
	grpc_host, grpc_port, grpc_tls := config.GRPCConfig1()
	certFile, keyFile := config.GRPCPathsConfig()

	add := fmt.Sprintf("%s:%s", grpc_host, grpc_port)

	lis, err := net.Listen("tcp", add)

	if err != nil {
		log.Fatalf("Failed to listen on %s\n: ", err)
	}

	log.Printf("GRPC Server Listening on %s\n: ", add)

	opts := []grpc.ServerOption{}

	if grpc_tls {
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Error while loading cert files: %v\n", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	admin_orders.RegisterRequestOrderStoreInventoryServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve on %s\n: ", err)
	}
}
