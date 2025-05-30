package emails

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EmailListData struct {
	Messages           []Message `json:"messages"`
	NextPageToken      string    `json:"nextPageToken"`
	ResultSizeEstimate int       `json:"resultSizeEstimate"`
}

type Message struct {
	Id string `json:"id"`
}
type MessageDetail struct {
	Id       string         `json:"id"`
	ThreadId string         `json:"threadId"`
	Snippet  string         `json:"snippet"`
	Payload  map[string]any `json:"payload"`
	// Add other fields as you need from the Gmail API message resource
}

func FetchMessageDetails(accessToken, messageID string) (*MessageDetail, error) {
	url := fmt.Sprintf("https://gmail.googleapis.com/gmail/v1/users/me/messages/%s", messageID)

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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var detail MessageDetail
	if err := json.Unmarshal(body, &detail); err != nil {
		return nil, err
	}

	return &detail, nil

}

func FetchEmailsHelper(accessToken string) (*EmailListData, error) {
	// Add query "in:sent" to get only sent emails
	req, _ := http.NewRequest("GET", "https://gmail.googleapis.com/gmail/v1/users/me/messages?q=in:sent", nil)
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
