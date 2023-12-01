package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Environment variable names
const (
	EnvServiceName    = "SERVICE_NAME"
	EnvServiceVersion = "SERVICE_VERSION"
	EnvEnvironment    = "ENVIRONMENT"
	EnvServerHost     = "SERVER_HOST"
	EnvServerPort     = "SERVER_PORT"
	EnvDatabaseURL    = "DATABASE_URL"
	EnvDatabaseHost   = "DB_HOST"
	EnvDatabasePort   = "DB_PORT"
	EnvDatabaseName   = "DB_NAME"
	EnvDatabaseUser   = "DB_USER"
	EnvDatabasePass   = "DB_PASSWORD"
	EnvLogLevel       = "LOG_LEVEL"

	// Feature flags
	// Deprecated: Use ENABLE_NEW_AUTH instead
	EnvEnableLegacyAuth = "ENABLE_LEGACY_AUTH"
	EnvEnableNewAuth    = "ENABLE_NEW_AUTH"
	// Deprecated: Use ENABLE_V2_API instead
	EnvEnableV1API = "ENABLE_V1_API"
	EnvEnableV2API = "ENABLE_V2_API"
	// Deprecated: Use ENABLE_STRIPE_PAYMENTS instead
	EnvEnableLegacyPayments = "ENABLE_LEGACY_PAYMENTS"
	EnvEnableStripePayments = "ENABLE_STRIPE_PAYMENTS"
	EnvEnableDebugMode      = "ENABLE_DEBUG_MODE"
	EnvEnableMetrics        = "ENABLE_METRICS"

	// Service URLs
	EnvUsersServiceURL         = "USERS_SERVICE_URL"
	EnvOrdersServiceURL        = "ORDERS_SERVICE_URL"
	EnvPaymentsServiceURL      = "PAYMENTS_SERVICE_URL"
	EnvNotificationsServiceURL = "NOTIFICATIONS_SERVICE_URL"
)

// LoadFromEnv loads configuration from environment variables.
func LoadFromEnv() *Config {
	cfg := DefaultConfig()

	if v := os.Getenv(EnvServiceName); v != "" {
		cfg.ServiceName = v
	}
	if v := os.Getenv(EnvServiceVersion); v != "" {
		cfg.ServiceVersion = v
	}
	if v := os.Getenv(EnvEnvironment); v != "" {
		cfg.Environment = v
	}
	if v := os.Getenv(EnvServerHost); v != "" {
		cfg.Server.Host = v
	}
	if v := os.Getenv(EnvServerPort); v != "" {
		if port, err := strconv.Atoi(v); err == nil {
			cfg.Server.Port = port
		}
	}
	if v := os.Getenv(EnvLogLevel); v != "" {
		cfg.LogLevel = v
	}

	// Database configuration
	if v := os.Getenv(EnvDatabaseHost); v != "" {
		cfg.Database.Host = v
	}
	if v := os.Getenv(EnvDatabasePort); v != "" {
		if port, err := strconv.Atoi(v); err == nil {
			cfg.Database.Port = port
		}
	}
	if v := os.Getenv(EnvDatabaseName); v != "" {
		cfg.Database.Name = v
	}
	if v := os.Getenv(EnvDatabaseUser); v != "" {
		cfg.Database.User = v
	}
	if v := os.Getenv(EnvDatabasePass); v != "" {
		cfg.Database.Password = v
	}

	// Feature flags
	cfg.Features.EnableLegacyAuth = getEnvBool(EnvEnableLegacyAuth, false)
	cfg.Features.EnableNewAuth = getEnvBool(EnvEnableNewAuth, true)
	cfg.Features.EnableV1API = getEnvBool(EnvEnableV1API, true)
	cfg.Features.EnableV2API = getEnvBool(EnvEnableV2API, true)
	cfg.Features.EnableLegacyPayments = getEnvBool(EnvEnableLegacyPayments, false)
	cfg.Features.EnableStripePayments = getEnvBool(EnvEnableStripePayments, true)
	cfg.Features.EnableDebugMode = getEnvBool(EnvEnableDebugMode, false)
	cfg.Features.EnableMetrics = getEnvBool(EnvEnableMetrics, true)

	// Service URLs
	if v := os.Getenv(EnvUsersServiceURL); v != "" {
		cfg.Services.UsersServiceURL = v
	}
	if v := os.Getenv(EnvOrdersServiceURL); v != "" {
		cfg.Services.OrdersServiceURL = v
	}
	if v := os.Getenv(EnvPaymentsServiceURL); v != "" {
		cfg.Services.PaymentsServiceURL = v
	}
	if v := os.Getenv(EnvNotificationsServiceURL); v != "" {
		cfg.Services.NotificationsServiceURL = v
	}

	return cfg
}

func getEnvBool(key string, defaultVal bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}
	v = strings.ToLower(v)
	return v == "true" || v == "1" || v == "yes"
}

func getEnvInt(key string, defaultVal int) int {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultVal
	}
	return i
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return defaultVal
	}
	return d
}
