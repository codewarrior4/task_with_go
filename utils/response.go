package utils

import "github.com/gofiber/fiber/v2"

// ErrorResponse formats error messages
func ErrorResponse(c *fiber.Ctx, status int, message string, details interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"errors":  details,
		"status":  "error",
		"message": message,
	})
}

// SuccessResponse formats success responses
func SuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}
