package errors

import (
	"errors"
	"fmt"
	"net/http"
)

// Standard error types
var (
	ErrNotFound       = errors.New("resource not found")
	ErrAlreadyExists  = errors.New("resource already exists")
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrInternal       = errors.New("internal error")
	ErrTimeout        = errors.New("operation timed out")
	ErrRateLimited    = errors.New("rate limit exceeded")
	ErrServiceDown    = errors.New("service unavailable")
)

// AppError represents an application error with additional context.
type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
	StatusCode int    `json:"-"`
	Err        error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// NewNotFoundError creates a not found error.
func NewNotFoundError(resource, id string) *AppError {
	return &AppError{
		Code:       "NOT_FOUND",
		Message:    fmt.Sprintf("%s not found", resource),
		Details:    fmt.Sprintf("%s with ID '%s' does not exist", resource, id),
		StatusCode: http.StatusNotFound,
		Err:        ErrNotFound,
	}
}

// NewValidationError creates a validation error.
func NewValidationError(field, message string) *AppError {
	return &AppError{
		Code:       "VALIDATION_ERROR",
		Message:    fmt.Sprintf("invalid %s", field),
		Details:    message,
		StatusCode: http.StatusBadRequest,
		Err:        ErrInvalidInput,
	}
}

// NewUnauthorizedError creates an unauthorized error.
func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:       "UNAUTHORIZED",
		Message:    "authentication required",
		Details:    message,
		StatusCode: http.StatusUnauthorized,
		Err:        ErrUnauthorized,
	}
}

// NewForbiddenError creates a forbidden error.
func NewForbiddenError(message string) *AppError {
	return &AppError{
		Code:       "FORBIDDEN",
		Message:    "access denied",
		Details:    message,
		StatusCode: http.StatusForbidden,
		Err:        ErrForbidden,
	}
}

// NewConflictError creates a conflict error.
func NewConflictError(resource, field, value string) *AppError {
	return &AppError{
		Code:       "CONFLICT",
		Message:    fmt.Sprintf("%s already exists", resource),
		Details:    fmt.Sprintf("%s with %s '%s' already exists", resource, field, value),
		StatusCode: http.StatusConflict,
		Err:        ErrAlreadyExists,
	}
}

// NewInternalError creates an internal error.
func NewInternalError(err error) *AppError {
	return &AppError{
		Code:       "INTERNAL_ERROR",
		Message:    "an internal error occurred",
		StatusCode: http.StatusInternalServerError,
		Err:        err,
	}
}

// NewTimeoutError creates a timeout error.
func NewTimeoutError(operation string) *AppError {
	return &AppError{
		Code:       "TIMEOUT",
		Message:    "operation timed out",
		Details:    fmt.Sprintf("%s did not complete in time", operation),
		StatusCode: http.StatusGatewayTimeout,
		Err:        ErrTimeout,
	}
}

// NewRateLimitError creates a rate limit error.
func NewRateLimitError() *AppError {
	return &AppError{
		Code:       "RATE_LIMITED",
		Message:    "too many requests",
		Details:    "please try again later",
		StatusCode: http.StatusTooManyRequests,
		Err:        ErrRateLimited,
	}
}

// NewServiceUnavailableError creates a service unavailable error.
func NewServiceUnavailableError(service string) *AppError {
	return &AppError{
		Code:       "SERVICE_UNAVAILABLE",
		Message:    "service temporarily unavailable",
		Details:    fmt.Sprintf("%s is currently unavailable", service),
		StatusCode: http.StatusServiceUnavailable,
		Err:        ErrServiceDown,
	}
}

// ValidationError is a type alias for validation errors.
// Deprecated: Use AppError with NewValidationError instead.
type ValidationError = AppError

// IsNotFound checks if an error is a not found error.
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// IsValidationError checks if an error is a validation error.
func IsValidationError(err error) bool {
	return errors.Is(err, ErrInvalidInput)
}

// IsUnauthorized checks if an error is an unauthorized error.
func IsUnauthorized(err error) bool {
	return errors.Is(err, ErrUnauthorized)
}

// IsForbidden checks if an error is a forbidden error.
func IsForbidden(err error) bool {
	return errors.Is(err, ErrForbidden)
}

// GetStatusCode returns the HTTP status code for an error.
func GetStatusCode(err error) int {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr.StatusCode
	}
	if IsNotFound(err) {
		return http.StatusNotFound
	}
	if IsValidationError(err) {
		return http.StatusBadRequest
	}
	if IsUnauthorized(err) {
		return http.StatusUnauthorized
	}
	if IsForbidden(err) {
		return http.StatusForbidden
	}
	return http.StatusInternalServerError
}
