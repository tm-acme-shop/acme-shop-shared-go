package models

import "time"

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusConfirmed  OrderStatus = "confirmed"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusRefunded   OrderStatus = "refunded"
)

type Order struct {
	ID              string      `json:"id"`
	UserID          string      `json:"user_id"`
	Status          OrderStatus `json:"status"`
	Items           []OrderItem `json:"items"`
	ShippingAddress Address     `json:"shipping_address"`
	BillingAddress  Address     `json:"billing_address"`
	Subtotal        Money       `json:"subtotal"`
	Tax             Money       `json:"tax"`
	ShippingCost    Money       `json:"shipping_cost"`
	Total           Money       `json:"total"`
	PaymentID       string      `json:"payment_id,omitempty"`
	Notes           string      `json:"notes,omitempty"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	ShippedAt       *time.Time  `json:"shipped_at,omitempty"`
	DeliveredAt     *time.Time  `json:"delivered_at,omitempty"`
}

type OrderItem struct {
	ID          string `json:"id"`
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	UnitPrice   Money  `json:"unit_price"`
	Total       Money  `json:"total"`
}

type Address struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2,omitempty"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type Money struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func (m Money) ToFloat() float64 {
	return float64(m.Amount) / 100
}

func NewMoney(amount float64, currency string) Money {
	return Money{
		Amount:   int64(amount * 100),
		Currency: currency,
	}
}

type CreateOrderRequest struct {
	UserID          string      `json:"user_id"`
	Items           []OrderItem `json:"items"`
	ShippingAddress Address     `json:"shipping_address"`
	BillingAddress  Address     `json:"billing_address"`
	Subtotal        Money       `json:"subtotal"`
	Tax             Money       `json:"tax"`
	Total           Money       `json:"total"`
	Notes           string      `json:"notes,omitempty"`
}

type UpdateOrderStatusRequest struct {
	Status OrderStatus `json:"status"`
	Notes  string      `json:"notes,omitempty"`
}

type OrderListFilter struct {
	UserID    string       `json:"user_id,omitempty"`
	Status    *OrderStatus `json:"status,omitempty"`
	StartDate *time.Time   `json:"start_date,omitempty"`
	EndDate   *time.Time   `json:"end_date,omitempty"`
	Limit     int          `json:"limit"`
	Offset    int          `json:"offset"`
}

func (o *Order) CanCancel() bool {
	return o.Status == OrderStatusPending || o.Status == OrderStatusConfirmed
}

func (o *Order) CanRefund() bool {
	return o.Status == OrderStatusDelivered && o.PaymentID != ""
}

func (o *Order) CalculateTotal() {
	var subtotal int64
	for _, item := range o.Items {
		subtotal += item.Total.Amount
	}
	o.Subtotal = Money{Amount: subtotal, Currency: o.Items[0].UnitPrice.Currency}
	o.Total = Money{
		Amount:   o.Subtotal.Amount + o.Tax.Amount + o.ShippingCost.Amount,
		Currency: o.Subtotal.Currency,
	}
}
