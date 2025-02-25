package models

import (
	"time"

	"gorm.io/gorm"
)

// APIKey represents an API key for authentication and rate limiting
type APIKey struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Key       string    `json:"key" gorm:"uniqueIndex" example:"test-api-key-123"`
	Limit     int       `json:"limit" example:"100"`
	CreatedAt time.Time `json:"created_at" example:"2024-03-20T15:04:05Z"`
}

func (key *APIKey) BeforeCreate(tx *gorm.DB) (err error) {
	key.CreatedAt = time.Now()
	return
}
