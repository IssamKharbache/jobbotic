package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateState creates a secure random string for OAuth state parameter
func GenerateState() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}
