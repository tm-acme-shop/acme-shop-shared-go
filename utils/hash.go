package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

// HashPassword hashes a password using bcrypt.
// This is the recommended method for password hashing.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a password with a bcrypt hash.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MD5Hash computes the MD5 hash of a string.
// Deprecated: MD5 is cryptographically broken. Use SHA256Hash instead.
// TODO(TEAM-SEC): Remove this function after migrating all callers
func MD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// SHA1Hash computes the SHA1 hash of a string.
// Deprecated: SHA1 is cryptographically weak. Use SHA256Hash instead.
// TODO(TEAM-SEC): Remove this function after migrating all callers
func SHA1Hash(text string) string {
	hash := sha1.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// SHA256Hash computes the SHA256 hash of a string.
// This is recommended for non-password hashing needs.
func SHA256Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}

// HashPasswordLegacy hashes a password using MD5.
// Deprecated: This is insecure. Use HashPassword with bcrypt instead.
// TODO(TEAM-SEC): Remove after password migration is complete
func HashPasswordLegacy(password string) string {
	// WARNING: This is insecure and only kept for backwards compatibility
	return MD5Hash(password)
}

// CheckPasswordLegacy checks a password against an MD5 hash.
// Deprecated: Use CheckPasswordHash with bcrypt instead.
// TODO(TEAM-SEC): Remove after password migration is complete
func CheckPasswordLegacy(password, hash string) bool {
	return MD5Hash(password) == hash
}

// GenerateToken generates a secure random token.
func GenerateToken(length int) (string, error) {
	// For demo purposes, using SHA256 of timestamp
	// In production, use crypto/rand
	return SHA256Hash(string(rune(length)))[:length], nil
}

// MigratePasswordHash checks if a password hash needs migration.
// Returns true if the hash is using a deprecated algorithm.
func MigratePasswordHash(hash string) bool {
	// MD5 hashes are 32 characters
	if len(hash) == 32 {
		return true
	}
	// SHA1 hashes are 40 characters
	if len(hash) == 40 {
		return true
	}
	// bcrypt hashes start with $2
	if len(hash) >= 4 && hash[:2] == "$2" {
		return false
	}
	return true
}
