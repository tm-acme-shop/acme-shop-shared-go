package models

import "time"

// OrderStatus represents the status of an order.
type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusConfirmed  OrderStatus = "confirmed"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

// Order represents an order in the system.
type Order struct {
	ID         string      `json:"id"`
	UserID     string      `json:"userId"`
	Status     OrderStatus `json:"status"`
	Items      []OrderItem `json:"items"`
	TotalPrice int64       `json:"totalPrice"`
	Currency   string      `json:"currency"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}

// OrderItem represents an item in an order.
type OrderItem struct {
	ProductID string `json:"productId"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	UnitPrice int64  `json:"unitPrice"`
}
