package models

import "time"

// LogEntry represents a log entry in the system
type LogEntry struct {
	ID        uint      `json:"id" gorm:"primaryKey" example:"1"`
	Level     string    `json:"level" gorm:"index" example:"INFO"`
	Timestamp time.Time `json:"timestamp" example:"2024-03-20T15:04:05Z"`
	Message   string    `json:"message" example:"Request processed successfully"`
	Endpoint  string    `json:"endpoint" gorm:"index" example:"/api/protected-endpoint"`
}
