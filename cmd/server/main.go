package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/commerce-app-demo/order-service/external/clients"
	"github.com/commerce-app-demo/order-service/internal/config"
	productspb "github.com/commerce-app-demo/product-service/proto"
	userspb "github.com/commerce-app-demo/user-service/proto"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var MAX_RETRY = 3

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	configFactory := config.Config{}
	dbConfig := configFactory.LoadDB()
	serverConfig := configFactory.LoadServer()

	db, err := sql.Open(dbConfig.Driver, dbConfig.GetDSN())

	if err != nil {
		log.Fatalf("failed to open db connection error: %s", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("ping returned error: %s\nPlease check whether database is not running", err)
	}

	// external calls
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	grpcClientProduct, err := grpc.NewClient(serverConfig.ProductGrpcPort, opts...)
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}

	defer grpcClientProduct.Close()

	productServiceClient := productspb.NewProductServiceClient(grpcClientProduct)
	productClient, err := clients.NewProductClient(ctx, &productServiceClient)

	if err != nil {
		log.Fatalf("Error creating grpc client: %s", err)
	}

	grpcClientUser, err := grpc.NewClient(serverConfig.UserGrpcPort, opts...)
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}
	defer grpcClientUser.Close()

	userServiceClient := userspb.NewUserServiceClient(grpcClientUser)
	userClient, err := clients.NewUserClient(ctx, &userServiceClient)

	if err != nil {
		log.Fatalf("Error creating grpc client: %s", err)
	}

	log.Printf("All dependencies have been initalized, %v\n%v\n%v\n", db, productClient, userClient)

}
