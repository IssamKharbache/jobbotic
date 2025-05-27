package handlers

import (
	"jobbotic-backend/database"
	"jobbotic-backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	//fetching the users and handling errors
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch users",
		})
	}
	//return users
	return c.JSON(users)
}
