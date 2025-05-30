package emails

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"google.golang.org/api/gmail/v1"
)

// FetchEmailsHelper fetches the list of messages (IDs only)
type EmailListData struct {
	Messages           []gmail.Message `json:"messages"`
	NextPageToken      string          `json:"nextPageToken"`
	ResultSizeEstimate int             `json:"resultSizeEstimate"`
}

func FetchEmailsHelper(accessToken string) (*EmailListData, error) {
	req, err := http.NewRequest("GET", "https://gmail.googleapis.com/gmail/v1/users/me/messages?q=in:sent", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result EmailListData
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FetchMessageDetails fetches the full message details for a single message ID
func FetchMessageDetails(accessToken, messageID string) (*gmail.Message, error) {
	url := fmt.Sprintf("https://gmail.googleapis.com/gmail/v1/users/me/messages/%s?format=full", messageID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var msg gmail.Message
	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

// ExtractEmailDetails extracts subject, from, to, and body text from a Gmail message.
func ExtractEmailDetails(msg *gmail.Message) (subject, from, to, body string) {
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
	body = extractPlainTextBody(msg.Payload)
	return
}

func extractPlainTextBody(part *gmail.MessagePart) string {
	if part == nil || part.Body == nil || part.Body.Data == "" {
		return ""
	}

	data := normalizeBase64(part.Body.Data)
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		decoded, err = base64.RawURLEncoding.DecodeString(data)
		if err != nil {
			fmt.Println("Final fallback: Error decoding body:", err)
			return ""
		}
	}

	return string(decoded)
}

// Recursive function to extract the first "text/plain" part
func normalizeBase64(data string) string {
	data = strings.ReplaceAll(data, "-", "+")
	data = strings.ReplaceAll(data, "_", "/")
	switch len(data) % 4 {
	case 2:
		data += "=="
	case 3:
		data += "="
	}
	return data
}
