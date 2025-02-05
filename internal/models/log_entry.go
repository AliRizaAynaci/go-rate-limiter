package models

import "time"

type LogEntry struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Level     string    `json:"level" gorm:"index"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Endpoint  string    `json:"endpoint" gorm:"index"`
}
