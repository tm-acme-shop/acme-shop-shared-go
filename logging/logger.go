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
