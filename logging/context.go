package logging

import (
	"context"
)

type contextKey string

const (
	ContextKeyRequestID contextKey = "request_id"
	ContextKeyUserID    contextKey = "user_id"
	ContextKeyTraceID   contextKey = "trace_id"
	ContextKeySpanID    contextKey = "span_id"
	ContextKeyService   contextKey = "service"
)

// SetRequestID adds a request ID to the context.
func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, ContextKeyRequestID, requestID)
}

// GetRequestID retrieves the request ID from the context.
func GetRequestID(ctx context.Context) string {
	if v := ctx.Value(ContextKeyRequestID); v != nil {
		return v.(string)
	}
	return ""
}

// SetUserID adds a user ID to the context.
func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, ContextKeyUserID, userID)
}

// GetUserID retrieves the user ID from the context.
func GetUserID(ctx context.Context) string {
	if v := ctx.Value(ContextKeyUserID); v != nil {
		return v.(string)
	}
	return ""
}

// SetTraceID adds a trace ID to the context.
func SetTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, ContextKeyTraceID, traceID)
}

// GetTraceID retrieves the trace ID from the context.
func GetTraceID(ctx context.Context) string {
	if v := ctx.Value(ContextKeyTraceID); v != nil {
		return v.(string)
	}
	return ""
}

// SetSpanID adds a span ID to the context.
func SetSpanID(ctx context.Context, spanID string) context.Context {
	return context.WithValue(ctx, ContextKeySpanID, spanID)
}

// GetSpanID retrieves the span ID from the context.
func GetSpanID(ctx context.Context) string {
	if v := ctx.Value(ContextKeySpanID); v != nil {
		return v.(string)
	}
	return ""
}

// ContextFields extracts all logging-relevant fields from a context.
func ContextFields(ctx context.Context) Fields {
	fields := make(Fields)

	if requestID := GetRequestID(ctx); requestID != "" {
		fields["request_id"] = requestID
	}
	if userID := GetUserID(ctx); userID != "" {
		fields["user_id"] = userID
	}
	if traceID := GetTraceID(ctx); traceID != "" {
		fields["trace_id"] = traceID
	}
	if spanID := GetSpanID(ctx); spanID != "" {
		fields["span_id"] = spanID
	}

	return fields
}

// LogFromContext creates a logger with all context fields.
func LogFromContext(ctx context.Context) *LoggerV2 {
	return defaultLoggerV2.WithFields(ContextFields(ctx))
}
