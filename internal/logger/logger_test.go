package logger_test

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"rate-limiter/internal/database"
	"rate-limiter/internal/logger"
	"rate-limiter/internal/models"
	"testing"
)

func setupTestDB() {
	testDb, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	testDb.AutoMigrate(&models.LogEntry{}, &models.APIKey{})
	database.SetDb(testDb)
}

func TestMain(m *testing.M) {
	setupTestDB()
	m.Run()
}

func TestLogMessage(t *testing.T) {
	logger.Info("Info Log Test", "/test")
	logger.Debug("Debug Log Test", "/test")
	logger.Warn("Warning Log Test", "/test")
	logger.Error("Error Log Test", "/test")

	var count int64
	db := database.GetDb()
	db.Model(&models.LogEntry{}).Count(&count)

	if count != 4 {
		t.Fatalf("Beklenen 4 log entry, ancak %d bulundu", count)
	}
}
