package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

// LoggerV2 is the new structured logger implementation.
// This replaces the legacy Logger with proper structured logging.
type LoggerV2 struct {
	mu      sync.Mutex
	output  io.Writer
	level   Level
	fields  Fields
	service string
}

// Fields represents structured log fields.
type Fields map[string]interface{}

// LogEntry represents a structured log entry.
type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Service   string                 `json:"service"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Caller    string                 `json:"caller,omitempty"`
}

var defaultLoggerV2 = NewLoggerV2("app")

// NewLoggerV2 creates a new structured logger instance.
func NewLoggerV2(service string) *LoggerV2 {
	return &LoggerV2{
		output:  os.Stdout,
		level:   LevelInfo,
		fields:  make(Fields),
		service: service,
	}
}

// WithField returns a new logger with the specified field added.
func (l *LoggerV2) WithField(key string, value interface{}) *LoggerV2 {
	newFields := make(Fields, len(l.fields)+1)
	for k, v := range l.fields {
		newFields[k] = v
	}
	newFields[key] = value

	return &LoggerV2{
		output:  l.output,
		level:   l.level,
		fields:  newFields,
		service: l.service,
	}
}

// WithFields returns a new logger with the specified fields added.
func (l *LoggerV2) WithFields(fields Fields) *LoggerV2 {
	newFields := make(Fields, len(l.fields)+len(fields))
	for k, v := range l.fields {
		newFields[k] = v
	}
	for k, v := range fields {
		newFields[k] = v
	}

	return &LoggerV2{
		output:  l.output,
		level:   l.level,
		fields:  newFields,
		service: l.service,
	}
}

// WithError returns a new logger with an error field added.
func (l *LoggerV2) WithError(err error) *LoggerV2 {
	return l.WithField("error", err.Error())
}

// Debug logs a debug message with structured fields.
func (l *LoggerV2) Debug(msg string, fields ...Fields) {
	if l.level <= LevelDebug {
		l.log("DEBUG", msg, mergeFields(fields...))
	}
}

// Info logs an info message with structured fields.
func (l *LoggerV2) Info(msg string, fields ...Fields) {
	if l.level <= LevelInfo {
		l.log("INFO", msg, mergeFields(fields...))
	}
}

// Warn logs a warning message with structured fields.
func (l *LoggerV2) Warn(msg string, fields ...Fields) {
	if l.level <= LevelWarn {
		l.log("WARN", msg, mergeFields(fields...))
	}
}

// Error logs an error message with structured fields.
func (l *LoggerV2) Error(msg string, fields ...Fields) {
	if l.level <= LevelError {
		l.log("ERROR", msg, mergeFields(fields...))
	}
}

// Fatal logs a fatal message and exits.
func (l *LoggerV2) Fatal(msg string, fields ...Fields) {
	l.log("FATAL", msg, mergeFields(fields...))
	os.Exit(1)
}

func (l *LoggerV2) log(level, msg string, additionalFields Fields) {
	l.mu.Lock()
	defer l.mu.Unlock()

	allFields := make(Fields, len(l.fields)+len(additionalFields))
	for k, v := range l.fields {
		allFields[k] = v
	}
	for k, v := range additionalFields {
		allFields[k] = v
	}

	_, file, line, ok := runtime.Caller(2)
	caller := ""
	if ok {
		caller = fmt.Sprintf("%s:%d", file, line)
	}

	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     level,
		Service:   l.service,
		Message:   msg,
		Fields:    allFields,
		Caller:    caller,
	}

	data, _ := json.Marshal(entry)
	l.output.Write(append(data, '\n'))
}

func mergeFields(fields ...Fields) Fields {
	result := make(Fields)
	for _, f := range fields {
		for k, v := range f {
			result[k] = v
		}
	}
	return result
}

// Package-level functions using the default V2 logger

// Debug logs a debug message using the default logger.
func Debug(msg string, fields ...Fields) {
	defaultLoggerV2.Debug(msg, fields...)
}

// Info logs an info message using the default logger.
func Info(msg string, fields ...Fields) {
	defaultLoggerV2.Info(msg, fields...)
}

// Warn logs a warning message using the default logger.
func Warn(msg string, fields ...Fields) {
	defaultLoggerV2.Warn(msg, fields...)
}

// Error logs an error message using the default logger.
func Error(msg string, fields ...Fields) {
	defaultLoggerV2.Error(msg, fields...)
}

// Fatal logs a fatal message using the default logger and exits.
func Fatal(msg string, fields ...Fields) {
	defaultLoggerV2.Fatal(msg, fields...)
}

// WithContext returns a logger with context fields (request ID, user ID, etc.)
func WithContext(ctx context.Context) *LoggerV2 {
	logger := defaultLoggerV2

	if requestID := ctx.Value(ContextKeyRequestID); requestID != nil {
		logger = logger.WithField("request_id", requestID)
	}
	if userID := ctx.Value(ContextKeyUserID); userID != nil {
		logger = logger.WithField("user_id", userID)
	}

	return logger
}
