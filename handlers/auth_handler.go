package handlers

import (
	"task/auth"
	"task/config"
	"task/models"
	"task/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Username  string `json:"username" validate:"required"`
		Firstname string `json:"firstname" validate:"required"`
		Lastname  string `json:"lastname" validate:"required"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required"`
	}

	var input RegisterInput

	// Parse request body
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body", err.Error())
	}

	// Validate input
	if err := utils.ValidateStruct(input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input", err) // Directly return `err`
	}

	// Check if user already exists
	var existingUser models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "User already exists", nil)
	}

	user := models.User{
		Username:  input.Username,
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Email:     input.Email,
		Password:  input.Password, // Password hashing is handled in User model
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
	}

	// generate verification code
	verificationCode := utils.GenerateVerificationCode()
	verification :=models.Verification{
		UserID:    user.ID,
		Code:      verificationCode,
		ExpiresAt: time.Now().Add(60 * time.Minute),
	}

	if err := config.DB.Create(&verification).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create verification code", err.Error())
	}

	// Send verification email
	go func () {
		subject := "Email Verification"
		message := []string{
			"Hello " + user.Firstname + ",",
			"Please use the following code to verify your email: " + verificationCode,
			"If you did not request this verification, please ignore this email.",
			"Best Regards,\nThe Team",
		}
		utils.SendEmail(user.Email, subject, message)
	}()
	

	// Generate JWT token
	token, err := auth.GenerateJWT(user.ID) // Pass user.ID (which is uint)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create token", err.Error())
	}

	// Send welcome email asynchronously
	go func() {
		subject := "Welcome to Our Platform!"
		message := []string{
			"Hello " + user.Firstname + ",",
			"Thank you for registering with us!",
			"We're excited to have you on board.",
			"If you have any questions, feel free to reach out.",
			"Best Regards,\nThe Team",
		}
		utils.SendEmail(user.Email, subject, message)
	}()

	

	return utils.SuccessResponse(c, fiber.StatusCreated, "User registered successfully", fiber.Map{
		"user": fiber.Map{
			"id":        user.ID,
			"username":  user.Username,
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
			"email":     user.Email,
		},
		"token": token,
	})

}

func Login(c *fiber.Ctx) error {

	type LoginInput struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	var input LoginInput

	// parse request body
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body", err.Error())
	}

	// validate Input
	if err := utils.ValidateStruct(input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input", err) // Directly return `err`
	}

	// check if user exists
	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid credentials", nil)
	}

	// check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Password is incrrect", nil)	
	}

	// check if user is email verified
	if !user.IsVerified {
		
		// generate verification code
		verificationCode := utils.GenerateVerificationCode()
		verification :=models.Verification{
			UserID:    user.ID,
			Code:      verificationCode,
			ExpiresAt: time.Now().Add(60 * time.Minute),
		}

		if err := config.DB.Create(&verification).Error; err != nil {
			return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create verification code", err.Error())
		}

		// Send verification email
		go func () {
			subject := "Email Verification"
			message := []string{
				"Hello " + user.Firstname + ",",
				"Please use the following code to verify your email: " + verificationCode,
				"If you did not request this verification, please ignore this email.",
				"Best Regards,\nThe Team",
			}
			utils.SendEmail(user.Email, subject, message)
		}()

		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Email not verified", nil)
	}

	// generate jwt token
	token, err := auth.GenerateJWT(user.ID)

	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create token", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Login successful", fiber.Map{
		"user": fiber.Map{
			"id":        user.ID,
			"username":  user.Username,
			"firstname": user.Firstname,
			"lastname":  user.Lastname,
			"email":     user.Email,
		},
		"token": token,
		"tasks": user.Tasks,
	})

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
