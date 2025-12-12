package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/tm-acme-shop/acme-shop-shared-go/logging"
)

// RecoveryMiddleware recovers from panics and logs them.
type RecoveryMiddleware struct {
	logger *logging.LoggerV2
}

// NewRecoveryMiddleware creates a new recovery middleware.
func NewRecoveryMiddleware() *RecoveryMiddleware {
	return &RecoveryMiddleware{
		logger: logging.NewLoggerV2("recovery-middleware"),
	}
}

// Handler returns an HTTP middleware handler.
func (m *RecoveryMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				stack := debug.Stack()

				m.logger.Error("panic recovered", logging.Fields{
					"error":      fmt.Sprintf("%v", err),
					"stack":      string(stack),
					"method":     r.Method,
					"path":       r.URL.Path,
					"request_id": logging.GetRequestID(r.Context()),
				})

				// Also log with legacy format for backwards compatibility
				// TODO(TEAM-PLATFORM): Remove legacy logging after migration
				logging.Errorf("PANIC: %v\n%s", err, stack)

				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// ResponseWriter wraps http.ResponseWriter to capture status code.
type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Written    int64
}

// NewResponseWriter creates a wrapped response writer.
func NewResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}

// WriteHeader captures the status code.
func (rw *ResponseWriter) WriteHeader(code int) {
	rw.StatusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// Write captures bytes written.
func (rw *ResponseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.Written += int64(n)
	return n, err
}
