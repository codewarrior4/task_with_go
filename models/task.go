package models

import (
	"time"

	"gorm.io/gorm"
)

// Task model for the tasks table
type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"type TEXT"`
	UserID      uint // Foreign key linking to User model
	User        User  // Association with User
	IsCompleted bool
	DueDate     *time.Time     // Optional due date
	Image       string
	DeletedAt   gorm.DeletedAt `gorm:"index"` // Enables soft deletes
}
