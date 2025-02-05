package logger

import (
	"encoding/json"
	"log"
	"rate-limiter/internal/database"
	"rate-limiter/internal/models"
	"time"
)

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

func LogMessage(level, message, endpoint string) {
	db := database.GetDb()

	if db == nil {
		log.Println("Error: Database connection is nil. Skipping log entry.")
		return
	}

	entry := models.LogEntry{
		Level:     level,
		Timestamp: time.Now(),
		Message:   message,
		Endpoint:  endpoint,
	}

	logData, err := json.Marshal(entry)
	if err != nil {
		log.Fatalf("JSON log oluşturulamadı: %v", err)
	}
	log.Println(string(logData))

	if err = db.Create(&entry).Error; err != nil {
		log.Println("Error saving log entry to database:", err)
	}
}

func Debug(message, endpoint string) {
	LogMessage(DEBUG, message, endpoint)
}

func Info(message, endpoint string) {
	LogMessage(INFO, message, endpoint)
}

func Warn(message, endpoint string) {
	LogMessage(WARN, message, endpoint)
}

func Error(message, endpoint string) {
	LogMessage(ERROR, message, endpoint)
}
