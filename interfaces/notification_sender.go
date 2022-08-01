package interfaces

import (
	"context"

	"github.com/tm-acme-shop/acme-shop-shared-go/models"
)

// NotificationSender defines the interface for sending notifications.
// Implementations: EmailNotificationSender, SMSNotificationSender
type NotificationSender interface {
	// Send sends a notification to the specified recipients.
	Send(ctx context.Context, notification *models.Notification) error

	// SendBatch sends multiple notifications.
	SendBatch(ctx context.Context, notifications []*models.Notification) error
}
