package main

import (
	"context"
	"log"
	"time"

	"github.com/commerce-app-demo/order-service/internal"
	"github.com/commerce-app-demo/order-service/internal/config"
	"github.com/commerce-app-demo/order-service/internal/repository/mysql"
	"github.com/commerce-app-demo/order-service/internal/service"
)

var MAX_RETRY = 3

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	configFactory := config.Config{}
	dbConfig := configFactory.LoadDB()
	serverConfig := configFactory.LoadServer()

	db := internal.InitDB(dbConfig)
	productClient := internal.InitProductClient(ctx, serverConfig)
	userClient := internal.InitUserClient(ctx, serverConfig)

	log.Printf("All dependencies have been initalized, %v\n%v\n%v\n", db, productClient, userClient)

	orderRepo := &mysql.OrderRepository{DB: db}
	orderService := &service.OrderService{Repo: orderRepo}
	internal.RunGRPCServer(orderService, ":50053") // or use serverConfig.Port
}
