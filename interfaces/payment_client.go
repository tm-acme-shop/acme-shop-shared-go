package interfaces

import (
	"context"

	"github.com/tm-acme-shop/acme-shop-shared-go/models"
)

// PaymentClient defines the interface for payment processing operations.
// Implementations: StripePaymentClient (payments-service), MockPaymentClient (tests)
type PaymentClient interface {
	// ProcessPayment initiates a payment for an order.
	ProcessPayment(ctx context.Context, req *models.ProcessPaymentRequest) (*models.ProcessPaymentResponse, error)

	// GetPaymentStatus retrieves the current status of a payment.
	GetPaymentStatus(ctx context.Context, paymentID string) (*models.Payment, error)

	// Refund processes a refund for a completed payment.
	Refund(ctx context.Context, req *models.RefundRequest) (*models.RefundResponse, error)

	// CancelPayment cancels a pending payment.
	CancelPayment(ctx context.Context, paymentID string) error

	// ValidateWebhook validates an incoming webhook from the payment provider.
	ValidateWebhook(ctx context.Context, payload []byte, signature string) (bool, error)
}

// LegacyPaymentClient is the old payment client interface.
// Deprecated: Use PaymentClient instead. Migration deadline: Q2 2024.
type LegacyPaymentClient interface {
	// ProcessLegacyPayment processes a payment using the old format.
	// Deprecated: Use PaymentClient.ProcessPayment instead.
	// TODO(TEAM-PAYMENTS): Remove after migration complete
	ProcessLegacyPayment(ctx context.Context, req *models.LegacyPaymentRequest) (string, error)

	// GetStatus retrieves payment status by order ID (legacy behavior).
	// Deprecated: Use PaymentClient.GetPaymentStatus with payment ID.
	GetStatus(ctx context.Context, orderID string) (string, error)
}

// PaymentProvider represents different payment gateway providers.
type PaymentProvider string

const (
	PaymentProviderStripe   PaymentProvider = "stripe"
	PaymentProviderPayPal   PaymentProvider = "paypal"
	PaymentProviderSquare   PaymentProvider = "square"
	PaymentProviderLegacy   PaymentProvider = "legacy" // Deprecated: Old in-house system
)

// PaymentClientFactory creates PaymentClient instances for different providers.
type PaymentClientFactory interface {
	Create(provider PaymentProvider) (PaymentClient, error)
}
