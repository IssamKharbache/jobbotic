package emails

import (
	"fmt"
	"io"
	"net/http"
)

func FetchEmailsHelper(accessToken string) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://gmail.googleapis.com/gmail/v1/users/me/messages?maxResults=5", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("Gmail API error: %s", body)
	}

	return io.ReadAll(res.Body)
}
