package handlers

import (
	"fmt"
	"math"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"task/config"
	"task/models"
	"task/utils"
)

func GetTasks(c *fiber.Ctx) error {
	db := config.DB
	userIDStr := c.Locals("userID").(string)
	userID, _ := strconv.ParseUint(userIDStr, 10, 64)

	// Pagination
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset := (page - 1) * limit

	// Filters
	isCompleted := c.Query("is_completed")
	dueDate := c.Query("due_date")

	var tasks []models.Task
	query := db.Where("user_id = ?", int(userID))

	if isCompleted != "" {
		completed, _ := strconv.ParseBool(isCompleted)
		query = query.Where("is_completed = ?", completed)
	}

	if dueDate != "" {
		query = query.Where("DATE(due_date) = ?", dueDate)
	}

	// Get total count
	var total int64
	query.Model(&models.Task{}).Count(&total)

	// Fetch paginated data
	err := query.
		Limit(limit).
		Offset(offset).
		Order("created_at desc").
		Find(&tasks).Error

	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch tasks", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusOK, "Fetched Tasks", fiber.Map{
		"data":       tasks,
		"total":      total,
		"page":       page,
		"per_page":   limit,
		"total_page": int(math.Ceil(float64(total) / float64(limit))),
	})
}

// GetTaskByID fetches a single task by its ID
func GetTaskByID(c *fiber.Ctx) error {
	// Extract the task ID from the URL
	taskIDStr := c.Params("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid task ID", err)
	}

	// Fetch the task from the database
	var task models.Task
	err = config.DB.First(&task, taskID).Error
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Task not found", err)
	}

	// Return the task as a JSON response
	return utils.SuccessResponse(c, fiber.StatusOK, "Fetched Task", task)
}

func CreateTask(c *fiber.Ctx) error {
	// Struct for task input validation
	type CreateTaskInput struct {
		Title       string    `json:"title" validate:"required"`
		Description string    `json:"description" validate:"required"`
		DueDate     string `json:"due_date" validate:"required"`
	}

	var input CreateTaskInput

	// Parse request body
	if err := c.BodyParser(&input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Failed to parse request body", err.Error())
	}

	// Validate input fields
	if err := utils.ValidateStruct(input); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input", err)
	}

	// Get userID from context
	userIDStr := c.Locals("userID").(string)
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	parsedDate, err := time.Parse("2006-01-02", input.DueDate)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid due date format. Use YYYY-MM-DD", err.Error())
	}


	// Handle file upload
	file, err := c.FormFile("image")
	if err != nil || file.Size == 0 {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Image file is required", "No file was uploaded or file is empty")
	}
	// Validate file type (only image/jpeg, image/png)
	if err := validateImageFile(file); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid file type", err.Error())
	}
	if file.Size > 2*1024*1024 {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "File too large", "Maximum allowed size is 2MB")
	}

	// Create folder if it does not exist
	imageFolder := "./uploads/images"
	if err := os.MkdirAll(imageFolder, os.ModePerm); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create image folder", err.Error())
	}

	// Generate a unique filename
	timestamp := time.Now().UnixNano()
	ext := filepath.Ext(file.Filename)
	ext = strings.ToLower(ext)
	fileName := fmt.Sprintf("%d%s", timestamp, ext)

	// Save the file to the specified folder
	err = c.SaveFile(file, fmt.Sprintf("%s/%s", imageFolder, fileName))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to save file", err.Error())
	}

	// Create the task in the database
	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     &parsedDate, // DueDate is passed as a pointer to time.Time
		Image:       fileName,       // Save the filename in the database
		UserID:      uint(userID),   // Convert userID to uint
	}

	// Insert the task into the database
	if err := config.DB.Create(&task).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create task", err.Error())
	}

	return utils.SuccessResponse(c, fiber.StatusCreated, "Task created successfully", task)
}

// validateImageFile checks if the uploaded file is an image (jpeg/png)
func validateImageFile(file *multipart.FileHeader) error {
	allowedTypes := []string{"image/jpeg", "image/png"}
	fileType := file.Header.Get("Content-Type")

	// Check if the file's MIME type is in the allowed types
	for _, validType := range allowedTypes {
		if fileType == validType {
			return nil
		}
	}
	return fmt.Errorf("invalid file type: %s", fileType)
}
