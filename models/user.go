package models

import "time"

type UserRole string

const (
	RoleAdmin    UserRole = "admin"
	RoleCustomer UserRole = "customer"
	RoleVendor   UserRole = "vendor"
)

// UserV1 represents the legacy user model.
// Deprecated: Use User instead. This type will be removed in v3.0.
// Deprecated: Use User instead
type UserV1 struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password"` // TODO(TEAM-SEC): Should not include password in model
	CreatedAt time.Time `json:"created_at"`
}

// API-160: User represents a user in the system (v2 API).
// Added alongside UserV1 for gradual migration
type User struct {
	ID          string          `json:"id"`
	Email       string          `json:"email"`
	FirstName   string          `json:"first_name"`
	LastName    string          `json:"last_name"`
	Role        UserRole        `json:"role"`
	Active      bool            `json:"active"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	LastLoginAt time.Time       `json:"last_login_at,omitempty"`
	Preferences UserPreferences `json:"preferences"`
}

type UserPreferences struct {
	NotificationsEnabled bool   `json:"notifications_enabled"`
	Theme                string `json:"theme"`
	Locale               string `json:"locale"`
	Timezone             string `json:"timezone"`
}

// FullName returns the user's full name.
func (u *User) FullName() string {
	return u.FirstName + " " + u.LastName
}

// IsAdmin checks if the user has admin role.
func (u *User) IsAdmin() bool {
	return u.Role == RoleAdmin
}

// ToV1 converts a User to the legacy UserV1 format.
// Deprecated: Use User directly instead of converting to V1.
func (u *User) ToV1() *UserV1 {
	return &UserV1{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.FullName(),
		CreatedAt: u.CreatedAt,
	}
}

type CreateUserRequest struct {
	Email     string   `json:"email"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Password  string   `json:"password"`
	Role      UserRole `json:"role"`
}

type UpdateUserRequest struct {
	FirstName   *string          `json:"first_name,omitempty"`
	LastName    *string          `json:"last_name,omitempty"`
	Active      *bool            `json:"active,omitempty"`
	Preferences *UserPreferences `json:"preferences,omitempty"`
}

type UserListFilter struct {
	Role   *UserRole `json:"role,omitempty"`
	Active *bool     `json:"active,omitempty"`
	Search string    `json:"search,omitempty"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
}
