package handlers

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Register endpoint"})
}

func Login(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Login endpoint"})
}

func ForgotPassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Forgot Password endpoint"})
}

func ResetPassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Reset Password endpoint"})
}

func ChangePassword(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Change Password endpoint"})
}
