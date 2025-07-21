package orders

// OrderEntity represents an order in the system.
type OrderEntity struct {
	Id         int
	UserId     int
	TotalPrice float64
	Status     string
	CreatedAt  string
	UpdatedAt  string
}

// OrderItemEntity represents an item in an order.
type OrderItemEntity struct {
	Id        int
	OrderId   int
	ProductId int
	Quantity  int
	Price     float64
	CreatedAt string
	UpdatedAt string
}
