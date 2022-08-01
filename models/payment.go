package models

import "time"

// PaymentStatus represents the status of a payment.
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

// Payment represents a payment in the system.
type Payment struct {
	ID            string        `json:"id"`
	OrderID       string        `json:"orderId"`
	Amount        int64         `json:"amount"`
	Currency      string        `json:"currency"`
	Status        PaymentStatus `json:"status"`
	PaymentMethod string        `json:"paymentMethod"`
	ProviderID    string        `json:"providerId"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
}

// ProcessPaymentRequest is the request to process a payment.
type ProcessPaymentRequest struct {
	OrderID       string `json:"orderId"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	PaymentMethod string `json:"paymentMethod"`
	CustomerID    string `json:"customerId"`
}

// ProcessPaymentResponse is the response from processing a payment.
type ProcessPaymentResponse struct {
	PaymentID  string        `json:"paymentId"`
	Status     PaymentStatus `json:"status"`
	ProviderID string        `json:"providerId"`
}

// RefundRequest is the request to refund a payment.
type RefundRequest struct {
	PaymentID string `json:"paymentId"`
	Amount    int64  `json:"amount"`
	Reason    string `json:"reason"`
}

// RefundResponse is the response from a refund.
type RefundResponse struct {
	RefundID string        `json:"refundId"`
	Status   PaymentStatus `json:"status"`
}

// LegacyPaymentRequest is the legacy payment request format.
type LegacyPaymentRequest struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Method  string  `json:"method"`
}
