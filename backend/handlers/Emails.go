package handlers

import (
	"context"
	"fmt"
	"jobbotic-backend/database"
	"jobbotic-backend/models"
	"jobbotic-backend/utils/emails"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type FormattedEmail struct {
	ID      string `json:"id"`
	Subject string `json:"subject"`
	From    string `json:"from"`
	To      string `json:"to"`
	Body    string `json:"body"`
}

func FetchEmails(c *fiber.Ctx) error {
	fmt.Println("Fetch emails handler triggered")
	userID := c.Params("id")
	if userID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "missing user id"})
	}

	var gmailAccount models.GmailAccount
	if err := database.DB.First(&gmailAccount, "user_id = ?", userID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "gmail account not found"})
	}

	token := &oauth2.Token{
		AccessToken:  gmailAccount.GoogleAccessToken,
		RefreshToken: gmailAccount.GoogleRefreshToken,
		Expiry:       gmailAccount.TokenExpiry,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	config := GetGoogleOauthConfig() // make sure you have this function returning *oauth2.Config
	ts := config.TokenSource(ctx, token)
	newToken, err := ts.Token()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to refresh token"})
	}

	if newToken.AccessToken != gmailAccount.GoogleAccessToken {
		gmailAccount.GoogleAccessToken = newToken.AccessToken
		gmailAccount.TokenExpiry = newToken.Expiry
		if err := database.DB.Save(&gmailAccount).Error; err != nil {
			fmt.Println("Failed to update GmailAccount:", err)
		}
	}

	fmt.Println("start fetching ")

	emailListData, err := emails.FetchEmailsHelper(newToken.AccessToken)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var formattedEmails []FormattedEmail
	messages := emailListData.Messages
	if len(messages) > 5 {
		messages = messages[:5]
	}

	for _, msg := range messages {
		gmailMsg, err := emails.FetchMessageDetails(newToken.AccessToken, msg.Id)
		if err != nil {
			fmt.Println("Failed to fetch message details for", msg.Id, ":", err)
			continue
		}
		subject, from, to, body := emails.ExtractEmailDetails(gmailMsg)
		formattedEmails = append(formattedEmails, FormattedEmail{
			ID:      gmailMsg.Id,
			Subject: subject,
			From:    from,
			To:      to,
			Body:    body,
		})
		if emails.IsJobApplicationEmail(subject, body) {
			jobApp := models.JobApplication{
				UserID:  userID, // <- make sure you have this in context
				EmailID: msg.Id,
				Subject: subject,
				From:    from,
				To:      to,
				Snippet: body,
				Date:    time.Now(), // or use msg.InternalDate if available
			}

			database.DB.Create(&jobApp)
		}
	}
	return c.JSON(fiber.Map{
		"messages":           formattedEmails,
		"nextPageToken":      emailListData.NextPageToken,
		"resultSizeEstimate": emailListData.ResultSizeEstimate,
	})
}
