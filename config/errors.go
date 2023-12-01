package config

import "errors"

var (
	ErrMissingServiceName = errors.New("service name is required")
	ErrInvalidPort        = errors.New("invalid port number")
	ErrMissingDatabase    = errors.New("database configuration is required")
	ErrInvalidLogLevel    = errors.New("invalid log level")
)
