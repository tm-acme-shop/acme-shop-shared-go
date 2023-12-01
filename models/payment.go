package models

import "time"

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard PaymentMethod = "credit_card"
	PaymentMethodDebitCard  PaymentMethod = "debit_card"
	PaymentMethodPayPal     PaymentMethod = "paypal"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
	PaymentMethodCrypto     PaymentMethod = "crypto"
)

type Payment struct {
	ID              string        `json:"id"`
	OrderID         string        `json:"order_id"`
	UserID          string        `json:"user_id"`
	Amount          Money         `json:"amount"`
	Method          PaymentMethod `json:"method"`
	Status          PaymentStatus `json:"status"`
	ProviderID      string        `json:"provider_id,omitempty"`
	ProviderRef     string        `json:"provider_ref,omitempty"`
	ErrorMessage    string        `json:"error_message,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	CompletedAt     *time.Time    `json:"completed_at,omitempty"`
}

type ProcessPaymentRequest struct {
	OrderID     string            `json:"order_id"`
	UserID      string            `json:"user_id"`
	Amount      Money             `json:"amount"`
	Method      PaymentMethod     `json:"method"`
	CardToken   string            `json:"card_token,omitempty"`
	ReturnURL   string            `json:"return_url,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

type ProcessPaymentResponse struct {
	PaymentID   string        `json:"payment_id"`
	Status      PaymentStatus `json:"status"`
	RedirectURL string        `json:"redirect_url,omitempty"`
	ProviderRef string        `json:"provider_ref,omitempty"`
}

type RefundRequest struct {
	PaymentID string `json:"payment_id"`
	Amount    Money  `json:"amount"`
	Reason    string `json:"reason"`
}

type RefundResponse struct {
	RefundID    string        `json:"refund_id"`
	PaymentID   string        `json:"payment_id"`
	Amount      Money         `json:"amount"`
	Status      PaymentStatus `json:"status"`
	ProviderRef string        `json:"provider_ref,omitempty"`
}

// LegacyPaymentRequest is the old payment request format.
// Deprecated: Use ProcessPaymentRequest instead.
type LegacyPaymentRequest struct {
	OrderID   string  `json:"order_id"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	CardNumber string `json:"card_number"` // TODO(TEAM-SEC): Never send raw card numbers
	CVV       string  `json:"cvv"`
}

func (p *Payment) IsSuccessful() bool {
	return p.Status == PaymentStatusCompleted
}

func (p *Payment) CanRefund() bool {
	return p.Status == PaymentStatusCompleted
}
