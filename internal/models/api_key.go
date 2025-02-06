package models

import (
	"time"

	"gorm.io/gorm"
)

type APIKey struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Key       string    `json:"key" gorm:"uniqueIndex"`
	RateLimit int       `json:"rate_limit"`
	CreatedAt time.Time `json:"created_at"`
}

func (key *APIKey) BeforeCreate(tx *gorm.DB) (err error) {
	key.CreatedAt = time.Now()
	return
}
