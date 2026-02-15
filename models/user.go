
package models

import (
	"gorm.io/gorm"
	"task/utils"
)


// User model for the users table
type User struct {
	Username string `gorm:"unique;not null"`
	Firstname string `gorm:"not null"`
	Lastname string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	IsVerified bool `gorm:"default:false" json:"is_verified"`
	Tasks    []Task `gorm:"foreignKey:UserID"` // One-to-many relationship with Task
	gorm.Model
}


// BeforeCreate GORM Hook: Automatically hash password before saving
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := utils.HashPassword(u.Password) // Use auth package instead of utils
	if err != nil {
		return err // Simply return the error (GORM expects an error, not an HTTP response)
	}

	u.Password = hashedPassword
	return nil
}


