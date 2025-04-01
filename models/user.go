
package models

import (
	"gorm.io/gorm"
)


// User model for the users table
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Firstname string `gorm:"not null"`
	Lastname string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Tasks    []Task `gorm:"foreignKey:UserID"` // One-to-many relationship with Task
}
