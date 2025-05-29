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
	userID := c.Params("id")
	fmt.Println("USer id :", userID)
	if userID == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "missing user id"})
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	// Check if access token expired, if yes refresh it
	token := &oauth2.Token{
		AccessToken:  user.GoogleAccessToken,
		RefreshToken: user.GoogleRefreshToken,
		Expiry:       user.TokenExpiry,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	config := GetGoogleOauthConfig()

	ts := config.TokenSource(ctx, token)
	newToken, err := ts.Token()
	if err != nil {
		fmt.Println("Token refresh error:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "failed to refresh token"})
	}
	// If token updated, save it back
	if newToken.AccessToken != user.GoogleAccessToken {
		user.GoogleAccessToken = newToken.AccessToken
		user.TokenExpiry = newToken.Expiry
		database.DB.Save(&user)
	}

	// Use the valid access token to call Gmail API
	emailData, err := emails.FetchEmailsHelper(newToken.AccessToken)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Send(emailData)
}
