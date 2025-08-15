package service

import (
	"context"

	"github.com/commerce-app-demo/order-service/internal/models/orders"
)

type OrderEventProducer interface {
	PublishOrderPlaced(ctx context.Context, order *orders.OrderEntity) error
}
