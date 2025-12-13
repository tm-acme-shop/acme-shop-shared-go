package utils

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

// DebugTimer is a simple timer for debugging performance.
// Deprecated: Use structured logging with timing fields instead.
type DebugTimer struct {
	name  string
	start time.Time
}

// NewDebugTimer creates a new debug timer.
// Deprecated: Use structured logging with timing fields instead.
// TODO(TEAM-PLATFORM): Remove this, use structured logging
func NewDebugTimer(name string) *DebugTimer {
	log.Printf("Starting timer: %s", name)
	return &DebugTimer{
		name:  name,
		start: time.Now(),
	}
}

// Stop logs the elapsed time.
// Deprecated: Use structured logging with timing fields instead.
func (t *DebugTimer) Stop() time.Duration {
	elapsed := time.Since(t.start)
	// TODO(TEAM-PLATFORM): Migrate to structured logging
	log.Printf("Timer %s completed in %v", t.name, elapsed)
	return elapsed
}

// DebugLog logs a debug message with caller information.
// Deprecated: Use LoggerV2.Debug instead.
// TODO(TEAM-PLATFORM): Remove this function
func DebugLog(format string, args ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf(format, args...)
	log.Printf("[DEBUG] %s:%d - %s", file, line, msg)
}

// TraceFunction logs function entry and exit.
// Deprecated: Use structured logging with tracing instead.
// TODO(TEAM-PLATFORM): Remove this, use OpenTelemetry
func TraceFunction(name string) func() {
	log.Printf("Entering function: %s", name)
	start := time.Now()
	return func() {
		log.Printf("Exiting function: %s (took %v)", name, time.Since(start))
	}
}

// LogMemStats logs memory statistics.
// Deprecated: Use metrics instead of logging.
// TODO(TEAM-PLATFORM): Move to Prometheus metrics
func LogMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Memory: Alloc=%v MiB, TotalAlloc=%v MiB, Sys=%v MiB, NumGC=%v",
		m.Alloc/1024/1024,
		m.TotalAlloc/1024/1024,
		m.Sys/1024/1024,
		m.NumGC,
	)
}
