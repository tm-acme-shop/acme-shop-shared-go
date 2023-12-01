# AcmeShop Shared Go Library

Shared domain models, interfaces, and utilities for AcmeShop Go services.

## Installation

```bash
go get github.com/tm-acme-shop/acme-shop-shared-go
```

## Packages

- `models` - Domain models (User, Order, Payment, Notification)
- `interfaces` - Service interfaces for dependency injection
- `logging` - Structured logging utilities
- `middleware` - HTTP middleware (auth, correlation, recovery)
- `config` - Configuration management
- `errors` - Custom error types
- `utils` - Common utilities

## Usage

```go
import (
    "github.com/tm-acme-shop/acme-shop-shared-go/models"
    "github.com/tm-acme-shop/acme-shop-shared-go/interfaces"
    "github.com/tm-acme-shop/acme-shop-shared-go/logging"
)
```

## Migration Notes

### Logging v2 Migration
We are migrating from `log.Infof()` style to structured `log.Info()` with fields.
See `logging/logger_v2.go` for the new API.

### User API v2 Migration  
`UserV1` is deprecated. Use `User` type instead.
