package handlers

import (
	// "task/models"
	"github.com/gofiber/fiber/v2"
 
	"task/utils"
	"task/config"
	"task/models"
)

func Dashboard(c *fiber.Ctx) error {
	// Retrieve user ID from the JWT claim stored in locals
	userID := c.Locals("userID")

	// Do something with the userID, like fetching user-specific data
	// For example, if you want to get the user from the DB based on their ID:
	var user models.User
	if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "User not found", err.Error())
	}

	// preload tasks
	if err := config.DB.Preload("Tasks").Where("id = ?", userID).First(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to retrieve user with tasks", err.Error())
	}

	// Respond with user data
	return utils.SuccessResponse(c, fiber.StatusOK, "Dashboard data", fiber.Map{
		"user": user,
		// "tasks" : user.Tasks,
	})
}