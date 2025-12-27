package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// GenerateID generates a unique ID with the given prefix.
// Format: prefix_<random hex string>
func GenerateID(prefix string) string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return fmt.Sprintf("%s_%s", prefix, "fallback")
	}
	return fmt.Sprintf("%s_%s", prefix, hex.EncodeToString(b))
}
