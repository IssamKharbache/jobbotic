package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateState creates a secure random string for OAuth state parameter
func GenerateState() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		// fallback or handle error properly
		return "default_state"
	}
	return base64.URLEncoding.EncodeToString(b)
}
