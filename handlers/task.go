package handlers

import (
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"task/config"
	"task/models"
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch tasks",
		})
	}

	return c.JSON(fiber.Map{
		"data":       tasks,
		"total":      total,
		"page":       page,
		"per_page":   limit,
		"total_page": int(math.Ceil(float64(total) / float64(limit))),
	})
}
