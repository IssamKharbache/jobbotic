
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

    config := GetGoogleOauthConfig()
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
    // 1. fetch list of message IDs
    emailListData, err := emails.FetchEmailsHelper(newToken.AccessToken)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    // 2. fetch details for each message
    var detailedMessages []emails.MessageDetail

messages := emailListData.Messages
if len(messages) > 5 {
    messages = messages[:5] // ✅ this is valid Go syntax
}
fmt.Println("Fetched email list, total:", len(emailListData.Messages))
    for _, msg := range messages.Messages {
        detail, err := emails.FetchMessageDetails(newToken.AccessToken, msg.Id)
        if err != nil {
            fmt.Println("Failed to fetch message details for", msg.Id, ":", err)
            continue
        }
        detailedMessages = append(detailedMessages, *detail)
    }

    // 3. return the detailed messages, not just IDs
    return c.JSON(fiber.Map{
        "messages": detailedMessages,
        "nextPageToken": emailListData.NextPageToken,
        "resultSizeEstimate": emailListData.ResultSizeEstimate,
    })
}
