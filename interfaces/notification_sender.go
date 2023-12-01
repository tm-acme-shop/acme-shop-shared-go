package interfaces

import (
	"context"

	"github.com/tm-acme-shop/acme-shop-shared-go/models"
)

// NotificationSender defines the interface for sending notifications.
// Implementations: EmailSender, SMSSender, PushSender (notifications-service)
type NotificationSender interface {
	// Send sends a single notification.
	Send(ctx context.Context, req *models.SendNotificationRequest) (*models.NotificationResult, error)

	// SendBatch sends multiple notifications in a batch.
	SendBatch(ctx context.Context, req *models.SendBatchRequest) ([]*models.NotificationResult, error)

	// GetStatus retrieves the status of a sent notification.
	GetStatus(ctx context.Context, notificationID string) (*models.Notification, error)

	// Cancel cancels a pending notification.
	Cancel(ctx context.Context, notificationID string) error
}

// EmailSender is a specialized interface for email notifications.
type EmailSender interface {
	NotificationSender

	// SendWithTemplate sends an email using a predefined template.
	SendWithTemplate(ctx context.Context, recipient, templateID string, data map[string]interface{}) (*models.NotificationResult, error)

	// ValidateEmail validates an email address format and deliverability.
	ValidateEmail(ctx context.Context, email string) (bool, error)
}

// LegacyEmailSender is the old email sending interface.
// Deprecated: Use EmailSender instead.
type LegacyEmailSender interface {
	// SendEmailLegacy sends an email using the old format.
	// Deprecated: Use EmailSender.Send instead.
	// TODO(TEAM-NOTIFICATIONS): Migrate all callers to new interface
	SendEmailLegacy(to, subject, body string) error

	// SendEmailWithAttachment sends an email with attachments.
	// Deprecated: Use EmailSender.Send with attachment support.
	SendEmailWithAttachment(to, subject, body string, attachments [][]byte) error
}

// SMSSender is a specialized interface for SMS notifications.
type SMSSender interface {
	NotificationSender

	// ValidatePhoneNumber validates a phone number format.
	ValidatePhoneNumber(ctx context.Context, phone string) (bool, error)
}

// PushSender is a specialized interface for push notifications.
type PushSender interface {
	NotificationSender

	// RegisterDevice registers a device for push notifications.
	RegisterDevice(ctx context.Context, userID, deviceToken string, platform string) error

	// UnregisterDevice removes a device from push notifications.
	UnregisterDevice(ctx context.Context, deviceToken string) error
}
