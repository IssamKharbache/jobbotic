package utils

import (
	"encoding/base64"
	"encoding/json"
)

func GenerateStateWithUserID(userID string) string {
	stateObj := map[string]string{
		"uid": userID,
	}
	jsonState, _ := json.Marshal(stateObj)
	return base64.URLEncoding.EncodeToString(jsonState)
}
