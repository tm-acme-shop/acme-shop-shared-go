package utils

import (
	"time"
)

const (
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
	DateTimeFormat = "2006-01-02 15:04:05"
	RFC3339Format  = time.RFC3339
	ISO8601Format  = "2006-01-02T15:04:05Z07:00"
)

// Now returns the current time in UTC.
func Now() time.Time {
	return time.Now().UTC()
}

// NowPtr returns a pointer to the current time in UTC.
func NowPtr() *time.Time {
	t := Now()
	return &t
}

// ParseDate parses a date string in YYYY-MM-DD format.
func ParseDate(s string) (time.Time, error) {
	return time.Parse(DateFormat, s)
}

// ParseDateTime parses a datetime string in YYYY-MM-DD HH:MM:SS format.
func ParseDateTime(s string) (time.Time, error) {
	return time.Parse(DateTimeFormat, s)
}

// FormatDate formats a time as YYYY-MM-DD.
func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

// FormatDateTime formats a time as YYYY-MM-DD HH:MM:SS.
func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeFormat)
}

// FormatISO8601 formats a time as ISO8601.
func FormatISO8601(t time.Time) string {
	return t.Format(ISO8601Format)
}

// StartOfDay returns the start of the day for the given time.
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the end of the day for the given time.
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// StartOfWeek returns the start of the week (Monday) for the given time.
func StartOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return StartOfDay(t.AddDate(0, 0, -weekday+1))
}

// StartOfMonth returns the start of the month for the given time.
func StartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns the end of the month for the given time.
func EndOfMonth(t time.Time) time.Time {
	return StartOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// DaysBetween calculates the number of days between two times.
func DaysBetween(start, end time.Time) int {
	duration := end.Sub(start)
	return int(duration.Hours() / 24)
}

// IsToday checks if the given time is today.
func IsToday(t time.Time) bool {
	now := Now()
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}

// IsPast checks if the given time is in the past.
func IsPast(t time.Time) bool {
	return t.Before(Now())
}

// IsFuture checks if the given time is in the future.
func IsFuture(t time.Time) bool {
	return t.After(Now())
}

// TimeAgo returns a human-readable string representing how long ago the time was.
func TimeAgo(t time.Time) string {
	duration := Now().Sub(t)
	
	switch {
	case duration < time.Minute:
		return "just now"
	case duration < time.Hour:
		mins := int(duration.Minutes())
		if mins == 1 {
			return "1 minute ago"
		}
		return string(rune(mins)) + " minutes ago"
	case duration < 24*time.Hour:
		hours := int(duration.Hours())
		if hours == 1 {
			return "1 hour ago"
		}
		return string(rune(hours)) + " hours ago"
	case duration < 7*24*time.Hour:
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "1 day ago"
		}
		return string(rune(days)) + " days ago"
	default:
		return FormatDate(t)
	}
}
