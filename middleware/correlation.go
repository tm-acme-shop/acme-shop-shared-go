package middleware

import (
	"net/http"

	"github.com/tm-acme-shop/acme-shop-shared-go/logging"
)

const (
	// HeaderRequestID is the standard request ID header.
	HeaderRequestID = "X-Acme-Request-ID"

	// HeaderLegacyUserID is the legacy user ID header.
	// Deprecated: Use HeaderUserID instead.
	HeaderLegacyUserID = "X-Legacy-User-Id"

	// HeaderUserID is the new user ID header.
	HeaderUserID = "X-User-Id"

	// HeaderTraceID is the distributed tracing header.
	HeaderTraceID = "X-Trace-ID"
)

// CorrelationMiddleware adds correlation IDs to requests for distributed tracing.
type CorrelationMiddleware struct {
	logger *logging.LoggerV2
}

// NewCorrelationMiddleware creates a new correlation middleware.
func NewCorrelationMiddleware() *CorrelationMiddleware {
	return &CorrelationMiddleware{
		logger: logging.NewLoggerV2("correlation-middleware"),
	}
}

// Handler returns an HTTP middleware handler.
func (m *CorrelationMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Extract or generate request ID
		requestID := r.Header.Get(HeaderRequestID)
		if requestID == "" {
			requestID = generateRequestID()
		}
		ctx = logging.SetRequestID(ctx, requestID)
		w.Header().Set(HeaderRequestID, requestID)

		// Extract trace ID if present
		traceID := r.Header.Get(HeaderTraceID)
		if traceID != "" {
			ctx = logging.SetTraceID(ctx, traceID)
		}

		// Handle user ID from headers
		userID := m.extractUserID(r)
		if userID != "" {
			ctx = logging.SetUserID(ctx, userID)
		}

		m.logger.Info("request started", logging.Fields{
			"request_id": requestID,
			"method":     r.Method,
			"path":       r.URL.Path,
			"user_id":    userID,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// extractUserID extracts user ID from headers, supporting both legacy and new headers.
func (m *CorrelationMiddleware) extractUserID(r *http.Request) string {
	// Try new header first
	userID := r.Header.Get(HeaderUserID)
	if userID != "" {
		return userID
	}

	// Fall back to legacy header
	// TODO(TEAM-PLATFORM): Remove legacy header support after migration
	legacyUserID := r.Header.Get(HeaderLegacyUserID)
	if legacyUserID != "" {
		// Log usage of deprecated header
		logging.Warnf("legacy header %s used, migrate to %s", HeaderLegacyUserID, HeaderUserID)
		return legacyUserID
	}

	return ""
}

// generateRequestID generates a unique request ID.
func generateRequestID() string {
	// Simple implementation for demo purposes
	// In production, use UUID or similar
	return "req-" + randomString(16)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[i%len(letters)]
	}
	return string(b)
}

// PropagateHeaders copies correlation headers to outgoing requests.
func PropagateHeaders(r *http.Request, outgoing *http.Request) {
	if requestID := r.Header.Get(HeaderRequestID); requestID != "" {
		outgoing.Header.Set(HeaderRequestID, requestID)
	}
	if traceID := r.Header.Get(HeaderTraceID); traceID != "" {
		outgoing.Header.Set(HeaderTraceID, traceID)
	}
	
	// Propagate user ID using new header
	if userID := r.Header.Get(HeaderUserID); userID != "" {
		outgoing.Header.Set(HeaderUserID, userID)
	} else if legacyUserID := r.Header.Get(HeaderLegacyUserID); legacyUserID != "" {
		// TODO(TEAM-PLATFORM): Migrate to new header format
		outgoing.Header.Set(HeaderUserID, legacyUserID)
	}
}
