package service

import (
	"github.com/commerce-app-demo/order-service/internal/models/orders"
	"github.com/commerce-app-demo/order-service/internal/repository/mysql"
)

type OrderService struct {
	Repo *mysql.OrderRepository
}

func (s *OrderService) GetOrders() ([]orders.OrderEntity, error) {
	return s.Repo.Orders()
}

func (s *OrderService) GetOrderById(id int) (*orders.OrderEntity, error) {
	return s.Repo.OrderById(id)
}

func (s *OrderService) CreateOrder(order *orders.OrderEntity) (*orders.OrderEntity, error) {
	return s.Repo.CreateOrder(order)
}

func (s *OrderService) UpdateOrder(id int, order *orders.OrderEntity) (*orders.OrderEntity, error) {
	updatedFields := make(map[string]any)
	if order.UserId != 0 {
		updatedFields["user_id"] = order.UserId
	}
	if order.TotalPrice != 0 {
		updatedFields["total_price"] = order.TotalPrice
	}
	if order.Status != "" {
		updatedFields["status"] = order.Status
	}
	return s.Repo.UpdateOrder(id, updatedFields)
}

func (s *OrderService) DeleteOrder(id int) (*orders.OrderEntity, error) {
	return s.Repo.DeleteOrder(id)
}

// Order Items
func (s *OrderService) GetOrderItems(orderId int) ([]orders.OrderItemEntity, error) {
	return s.Repo.OrderItemsByOrderId(orderId)
}

func (s *OrderService) CreateOrderItem(item *orders.OrderItemEntity) (*orders.OrderItemEntity, error) {
	return s.Repo.CreateOrderItem(item)
}

func (s *OrderService) UpdateOrderItem(id int, item *orders.OrderItemEntity) (*orders.OrderItemEntity, error) {
	updatedFields := make(map[string]any)
	if item.OrderId != 0 {
		updatedFields["order_id"] = item.OrderId
	}
	if item.ProductId != 0 {
		updatedFields["product_id"] = item.ProductId
	}
	if item.Quantity != 0 {
		updatedFields["quantity"] = item.Quantity
	}
	if item.Price != 0 {
		updatedFields["price"] = item.Price
	}
	return s.Repo.UpdateOrderItem(id, updatedFields)
}

func (s *OrderService) DeleteOrderItem(id int) (*orders.OrderItemEntity, error) {
	return s.Repo.DeleteOrderItem(id)
}
