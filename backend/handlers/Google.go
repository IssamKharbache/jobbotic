
package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"jobbotic-backend/database"
	"jobbotic-backend/models"
	"jobbotic-backend/utils"
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

// OAuth config
func GetGoogleOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/gmail.readonly",
		},
		Endpoint: google.Endpoint,
	}
}

// Redirect to Google login for signup
func GoogleLogin(c *fiber.Ctx) error {
	state := generateState() // random fallback if not using frontend state
	c.Cookie(&fiber.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  time.Now().Add(1 * time.Hour),
		HTTPOnly: true,
	})
	url := GetGoogleOauthConfig().AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

// Google Signup/Login Callback
func GoogleCallback(c *fiber.Ctx) error {
	// Decode base64 state (containing JWT)
	stateParam := c.Query("state")
	if stateParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing state parameter"})
	}

	decodedState, err := base64.StdEncoding.DecodeString(stateParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid state encoding"})
	}

	var stateData struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(decodedState, &stateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid state format"})
	}

	// Validate JWT
	claims, err := utils.ValidateToken(stateData.Token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
	}

	code := c.Query("code")
	token, err := GetGoogleOauthConfig().Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange code"})
	}

	client := GetGoogleOauthConfig().Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info"})
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email      string `json:"email"`
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to decode user info"})
	}

	db := database.DB
	var user models.User
	result := db.Where("email = ?", userInfo.Email).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		user = models.User{
			FirstName:          userInfo.GivenName,
			LastName:           userInfo.FamilyName,
			Email:              userInfo.Email,
			GoogleAccessToken:  token.AccessToken,
			GoogleRefreshToken: token.RefreshToken,
			TokenExpiry:        token.Expiry,
			IsGmailLinked:      true,
		}
		if err := db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
		}
	} else if result.Error == nil {
		user.GoogleAccessToken = token.AccessToken
		user.TokenExpiry = token.Expiry
		if token.RefreshToken != "" {
			user.GoogleRefreshToken = token.RefreshToken
		}
		user.IsGmailLinked = true
		db.Save(&user)
	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to query user"})
	}

	// Generate JWT token
	jwtToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate JWT"})
	}

	return c.JSON(fiber.Map{
		"token": jwtToken,
		"user": fiber.Map{
			"id":         user.ID,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
		},
	})
}

// Generate Gmail Link URL
func GoogleLink(c *fiber.Ctx) error {
	state := c.Query("state")
	if state == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing state"})
	}

	authURL := fmt.Sprintf(
		"https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=%s&prompt=consent&redirect_uri=%s&response_type=code&scope=%s&state=%s",
		os.Getenv("GOOGLE_CLIENT_ID"),
		url.QueryEscape(os.Getenv("GOOGLE_REDIRECT_URL")+"/link"),
		url.QueryEscape("https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/gmail.readonly"),
		state,
	)

	return c.JSON(fiber.Map{"url": authURL})
}

// Gmail Link Callback
func GoogleLinkCallback(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	code := c.Query("code")
	token, err := GetGoogleOauthConfig().Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange token"})
	}

	db := database.DB
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	user.GoogleAccessToken = token.AccessToken
	user.GoogleRefreshToken = token.RefreshToken
	user.TokenExpiry = token.Expiry
	user.IsGmailLinked = true

	if err := db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save tokens"})
	}

	return c.JSON(fiber.Map{"message": "Gmail successfully linked!"})
}

// Helper
func generateState() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
