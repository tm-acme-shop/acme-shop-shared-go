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
}

// LegacyPaymentClient is the old payment client interface.
// Deprecated: Use PaymentClient instead. Migration deadline: Q2 2024.
// TODO(TEAM-PAYMENTS): Remove after all services migrated to PaymentClient
type LegacyPaymentClient interface {
	// ProcessLegacyPayment processes a payment using the old format.
	// Deprecated: Use PaymentClient.ProcessPayment instead.
	ProcessLegacyPayment(ctx context.Context, req *models.LegacyPaymentRequest) (string, error)

	// GetLegacyPaymentStatus gets payment status.
	// Deprecated: Use PaymentClient.GetPaymentStatus instead.
	GetLegacyPaymentStatus(ctx context.Context, paymentID string) (string, error)
}
