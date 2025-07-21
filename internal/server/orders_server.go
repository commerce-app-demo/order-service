package server

import (
	"context"
	"fmt"

	"github.com/commerce-app-demo/order-service/external/clients"
	"github.com/commerce-app-demo/order-service/internal/models/orders"
	"github.com/commerce-app-demo/order-service/internal/service"
	orderspb "github.com/commerce-app-demo/order-service/proto"
)

type OrderServiceServer struct {
	orderspb.UnimplementedOrderServiceServer
	OrderService  *service.OrderService
	UserClient    *clients.UserClient
	ProductClient *clients.ProductClient
}

func (s *OrderServiceServer) GetOrders(ctx context.Context, req *orderspb.Empty) (*orderspb.OrderArray, error) {
	ordersList, err := s.OrderService.GetOrders()
	if err != nil {
		return nil, err
	}
	var orderArrayPb []*orderspb.Order
	for _, o := range ordersList {
		orderPb := &orderspb.Order{
			Id:         int32(o.Id),
			UserId:     int32(o.UserId),
			TotalPrice: o.TotalPrice,
			Status:     o.Status,
			CreatedAt:  o.CreatedAt,
			UpdatedAt:  o.UpdatedAt,
		}
		orderArrayPb = append(orderArrayPb, orderPb)
	}
	return &orderspb.OrderArray{Orders: orderArrayPb}, nil
}

func (s *OrderServiceServer) GetOrder(ctx context.Context, req *orderspb.GetOrderRequest) (*orderspb.Order, error) {
	order, err := s.OrderService.GetOrderById(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &orderspb.Order{
		Id:         int32(order.Id),
		UserId:     int32(order.UserId),
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}, nil
}

func (s *OrderServiceServer) CreateOrder(ctx context.Context, req *orderspb.CreateOrderRequest) (*orderspb.Order, error) {
	if req.UserId == 0 || req.Status == "" {
		return nil, fmt.Errorf("invalid request")
	}
	// Validate user exists
	_, err := s.UserClient.ValidateUser(fmt.Sprint(req.UserId))
	if err != nil {
		return nil, fmt.Errorf("user not found or not valid: %w", err)
	}
	order := &orders.OrderEntity{
		UserId:     int(req.UserId),
		TotalPrice: req.TotalPrice,
		Status:     req.Status,
	}
	created, err := s.OrderService.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	return &orderspb.Order{
		Id:         int32(created.Id),
		UserId:     int32(created.UserId),
		TotalPrice: created.TotalPrice,
		Status:     created.Status,
		CreatedAt:  created.CreatedAt,
		UpdatedAt:  created.UpdatedAt,
	}, nil
}

func (s *OrderServiceServer) UpdateOrder(ctx context.Context, req *orderspb.UpdateOrderRequest) (*orderspb.UpdateOrderResponse, error) {
	if req.Id == 0 {
		return nil, fmt.Errorf("invalid request id")
	}
	order := &orders.OrderEntity{
		UserId:     int(req.Order.UserId),
		TotalPrice: req.Order.TotalPrice,
		Status:     req.Order.Status,
	}
	updated, err := s.OrderService.UpdateOrder(int(req.Id), order)
	if err != nil {
		return nil, err
	}
	return &orderspb.UpdateOrderResponse{
		Success: true,
		UpdatedOrder: &orderspb.Order{
			Id:         int32(updated.Id),
			UserId:     int32(updated.UserId),
			TotalPrice: updated.TotalPrice,
			Status:     updated.Status,
			CreatedAt:  updated.CreatedAt,
			UpdatedAt:  updated.UpdatedAt,
		},
	}, nil
}

func (s *OrderServiceServer) DeleteOrder(ctx context.Context, req *orderspb.DeleteOrderRequest) (*orderspb.DeleteOrderResponse, error) {
	if req.Id == 0 {
		return nil, fmt.Errorf("invalid request id")
	}
	deleted, err := s.OrderService.DeleteOrder(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &orderspb.DeleteOrderResponse{
		Success: true,
		DeletedOrder: &orderspb.Order{
			Id:         int32(deleted.Id),
			UserId:     int32(deleted.UserId),
			TotalPrice: deleted.TotalPrice,
			Status:     deleted.Status,
			CreatedAt:  deleted.CreatedAt,
			UpdatedAt:  deleted.UpdatedAt,
		},
	}, nil
}

// Order Items
func (s *OrderServiceServer) GetOrderItems(ctx context.Context, req *orderspb.GetOrderItemsRequest) (*orderspb.OrderItemArray, error) {
	items, err := s.OrderService.GetOrderItems(int(req.OrderId))
	if err != nil {
		return nil, err
	}
	var itemArrayPb []*orderspb.OrderItem
	for _, i := range items {
		itemPb := &orderspb.OrderItem{
			Id:        int32(i.Id),
			OrderId:   int32(i.OrderId),
			ProductId: int32(i.ProductId),
			Quantity:  int32(i.Quantity),
			Price:     i.Price,
			CreatedAt: i.CreatedAt,
			UpdatedAt: i.UpdatedAt,
		}
		itemArrayPb = append(itemArrayPb, itemPb)
	}
	return &orderspb.OrderItemArray{Items: itemArrayPb}, nil
}

func (s *OrderServiceServer) CreateOrderItem(ctx context.Context, req *orderspb.CreateOrderItemRequest) (*orderspb.OrderItem, error) {
	// Validate product exists
	_, err := s.ProductClient.ValidateProduct(fmt.Sprint(req.ProductId))
	if err != nil {
		return nil, fmt.Errorf("product not found or not valid: %w", err)
	}
	item := &orders.OrderItemEntity{
		OrderId:   int(req.OrderId),
		ProductId: int(req.ProductId),
		Quantity:  int(req.Quantity),
		Price:     req.Price,
	}
	created, err := s.OrderService.CreateOrderItem(item)
	if err != nil {
		return nil, err
	}
	return &orderspb.OrderItem{
		Id:        int32(created.Id),
		OrderId:   int32(created.OrderId),
		ProductId: int32(created.ProductId),
		Quantity:  int32(created.Quantity),
		Price:     created.Price,
		CreatedAt: created.CreatedAt,
		UpdatedAt: created.UpdatedAt,
	}, nil
}

func (s *OrderServiceServer) UpdateOrderItem(ctx context.Context, req *orderspb.UpdateOrderItemRequest) (*orderspb.UpdateOrderItemResponse, error) {
	item := &orders.OrderItemEntity{
		OrderId:   int(req.Item.OrderId),
		ProductId: int(req.Item.ProductId),
		Quantity:  int(req.Item.Quantity),
		Price:     req.Item.Price,
	}
	updated, err := s.OrderService.UpdateOrderItem(int(req.Id), item)
	if err != nil {
		return nil, err
	}
	return &orderspb.UpdateOrderItemResponse{
		Success: true,
		UpdatedItem: &orderspb.OrderItem{
			Id:        int32(updated.Id),
			OrderId:   int32(updated.OrderId),
			ProductId: int32(updated.ProductId),
			Quantity:  int32(updated.Quantity),
			Price:     updated.Price,
			CreatedAt: updated.CreatedAt,
			UpdatedAt: updated.UpdatedAt,
		},
	}, nil
}

func (s *OrderServiceServer) DeleteOrderItem(ctx context.Context, req *orderspb.DeleteOrderItemRequest) (*orderspb.DeleteOrderItemResponse, error) {
	deleted, err := s.OrderService.DeleteOrderItem(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &orderspb.DeleteOrderItemResponse{
		Success: true,
		DeletedItem: &orderspb.OrderItem{
			Id:        int32(deleted.Id),
			OrderId:   int32(deleted.OrderId),
			ProductId: int32(deleted.ProductId),
			Quantity:  int32(deleted.Quantity),
			Price:     deleted.Price,
			CreatedAt: deleted.CreatedAt,
			UpdatedAt: deleted.UpdatedAt,
		},
	}, nil
}
