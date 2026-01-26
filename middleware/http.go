package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tm-acme-shop/acme-shop-shared-go/logging"
)

type requestIDKey string

// RequestIDKey is the context key for request IDs.
const RequestIDKey requestIDKey = "request_id"

// RequestIDMiddleware adds a unique request ID to each request.
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader(HeaderRequestID)
		if requestID == "" {
			requestID = generateRequestID()
		}
		c.Set(string(RequestIDKey), requestID)
		c.Request = c.Request.WithContext(
			logging.SetRequestID(c.Request.Context(), requestID),
		)
		c.Header(HeaderRequestID, requestID)
		c.Next()
	}
}

// LoggingMiddleware logs HTTP requests.
func LoggingMiddleware() gin.HandlerFunc {
	logger := logging.NewLoggerV2("http")
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		logger.Info("request completed", logging.Fields{
			"method":   c.Request.Method,
			"path":     path,
			"status":   c.Writer.Status(),
			"duration": time.Since(start).String(),
		})
	}
}

// CORSMiddleware handles Cross-Origin Resource Sharing.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Request-ID")
		c.Header("Access-Control-Expose-Headers", "X-Request-ID")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
