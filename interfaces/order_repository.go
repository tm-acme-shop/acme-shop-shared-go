package interfaces

import (
	"context"

	"github.com/tm-acme-shop/acme-shop-shared-go/models"
)

// OrderRepository defines the interface for order persistence operations.
// Implementations: PostgresOrderRepository (orders-service)
type OrderRepository interface {
	// GetByID retrieves an order by ID.
	GetByID(ctx context.Context, id string) (*models.Order, error)

	// Create creates a new order.
	Create(ctx context.Context, order *models.Order) (*models.Order, error)

	// Update updates an existing order.
	Update(ctx context.Context, order *models.Order) (*models.Order, error)

	// ListByUser retrieves all orders for a user.
	ListByUser(ctx context.Context, userID string, limit, offset int) ([]*models.Order, int, error)

	// UpdateStatus updates the status of an order.
	UpdateStatus(ctx context.Context, id string, status models.OrderStatus) error
}
