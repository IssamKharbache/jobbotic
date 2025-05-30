package emails

import (
	"encoding/base64"
	"fmt"
	"strings"

	"google.golang.org/api/gmail/v1"
)

// ExtractEmailDetails extracts subject, from, to, and body text from a Gmail message.
func ExtractEmailDetails(msg *gmail.Message) (subject, from, to, body string) {
	// 1. Extract headers
	for _, header := range msg.Payload.Headers {
		switch header.Name {
		case "Subject":
			subject = header.Value
		case "From":
			from = header.Value
		case "To":
			to = header.Value
		}
	}

	// 2. Extract body (text/plain)
	body = extractPlainTextBody(msg.Payload)

	return
}

// Recursive function to extract the first "text/plain" part
func extractPlainTextBody(part *gmail.MessagePart) string {
	if part == nil {
		return ""
	}

	// If this is the plain text part, decode and return it
	if part.MimeType == "text/plain" && part.Body != nil && part.Body.Data != "" {
		decoded, err := base64.URLEncoding.DecodeString(normalizeBase64(part.Body.Data))
		if err != nil {
			fmt.Println("Error decoding body:", err)
			return ""
		}
		return string(decoded)
	}

	// If this part has nested parts, search recursively
	for _, subPart := range part.Parts {
		text := extractPlainTextBody(subPart)
		if text != "" {
			return text
		}
	}

	return ""
}

// Gmail uses base64url encoding that may include padding removal
func normalizeBase64(data string) string {
	// Replace URL-safe characters
	data = strings.ReplaceAll(data, "-", "+")
	data = strings.ReplaceAll(data, "_", "/")
	// Add padding if missing
	switch len(data) % 4 {
	case 2:
		data += "=="
	case 3:
		data += "="
	}
	return data
}
