package utils

import "github.com/gofiber/fiber/v2"

// ErrorResponse formats error messages
func ErrorResponse(c *fiber.Ctx, status int, message string, details interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"success":  false,
		"errors":  details,
		"message": message,
	})
}

// SuccessResponse formats success responses
func SuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"success":  true,
		"message": message,
		"data":    data,
	})
}
