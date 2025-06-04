package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"jobbotic-backend/database"
	"jobbotic-backend/models"
	"jobbotic-backend/utils"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

func GoogleLogin(c *fiber.Ctx) error {
	state := utils.GenerateState() // Generate a random state for CSRF protection
	url := GetGoogleOauthConfig().AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("prompt", "consent"))
	return c.Redirect(url)
}

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
	userIDVal := c.Locals("userID")
	userID, ok := userIDVal.(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User ID not found in context or invalid",
		})
	}
	state := utils.GenerateStateWithUserID(userID)
	authURL := GetGoogleOauthConfig().AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.SetAuthURLParam("prompt", "consent"))

	return c.JSON(fiber.Map{
		"url": authURL,
	})
}

// Gmail Link Callback

func GoogleLinkCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing state parameter",
		})
	}

	decodedBytes, err := base64.URLEncoding.DecodeString(state)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid state encoding",
		})
	}

	var stateData map[string]string
	if err := json.Unmarshal(decodedBytes, &stateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid state format",
		})
	}

	userID := stateData["uid"]
	if userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing user ID in state",
		})
	}

	code := c.Query("code")
	if code == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing authorization code",
		})
	}

	// Exchange code for token
	token, err := GetGoogleOauthConfig().Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to exchange code for token",
		})
	}

	// Fetch user's Gmail address
	client := GetGoogleOauthConfig().Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil || resp.StatusCode != 200 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user info from Google",
		})
	}
	defer resp.Body.Close()

	var googleUser struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to decode Google user info",
		})
	}

	// Fetch user from DB
	var user models.User
	if err := database.DB.First(&user, "id = ?", userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// ✅ Update User basic flags
	user.IsGmailLinked = true
	user.LinkedEmail = &googleUser.Email // optional field in User
	user.GoogleAccessToken = token.AccessToken
	user.GoogleRefreshToken = token.RefreshToken
	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user info",
		})
	}

	// ✅ Create or update GmailAccount record
	gmailAccount := models.GmailAccount{
		UserID:             user.ID,
		Email:              googleUser.Email,
		GoogleAccessToken:  token.AccessToken,
		GoogleRefreshToken: token.RefreshToken,
		TokenExpiry:        token.Expiry,
	}

	// Check if Gmail account already exists
	var existing models.GmailAccount
	err = database.DB.Where("email = ?", googleUser.Email).First(&existing).Error
	if err == nil {
		// Update existing
		existing.GoogleAccessToken = token.AccessToken
		existing.GoogleRefreshToken = token.RefreshToken
		existing.TokenExpiry = token.Expiry
		database.DB.Save(&existing)
	} else {
		// Create new
		database.DB.Create(&gmailAccount)
	}

	return c.JSON(fiber.Map{
		"message": "Gmail linked successfully",
	})
}

func GetValidAccessToken(account *models.GmailAccount) (string, error) {
	if time.Now().Before(account.TokenExpiry) {
		return account.GoogleAccessToken, nil
	}

	config := GetGoogleOauthConfig()
	token := &oauth2.Token{
		RefreshToken: account.GoogleRefreshToken,
	}
	tokenSource := config.TokenSource(context.Background(), token)
	newToken, err := tokenSource.Token()
	if err != nil {
		return "", err
	}

	// Update DB
	account.GoogleAccessToken = newToken.AccessToken
	account.TokenExpiry = newToken.Expiry
	database.DB.Save(account)

	return newToken.AccessToken, nil
}

func TestAccessToken(c *fiber.Ctx) error {
	userID := c.Params("id") // or get from auth session
	var account models.GmailAccount
	err := database.DB.Where("user_id = ?", userID).First(&account).Error
	if err != nil {
		return c.Status(404).SendString("Gmail account not found")
	}

	token, err := GetValidAccessToken(&account)
	if err != nil {
		return c.Status(500).SendString("Error: " + err.Error())
	}

	return c.SendString("Access Token: " + token)
}
