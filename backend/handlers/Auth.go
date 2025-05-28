package handlers

import (
	"jobbotic-backend/database"
	"jobbotic-backend/models"
	"jobbotic-backend/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// RegisterSchema struct
type RegisterSchema struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}

// LoginSchema struct
type LoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var registerData RegisterSchema
	if err := c.BodyParser(&registerData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid registerData"})
	}

	lowerCaseEmail := strings.ToLower(registerData.Email)
	// Check if email already exists
	var existingUser models.User
	if err := database.DB.Where("email = ?", registerData.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email already in use"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), 12)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not hash password"})
	}

	user := models.User{
		FirstName:      registerData.FirstName,
		LastName:       registerData.LastName,
		Email:          lowerCaseEmail,
		HashedPassword: string(hashedPassword),
		IsGmailLinked:  false,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

func Login(c *fiber.Ctx) error {
	var loginValues LoginSchema
	if err := c.BodyParser(&loginValues); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid values",
		})
	}

	// Normalize email
	email := strings.ToLower(loginValues.Email)

	// Find the user
	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(loginValues.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	// Create JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}
	//return token if the user is logged in successfully
	return c.JSON(fiber.Map{
		"message": "Logged in successful",
		"token":   token,
	})
}
