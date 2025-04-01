package models

import "gorm.io/gorm"

// Task model for the tasks table
type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"type TEXT"`
	UserID      uint // Foreign key linking to User model
	User        User  // Association with User
	IsCompleted bool
	Image       string
}
