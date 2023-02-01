package models

import "time"

// UserV1 represents a user in the legacy format.
// Deprecated: Use User instead. This type will be removed in v3.0.
// TODO(TEAM-BACKEND): Remove after all services migrated to v2.
type UserV1 struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// UserRole represents the role of a user.
type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleCustomer UserRole = "customer"
	UserRoleVendor   UserRole = "vendor"
)

// User represents a user in the system (v2 API).
type User struct {
	ID          string          `json:"id"`
	Email       string          `json:"email"`
	FirstName   string          `json:"firstName"`
	LastName    string          `json:"lastName"`
	Role        UserRole        `json:"role"`
	Active      bool            `json:"active"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	LastLoginAt *time.Time      `json:"lastLoginAt,omitempty"`
	Preferences UserPreferences `json:"preferences"`
}

// UserPreferences represents user preferences.
type UserPreferences struct {
	NotificationsEnabled bool   `json:"notificationsEnabled"`
	Theme                string `json:"theme"`
	Locale               string `json:"locale"`
	Timezone             string `json:"timezone"`
}

// CreateUserRequest is the request to create a new user.
type CreateUserRequest struct {
	Email     string   `json:"email"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Password  string   `json:"password"`
	Role      UserRole `json:"role"`
}

// UpdateUserRequest is the request to update a user.
type UpdateUserRequest struct {
	FirstName   *string          `json:"firstName,omitempty"`
	LastName    *string          `json:"lastName,omitempty"`
	Active      *bool            `json:"active,omitempty"`
	Preferences *UserPreferences `json:"preferences,omitempty"`
}

// UserListFilter defines filter criteria for listing users.
type UserListFilter struct {
	Role   *UserRole `json:"role,omitempty"`
	Active *bool     `json:"active,omitempty"`
	Search string    `json:"search,omitempty"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
}

// ToUserV1 converts a User to the legacy UserV1 format.
// Deprecated: Use User directly instead of converting to V1.
// TODO(TEAM-BACKEND): Remove after v1 API is disabled
func ToUserV1(user *User) *UserV1 {
	return &UserV1{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.FirstName + " " + user.LastName,
		CreatedAt: user.CreatedAt,
	}
}

// FromUserV1 converts a UserV1 to the new User format.
// Deprecated: Will be removed when v1 API is disabled.
// TODO(TEAM-BACKEND): Remove after v1 API is disabled
func FromUserV1(userV1 *UserV1) *User {
	return &User{
		ID:        userV1.ID,
		Email:     userV1.Email,
		FirstName: userV1.Name,
		LastName:  "",
		Role:      UserRoleCustomer,
		Active:    true,
		CreatedAt: userV1.CreatedAt,
		UpdatedAt: userV1.CreatedAt,
		Preferences: UserPreferences{
			NotificationsEnabled: true,
			Theme:                "system",
			Locale:               "en-US",
			Timezone:             "UTC",
		},
	}
}
