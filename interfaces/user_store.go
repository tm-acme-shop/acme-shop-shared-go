package interfaces

import (
	"context"

	"github.com/tm-acme-shop/acme-shop-shared-go/models"
)

// UserStoreV1 defines the interface for user persistence operations.
type UserStoreV1 interface {
	// GetUserByID retrieves a user by their unique identifier.
	GetUserByID(ctx context.Context, id string) (*models.UserV1, error)

	// GetUserByEmail retrieves a user by their email address.
	GetUserByEmail(ctx context.Context, email string) (*models.UserV1, error)

	// CreateUser creates a new user in the store.
	CreateUser(ctx context.Context, email, name, password string) (*models.UserV1, error)

	// UpdateUser updates an existing user.
	UpdateUser(ctx context.Context, id, email, name string) (*models.UserV1, error)

	// DeleteUser removes a user from the store.
	DeleteUser(ctx context.Context, id string) error

	// ListUsers retrieves all users with pagination.
	ListUsers(ctx context.Context, limit, offset int) ([]*models.UserV1, error)
}
