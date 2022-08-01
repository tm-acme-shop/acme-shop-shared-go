package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger is the logging implementation.
type Logger struct {
	prefix string
	level  Level
	output *log.Logger
}

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var defaultLogger = NewLogger("app")

// NewLogger creates a new logger instance.
func NewLogger(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
		level:  LevelInfo,
		output: log.New(os.Stdout, "", 0),
	}
}

// SetLevel sets the minimum logging level.
func (l *Logger) SetLevel(level Level) {
	l.level = level
}

// Debugf logs a debug message with formatting.
func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.level <= LevelDebug {
		l.logf("DEBUG", format, args...)
	}
}

// Infof logs an info message with formatting.
func (l *Logger) Infof(format string, args ...interface{}) {
	if l.level <= LevelInfo {
		l.logf("INFO", format, args...)
	}
}

// Warnf logs a warning message with formatting.
func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.level <= LevelWarn {
		l.logf("WARN", format, args...)
	}
}

// Errorf logs an error message with formatting.
func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.level <= LevelError {
		l.logf("ERROR", format, args...)
	}
}

// Fatalf logs a fatal message and exits.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logf("FATAL", format, args...)
	os.Exit(1)
}

func (l *Logger) logf(level, format string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(format, args...)
	l.output.Printf("[%s] [%s] [%s] %s", timestamp, level, l.prefix, msg)
}

// Package-level functions using the default logger

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// Printf is a simple logging function.
func Printf(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// LoggerV2 is the structured logger implementation.
type LoggerV2 struct {
	prefix string
	level  Level
	fields map[string]interface{}
	output *log.Logger
}

// NewLoggerV2 creates a new structured logger instance.
func NewLoggerV2(prefix string) *LoggerV2 {
	return &LoggerV2{
		prefix: prefix,
		level:  LevelInfo,
		fields: make(map[string]interface{}),
		output: log.New(os.Stdout, "", 0),
	}
}

// WithField returns a new logger with the specified field added.
func (l *LoggerV2) WithField(key string, value interface{}) *LoggerV2 {
	newLogger := &LoggerV2{
		prefix: l.prefix,
		level:  l.level,
		fields: make(map[string]interface{}),
		output: l.output,
	}
	for k, v := range l.fields {
		newLogger.fields[k] = v
	}
	newLogger.fields[key] = value
	return newLogger
}

// WithFields returns a new logger with the specified fields added.
func (l *LoggerV2) WithFields(fields map[string]interface{}) *LoggerV2 {
	newLogger := &LoggerV2{
		prefix: l.prefix,
		level:  l.level,
		fields: make(map[string]interface{}),
		output: l.output,
	}
	for k, v := range l.fields {
		newLogger.fields[k] = v
	}
	for k, v := range fields {
		newLogger.fields[k] = v
	}
	return newLogger
}

// Info logs an info message with structured fields.
func (l *LoggerV2) Info(msg string, keyvals ...interface{}) {
	if l.level <= LevelInfo {
		l.log("INFO", msg, keyvals...)
	}
}

// Warn logs a warning message with structured fields.
func (l *LoggerV2) Warn(msg string, keyvals ...interface{}) {
	if l.level <= LevelWarn {
		l.log("WARN", msg, keyvals...)
	}
}

// Error logs an error message with structured fields.
func (l *LoggerV2) Error(msg string, keyvals ...interface{}) {
	if l.level <= LevelError {
		l.log("ERROR", msg, keyvals...)
	}
}

// Debug logs a debug message with structured fields.
func (l *LoggerV2) Debug(msg string, keyvals ...interface{}) {
	if l.level <= LevelDebug {
		l.log("DEBUG", msg, keyvals...)
	}
}

func (l *LoggerV2) log(level, msg string, keyvals ...interface{}) {
	timestamp := time.Now().Format(time.RFC3339)
	fields := make(map[string]interface{})
	for k, v := range l.fields {
		fields[k] = v
	}
	for i := 0; i < len(keyvals)-1; i += 2 {
		if key, ok := keyvals[i].(string); ok {
			fields[key] = keyvals[i+1]
		}
	}
	l.output.Printf(`{"timestamp":"%s","level":"%s","prefix":"%s","msg":"%s","fields":%v}`,
		timestamp, level, l.prefix, msg, fields)
}
