package handlers

import (
	"context"
	"encoding/json"
	"jobbotic-backend/database"
	"jobbotic-backend/models"
	"jobbotic-backend/utils"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

// GetGetGoogleOauthConfig() returns a new OAuth2 config for Google
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

// generateStateOauthCookie generates a random state string and sets it as a cookie
func GenerateStateOauthCookie(c *fiber.Ctx) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	state := string(b)

	c.Cookie(&fiber.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  time.Now().Add(1 * time.Hour),
		HTTPOnly: true,
	})

	return state
}

// GoogleLogin redirects user to Google's OAuth consent page

func GoogleLogin(c *fiber.Ctx) error {
	state := GenerateStateOauthCookie(c)
	url := GetGoogleOauthConfig().AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

// GoogleCallback handles the callback from Google after user consents
func GoogleCallback(c *fiber.Ctx) error {
	state := c.Cookies("oauthstate")
	if c.Query("state") != state {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid OAuth state"})
	}

	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Code not found"})
	}

	token, err := GetGoogleOauthConfig().Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange token"})
	}

	// TODO: Save token.AccessToken, token.RefreshToken, token.Expiry in DB linked to the user

	return c.JSON(fiber.Map{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"expiry":        token.Expiry,
	})
}

//Google sign up handler

func GoogleSignup(c *fiber.Ctx) error {
	state := c.Cookies("oauthstate")
	if c.Query("state") != state {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid OAuth state"})
	}

	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Code not found"})
	}

	token, err := GetGoogleOauthConfig().Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange token"})
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
	} else if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to query user"})
	} else {
		// Existing user: optionally update token info
		user.GoogleAccessToken = token.AccessToken
		user.TokenExpiry = token.Expiry
		if token.RefreshToken != "" {
			user.GoogleRefreshToken = token.RefreshToken
		}
		db.Save(&user)
	}

	// Generate JWT token
	jwtToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
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

func GoogleLink(c *fiber.Ctx) error {
	// Generate a secure random state
	state := uuid.New().String()

	c.Cookie(&fiber.Cookie{
		Name:     "oauthstate",
		Value:    state,
		HTTPOnly: true,
		Secure:   true, // important for production
		Path:     "/",
	})

	// Generate URL
	url := GetGoogleOauthConfig().AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)

	// Return the URL instead of redirecting
	return c.JSON(fiber.Map{
		"authUrl": url,
	})
}

func GoogleLinkCallback(c *fiber.Ctx) error {
	state := c.Cookies("oauthstate")
	if c.Query("state") != state {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid OAuth state"})
	}

	code := c.Query("code")
	token, err := GetGoogleOauthConfig().Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange token"})
	}

	client := GetGoogleOauthConfig().Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user info"})
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to decode user info"})
	}

	// Get current user ID from token (requires auth middleware)
	userID := c.Locals("userID").(uint)

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
