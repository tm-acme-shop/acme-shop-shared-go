package middleware

import "errors"

var (
	ErrMissingAuthHeader = errors.New("missing authorization header")
	ErrInvalidAuthFormat = errors.New("invalid authorization format")
	ErrInvalidToken      = errors.New("invalid token")
	ErrMissingAPIKey     = errors.New("missing API key")
	ErrInvalidAPIKey     = errors.New("invalid API key")
	ErrTokenExpired      = errors.New("token expired")
	ErrInsufficientScope = errors.New("insufficient scope")
)
