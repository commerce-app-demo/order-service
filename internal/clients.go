package internal

import (
	"context"
	"log"

	"github.com/commerce-app-demo/order-service/external/clients"
	"github.com/commerce-app-demo/order-service/internal/config"
	productspb "github.com/commerce-app-demo/product-service/proto"
	userspb "github.com/commerce-app-demo/user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitProductClient(ctx context.Context, serverConfig config.ServerConfig) *clients.ProductClient {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcClientProduct, err := grpc.NewClient(serverConfig.ProductGrpcPort, opts...)
	if err != nil {
		log.Fatalf("Error creating product client: %s", err)
	}
	productServiceClient := productspb.NewProductServiceClient(grpcClientProduct)
	productClient, err := clients.NewProductClient(ctx, &productServiceClient)
	if err != nil {
		log.Fatalf("Error creating grpc product client: %s", err)
	}
	return productClient
}

func InitUserClient(ctx context.Context, serverConfig config.ServerConfig) *clients.UserClient {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcClientUser, err := grpc.NewClient(serverConfig.UserGrpcPort, opts...)
	if err != nil {
		log.Fatalf("Error creating user client: %s", err)
	}
	userServiceClient := userspb.NewUserServiceClient(grpcClientUser)
	userClient, err := clients.NewUserClient(ctx, &userServiceClient)
	if err != nil {
		log.Fatalf("Error creating grpc user client: %s", err)
	}
	return userClient
}
