package models

import (
	"time"

	"gorm.io/gorm"
)

type Verification struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"uniqueIndex"`  // Link to user
	Code      string    `json:"code" gorm:"size:6"`
	ExpiresAt time.Time `json:"expires_at"` // Expiry time
	CreatedAt time.Time `json:"created_at"`
}

func (v *Verification) BeforeCreate(tx *gorm.DB) error {
	v.ExpiresAt = time.Now().Add(60 * time.Minute)  // Set expiry time
	return nil
}