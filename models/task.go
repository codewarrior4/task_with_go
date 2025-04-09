package models

import (
	"encoding/json"
	"fmt"
	"task/config"
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


func (t *Task) MarshalJSON() ([]byte, error) {
	type Alias Task // avoid recursion

	baseURL := config.GetEnv("APP_URL", "http://localhost:4023")

	return json.Marshal(&struct {
		*Alias
		ImageURL string `json:"image_url"`
	}{
		Alias:    (*Alias)(t),
		ImageURL: fmt.Sprintf("%s/uploads/images/%s", baseURL, t.Image),
	})
}
