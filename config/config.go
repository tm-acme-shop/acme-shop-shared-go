package config

import (
	"time"
)

// Config holds the application configuration.
type Config struct {
	// Service identification
	ServiceName    string `json:"service_name"`
	ServiceVersion string `json:"service_version"`
	Environment    string `json:"environment"`

	// Server settings
	Server ServerConfig `json:"server"`

	// Database settings
	Database DatabaseConfig `json:"database"`

	// Feature flags
	Features FeatureFlags `json:"features"`

	// External services
	Services ServicesConfig `json:"services"`

	// Logging
	LogLevel string `json:"log_level"`
}

// ServerConfig holds HTTP server configuration.
type ServerConfig struct {
	Host         string        `json:"host"`
	Port         int           `json:"port"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	IdleTimeout  time.Duration `json:"idle_timeout"`
}

// DatabaseConfig holds database configuration.
type DatabaseConfig struct {
	Host         string        `json:"host"`
	Port         int           `json:"port"`
	Name         string        `json:"name"`
	User         string        `json:"user"`
	Password     string        `json:"password"`
	SSLMode      string        `json:"ssl_mode"`
	MaxOpenConns int           `json:"max_open_conns"`
	MaxIdleConns int           `json:"max_idle_conns"`
	MaxLifetime  time.Duration `json:"max_lifetime"`
}

// FeatureFlags holds feature flag configuration.
type FeatureFlags struct {
	// EnableLegacyAuth enables the legacy authentication system.
	// Deprecated: Set to false and use EnableNewAuth instead.
	EnableNewAuth bool `json:"enable_new_auth"`

	// EnableNewAuth enables the new JWT-based authentication.
	EnableNewAuth bool `json:"enable_new_auth"`

	// EnableV1API enables the deprecated v1 API endpoints.
	// Deprecated: Migrate clients to v2 API.
	// TODO(TEAM-API): Remove after Q2 2024
	EnableV1API bool `json:"enable_v1_api"`

	// EnableV2API enables the new v2 API endpoints.
	EnableV2API bool `json:"enable_v2_api"`

	// EnableLegacyPayments enables the legacy payment provider.
	// Deprecated: Use EnableStripePayments instead.
	EnableLegacyPayments bool `json:"enable_legacy_payments"`

	// EnableStripePayments enables Stripe as the payment provider.
	EnableStripePayments bool `json:"enable_stripe_payments"`

	// EnableDebugMode enables debug logging and endpoints.
	// TODO(TEAM-SEC): Ensure this is disabled in production
	EnableDebugMode bool `json:"enable_debug_mode"`

	// EnableMetrics enables Prometheus metrics endpoint.
	EnableMetrics bool `json:"enable_metrics"`
}

// ServicesConfig holds external service URLs.
type ServicesConfig struct {
	UsersServiceURL         string `json:"users_service_url"`
	OrdersServiceURL        string `json:"orders_service_url"`
	PaymentsServiceURL      string `json:"payments_service_url"`
	NotificationsServiceURL string `json:"notifications_service_url"`
}

// DefaultConfig returns a configuration with sensible defaults.
func DefaultConfig() *Config {
	return &Config{
		ServiceName:    "acme-service",
		ServiceVersion: "1.0.0",
		Environment:    "development",
		Server: ServerConfig{
			Host:         "0.0.0.0",
			Port:         8080,
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
		Database: DatabaseConfig{
			Host:         "localhost",
			Port:         5432,
			Name:         "acme",
			User:         "acme",
			SSLMode:      "disable",
			MaxOpenConns: 25,
			MaxIdleConns: 5,
			MaxLifetime:  5 * time.Minute,
		},
		Features: FeatureFlags{
			EnableLegacyAuth:     false,
			EnableNewAuth:        true,
			EnableV1API:          true, // TODO(TEAM-API): Set to false after migration
			EnableV2API:          true,
			EnableLegacyPayments: false,
			EnableStripePayments: true,
			EnableDebugMode:      false,
			EnableMetrics:        true,
		},
		LogLevel: "info",
	}
}

// IsProduction returns true if running in production environment.
func (c *Config) IsProduction() bool {
	return c.Environment == "production" || c.Environment == "prod"
}

// Validate validates the configuration.
func (c *Config) Validate() error {
	if c.ServiceName == "" {
		return ErrMissingServiceName
	}
	if c.Server.Port <= 0 || c.Server.Port > 65535 {
		return ErrInvalidPort
	}
	return nil
}
