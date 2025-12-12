package interfaces

import (
	"context"

	"github.com/tm-acme-shop/acme-shop-shared-go/models"
)

// OrderRepository defines the interface for order persistence operations.
// Implementations: PostgresOrderRepository (orders-service)
type OrderRepository interface {
	// GetByID retrieves an order by its unique identifier.
	GetByID(ctx context.Context, id string) (*models.Order, error)

	// Create creates a new order.
	Create(ctx context.Context, req *models.CreateOrderRequest) (*models.Order, error)

	// UpdateStatus updates the status of an order.
	UpdateStatus(ctx context.Context, id string, req *models.UpdateOrderStatusRequest) (*models.Order, error)

	// List retrieves orders based on filter criteria.
	List(ctx context.Context, filter *models.OrderListFilter) ([]*models.Order, int, error)

	// GetByUserID retrieves all orders for a specific user.
	GetByUserID(ctx context.Context, userID string, limit, offset int) ([]*models.Order, int, error)

	// Delete soft-deletes an order (marks as cancelled).
	Delete(ctx context.Context, id string) error

	// SetPaymentID associates a payment with an order.
	SetPaymentID(ctx context.Context, orderID, paymentID string) error
}

// OrderEventPublisher publishes order events to a message queue.
type OrderEventPublisher interface {
	// PublishOrderCreated publishes an order created event.
	PublishOrderCreated(ctx context.Context, order *models.Order) error

	// PublishOrderStatusChanged publishes an order status change event.
	PublishOrderStatusChanged(ctx context.Context, order *models.Order, previousStatus models.OrderStatus) error

	// PublishOrderCancelled publishes an order cancellation event.
	PublishOrderCancelled(ctx context.Context, order *models.Order, reason string) error
}

// OrderAnalytics provides analytics operations for orders.
type OrderAnalytics interface {
	// GetDailyStats retrieves daily order statistics.
	GetDailyStats(ctx context.Context, startDate, endDate string) ([]OrderDailyStat, error)

	// GetTopProducts retrieves the most ordered products.
	GetTopProducts(ctx context.Context, limit int) ([]ProductStat, error)

	// GetRevenueByPeriod retrieves revenue grouped by time period.
	GetRevenueByPeriod(ctx context.Context, period string, startDate, endDate string) ([]RevenueStat, error)
}

type OrderDailyStat struct {
	Date       string `json:"date"`
	OrderCount int    `json:"order_count"`
	Revenue    int64  `json:"revenue"`
}

type ProductStat struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	OrderCount  int    `json:"order_count"`
	Quantity    int    `json:"quantity"`
	Revenue     int64  `json:"revenue"`
}

type RevenueStat struct {
	Period  string `json:"period"`
	Revenue int64  `json:"revenue"`
	Orders  int    `json:"orders"`
}
