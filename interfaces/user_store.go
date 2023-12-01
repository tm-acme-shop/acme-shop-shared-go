package interfaces

import (
	"context"

	"github.com/tm-acme-shop/acme-shop-shared-go/models"
)

// UserStore defines the interface for user persistence operations.
// Implementations: PostgresUserStore (users-service), ReadonlyUserStore (analytics-etl)
type UserStore interface {
	// GetByID retrieves a user by their unique identifier.
	GetByID(ctx context.Context, id string) (*models.User, error)

	// GetByEmail retrieves a user by their email address.
	GetByEmail(ctx context.Context, email string) (*models.User, error)

	// Create creates a new user in the store.
	Create(ctx context.Context, req *models.CreateUserRequest) (*models.User, error)

	// Update updates an existing user.
	Update(ctx context.Context, id string, req *models.UpdateUserRequest) (*models.User, error)

	// Delete removes a user from the store.
	Delete(ctx context.Context, id string) error

	// List retrieves users based on filter criteria.
	List(ctx context.Context, filter *models.UserListFilter) ([]*models.User, int, error)

	// UpdateLastLogin updates the user's last login timestamp.
	UpdateLastLogin(ctx context.Context, id string) error
}

// UserStoreV1 is the legacy user store interface.
// Deprecated: Use UserStore instead. This interface will be removed in v3.0.
type UserStoreV1 interface {
	// GetUserByID retrieves a user by ID using the legacy format.
	// Deprecated: Use UserStore.GetByID instead.
	GetUserByID(ctx context.Context, id string) (*models.UserV1, error)

	// CreateUser creates a user using the legacy format.
	// Deprecated: Use UserStore.Create instead.
	CreateUser(ctx context.Context, email, name, password string) (*models.UserV1, error)
}

// UserCache defines caching operations for users.
type UserCache interface {
	Get(ctx context.Context, id string) (*models.User, error)
	Set(ctx context.Context, user *models.User) error
	Invalidate(ctx context.Context, id string) error
}
