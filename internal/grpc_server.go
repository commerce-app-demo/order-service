package internal

import (
	"log"
	"net"

	"github.com/commerce-app-demo/order-service/internal/server"
	"github.com/commerce-app-demo/order-service/internal/service"
	orderspb "github.com/commerce-app-demo/order-service/proto"
	"google.golang.org/grpc"
)

func RunGRPCServer(orderService *service.OrderService, listenAddr string) {
	grpcServer := grpc.NewServer()
	orderspb.RegisterOrderServiceServer(grpcServer, &server.OrderServiceServer{OrderService: orderService})
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("OrderService gRPC server listening at %v", listenAddr)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
