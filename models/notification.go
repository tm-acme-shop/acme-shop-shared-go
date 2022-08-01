package models

import "time"

// NotificationType represents the type of notification.
type NotificationType string

const (
	NotificationTypeEmail NotificationType = "email"
	NotificationTypeSMS   NotificationType = "sms"
	NotificationTypePush  NotificationType = "push"
)

// Notification represents a notification to be sent.
type Notification struct {
	ID        string           `json:"id"`
	Type      NotificationType `json:"type"`
	Recipient string           `json:"recipient"`
	Subject   string           `json:"subject"`
	Body      string           `json:"body"`
	Metadata  map[string]any   `json:"metadata,omitempty"`
	CreatedAt time.Time        `json:"createdAt"`
}
