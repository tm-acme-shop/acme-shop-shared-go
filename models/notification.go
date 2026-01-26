package models

import "time"

type NotificationType string

const (
	NotificationTypeEmail              NotificationType = "email"
	NotificationTypeSMS                NotificationType = "sms"
	NotificationTypePush               NotificationType = "push"
	NotificationTypeSlack              NotificationType = "slack"
	NotificationTypeOrderConfirmation  NotificationType = "order_confirmation"
	NotificationTypeOrderShipped       NotificationType = "order_shipped"
	NotificationTypeOrderDelivered     NotificationType = "order_delivered"
	NotificationTypeOrderCancelled     NotificationType = "order_cancelled"
)

type NotificationChannel string

const (
	NotificationChannelEmail NotificationChannel = "email"
	NotificationChannelSMS   NotificationChannel = "sms"
	NotificationChannelPush  NotificationChannel = "push"
)

type NotificationStatus string

const (
	NotificationStatusPending   NotificationStatus = "pending"
	NotificationStatusSent      NotificationStatus = "sent"
	NotificationStatusDelivered NotificationStatus = "delivered"
	NotificationStatusFailed    NotificationStatus = "failed"
	NotificationStatusBounced   NotificationStatus = "bounced"
)

type NotificationPriority string

const (
	NotificationPriorityLow    NotificationPriority = "low"
	NotificationPriorityNormal NotificationPriority = "normal"
	NotificationPriorityHigh   NotificationPriority = "high"
	NotificationPriorityCritical NotificationPriority = "critical"
)

type Notification struct {
	ID           string                 `json:"id"`
	UserID       string                 `json:"user_id"`
	Type         NotificationType       `json:"type"`
	Channel      NotificationChannel    `json:"channel"`
	Status       NotificationStatus     `json:"status"`
	Priority     NotificationPriority   `json:"priority"`
	Recipient    string                 `json:"recipient"`
	Title        string                 `json:"title,omitempty"`
	Subject      string                 `json:"subject,omitempty"`
	Body         string                 `json:"body"`
	TemplateID   string                 `json:"template_id,omitempty"`
	TemplateData map[string]interface{} `json:"template_data,omitempty"`
	Data         map[string]interface{} `json:"data,omitempty"`
	Metadata     map[string]string      `json:"metadata,omitempty"`
	RetryCount   int                    `json:"retry_count"`
	MaxRetries   int                    `json:"max_retries"`
	ErrorMessage string                 `json:"error_message,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
	SentAt       *time.Time             `json:"sent_at,omitempty"`
	DeliveredAt  *time.Time             `json:"delivered_at,omitempty"`
}

type SendNotificationRequest struct {
	Type        NotificationType       `json:"type"`
	Priority    NotificationPriority   `json:"priority"`
	Recipient   string                 `json:"recipient"`
	Subject     string                 `json:"subject,omitempty"`
	Body        string                 `json:"body,omitempty"`
	TemplateID  string                 `json:"template_id,omitempty"`
	TemplateData map[string]interface{} `json:"template_data,omitempty"`
	Metadata    map[string]string      `json:"metadata,omitempty"`
}

type SendBatchRequest struct {
	Notifications []SendNotificationRequest `json:"notifications"`
}

type NotificationResult struct {
	NotificationID string             `json:"notification_id"`
	Status         NotificationStatus `json:"status"`
	ErrorMessage   string             `json:"error_message,omitempty"`
}

type EmailTemplate struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Subject     string `json:"subject"`
	HTMLBody    string `json:"html_body"`
	TextBody    string `json:"text_body"`
	Variables   []string `json:"variables"`
}

type SendEmailRequest struct {
	To          string                 `json:"to"`
	Subject     string                 `json:"subject,omitempty"`
	Body        string                 `json:"body,omitempty"`
	Template    string                 `json:"template,omitempty"`
	TemplateData map[string]interface{} `json:"template_data,omitempty"`
}

type SendSMSRequest struct {
	To          string                 `json:"to"`
	Message     string                 `json:"message,omitempty"`
	Template    string                 `json:"template,omitempty"`
	TemplateData map[string]interface{} `json:"template_data,omitempty"`
}

type SendPushRequest struct {
	UserID      string                 `json:"user_id"`
	Title       string                 `json:"title"`
	Body        string                 `json:"body"`
	Data        map[string]interface{} `json:"data,omitempty"`
}

func (n *Notification) CanRetry() bool {
	return n.Status == NotificationStatusFailed && n.RetryCount < n.MaxRetries
}

func (n *Notification) IncrementRetry() {
	n.RetryCount++
}
