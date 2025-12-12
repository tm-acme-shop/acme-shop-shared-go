package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/tm-acme-shop/acme-shop-shared-go/logging"
)

type contextKey string

const (
	ContextKeyUser contextKey = "user"
)

// AuthConfig holds authentication middleware configuration.
type AuthConfig struct {
	// EnableLegacyAuth enables the legacy authentication mechanism.
	// Deprecated: Set to false and use EnableNewAuth instead.
	EnableLegacyAuth bool

	// EnableNewAuth enables the new JWT-based authentication.
	EnableNewAuth bool

	// JWTSecret is the secret key for JWT validation.
	JWTSecret string

	// SkipPaths are paths that don't require authentication.
	SkipPaths []string
}

// AuthMiddleware provides authentication middleware for HTTP handlers.
type AuthMiddleware struct {
	config AuthConfig
	logger *logging.LoggerV2
}

// NewAuthMiddleware creates a new authentication middleware.
func NewAuthMiddleware(config AuthConfig) *AuthMiddleware {
	return &AuthMiddleware{
		config: config,
		logger: logging.NewLoggerV2("auth-middleware"),
	}
}

// Handler returns an HTTP middleware handler.
func (m *AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, path := range m.config.SkipPaths {
			if strings.HasPrefix(r.URL.Path, path) {
				next.ServeHTTP(w, r)
				return
			}
		}

		var userID string
		var err error

		if m.config.EnableNewAuth {
			userID, err = m.validateNewAuth(r)
		} else if m.config.EnableLegacyAuth {
			// TODO(TEAM-SEC): Remove legacy auth after migration
			userID, err = m.validateLegacyAuth(r)
		}

		if err != nil {
			m.logger.Warn("authentication failed", logging.Fields{
				"path":  r.URL.Path,
				"error": err.Error(),
			})
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyUser, userID)
		ctx = logging.SetUserID(ctx, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// validateNewAuth validates the new JWT-based authentication.
func (m *AuthMiddleware) validateNewAuth(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrMissingAuthHeader
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", ErrInvalidAuthFormat
	}

	token := parts[1]
	
	m.logger.Info("validating JWT token", logging.Fields{
		"token_length": len(token),
	})

	// In a real implementation, this would validate the JWT
	// For demo purposes, we'll just extract a mock user ID
	if len(token) > 10 {
		return "user-" + token[:8], nil
	}

	return "", ErrInvalidToken
}

// validateLegacyAuth validates the legacy API key authentication.
// Deprecated: Use validateNewAuth instead.
func (m *AuthMiddleware) validateLegacyAuth(r *http.Request) (string, error) {
	// TODO(TEAM-SEC): Remove this function after migration to JWT
	apiKey := r.Header.Get("X-API-Key")
	if apiKey == "" {
		apiKey = r.URL.Query().Get("api_key")
	}

	if apiKey == "" {
		return "", ErrMissingAPIKey
	}

	// Legacy: log using old format
	logging.Infof("validating legacy API key: %s...", apiKey[:min(8, len(apiKey))])

	// In a real implementation, this would validate against a database
	if len(apiKey) >= 16 {
		return "legacy-user-" + apiKey[:8], nil
	}

	return "", ErrInvalidAPIKey
}

// GetUserFromContext retrieves the authenticated user ID from context.
func GetUserFromContext(ctx context.Context) string {
	if v := ctx.Value(ContextKeyUser); v != nil {
		return v.(string)
	}
	return ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
