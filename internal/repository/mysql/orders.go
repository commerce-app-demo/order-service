package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/commerce-app-demo/order-service/internal/models/orders"
)

type OrderRepository struct {
	DB *sql.DB
}

// Orders returns a list of orders (limit 50)
func (r *OrderRepository) Orders() ([]orders.OrderEntity, error) {
	query := "SELECT id, user_id, total_price, status, created_at, updated_at FROM orders LIMIT 50"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orderArray []orders.OrderEntity
	for rows.Next() {
		var o orders.OrderEntity
		err := rows.Scan(&o.Id, &o.UserId, &o.TotalPrice, &o.Status, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return nil, err
		}
		orderArray = append(orderArray, o)
	}
	if len(orderArray) < 1 {
		return nil, fmt.Errorf("Order list is empty")
	}
	return orderArray, nil
}

// OrderById returns a single order by id
func (r *OrderRepository) OrderById(id int) (*orders.OrderEntity, error) {
	query := "SELECT id, user_id, total_price, status, created_at, updated_at FROM orders WHERE id = ?"
	row := r.DB.QueryRow(query, id)
	var o orders.OrderEntity
	err := row.Scan(&o.Id, &o.UserId, &o.TotalPrice, &o.Status, &o.CreatedAt, &o.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &o, nil
}

// CreateOrder inserts a new order
func (r *OrderRepository) CreateOrder(order *orders.OrderEntity) (*orders.OrderEntity, error) {
	query := "INSERT INTO orders (user_id, total_price, status) VALUES (?, ?, ?)"
	res, err := r.DB.Exec(query, order.UserId, order.TotalPrice, order.Status)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.OrderById(int(id))
}

// UpdateOrder updates an order by id
func (r *OrderRepository) UpdateOrder(id int, columns map[string]any) (*orders.OrderEntity, error) {
	queryArgs := ""
	var args []any
	for colName, col := range columns {
		if queryArgs == "" {
			queryArgs = fmt.Sprintf("%s = ?", colName)
			args = append(args, col)
		} else {
			queryArgs = fmt.Sprintf("%s, %s = ?", queryArgs, colName)
			args = append(args, col)
		}
	}
	args = append(args, id)
	query := fmt.Sprintf("UPDATE orders SET %s WHERE id = ?", queryArgs)
	_, err := r.DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return r.OrderById(id)
}

// DeleteOrder deletes an order by id
func (r *OrderRepository) DeleteOrder(id int) (*orders.OrderEntity, error) {
	order, err := r.OrderById(id)
	if err != nil {
		return nil, err
	}
	_, err = r.DB.Exec("DELETE FROM orders WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// CRUD for OrderItems

// OrderItemsByOrderId returns all items for a given order
func (r *OrderRepository) OrderItemsByOrderId(orderId int) ([]orders.OrderItemEntity, error) {
	query := "SELECT id, order_id, product_id, quantity, price, created_at, updated_at FROM order_items WHERE order_id = ?"
	rows, err := r.DB.Query(query, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []orders.OrderItemEntity
	for rows.Next() {
		var item orders.OrderItemEntity
		err := rows.Scan(&item.Id, &item.OrderId, &item.ProductId, &item.Quantity, &item.Price, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// CreateOrderItem inserts a new order item
func (r *OrderRepository) CreateOrderItem(item *orders.OrderItemEntity) (*orders.OrderItemEntity, error) {
	query := "INSERT INTO order_items (order_id, product_id, quantity, price) VALUES (?, ?, ?, ?)"
	res, err := r.DB.Exec(query, item.OrderId, item.ProductId, item.Quantity, item.Price)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return r.OrderItemById(int(id))
}

// OrderItemById returns a single order item by id
func (r *OrderRepository) OrderItemById(id int) (*orders.OrderItemEntity, error) {
	query := "SELECT id, order_id, product_id, quantity, price, created_at, updated_at FROM order_items WHERE id = ?"
	row := r.DB.QueryRow(query, id)
	var item orders.OrderItemEntity
	err := row.Scan(&item.Id, &item.OrderId, &item.ProductId, &item.Quantity, &item.Price, &item.CreatedAt, &item.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// UpdateOrderItem updates an order item by id
func (r *OrderRepository) UpdateOrderItem(id int, columns map[string]any) (*orders.OrderItemEntity, error) {
	queryArgs := ""
	var args []any
	for colName, col := range columns {
		if queryArgs == "" {
			queryArgs = fmt.Sprintf("%s = ?", colName)
			args = append(args, col)
		} else {
			queryArgs = fmt.Sprintf("%s, %s = ?", queryArgs, colName)
			args = append(args, col)
		}
	}
	args = append(args, id)
	query := fmt.Sprintf("UPDATE order_items SET %s WHERE id = ?", queryArgs)
	_, err := r.DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return r.OrderItemById(id)
}

// DeleteOrderItem deletes an order item by id
func (r *OrderRepository) DeleteOrderItem(id int) (*orders.OrderItemEntity, error) {
	item, err := r.OrderItemById(id)
	if err != nil {
		return nil, err
	}
	_, err = r.DB.Exec("DELETE FROM order_items WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
