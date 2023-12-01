package utils

import (
	"regexp"
	"strings"
	"unicode"
)

// Truncate truncates a string to the specified length.
func Truncate(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length]
}

// TruncateWithEllipsis truncates a string and adds ellipsis if needed.
func TruncateWithEllipsis(s string, length int) string {
	if len(s) <= length {
		return s
	}
	if length <= 3 {
		return s[:length]
	}
	return s[:length-3] + "..."
}

// Slugify converts a string to a URL-friendly slug.
func Slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.TrimSpace(s)
	
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "-")
	
	s = strings.Trim(s, "-")
	return s
}

// CamelToSnake converts camelCase to snake_case.
func CamelToSnake(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// SnakeToCamel converts snake_case to camelCase.
func SnakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}

// MaskEmail masks an email address for display.
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "****"
	}
	
	name := parts[0]
	domain := parts[1]
	
	if len(name) <= 2 {
		return name + "****@" + domain
	}
	
	return name[:2] + strings.Repeat("*", len(name)-2) + "@" + domain
}

// MaskPhone masks a phone number for display.
func MaskPhone(phone string) string {
	digits := regexp.MustCompile(`\D`).ReplaceAllString(phone, "")
	if len(digits) < 4 {
		return "****"
	}
	return strings.Repeat("*", len(digits)-4) + digits[len(digits)-4:]
}

// ContainsAny checks if a string contains any of the given substrings.
func ContainsAny(s string, substrs ...string) bool {
	for _, substr := range substrs {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// IsEmpty checks if a string is empty or contains only whitespace.
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// DefaultIfEmpty returns the default value if the string is empty.
func DefaultIfEmpty(s, defaultVal string) string {
	if IsEmpty(s) {
		return defaultVal
	}
	return s
}
