package emails

import (
	"strings"
)

var applicationKeywords = []string{
	// English
	"application", "applied", "position", "resume", "cv", "job",
	// French
	"candidature", "poste", "cv", "emploi",
	// German
	"bewerbung", "stelle", "job", "lebenslauf",
}

func IsJobApplicationEmail(subject, body string) bool {
	content := strings.ToLower(subject + " " + body)
	for _, keyword := range applicationKeywords {
		if strings.Contains(content, keyword) {
			return true
		}
	}
	return false
}
