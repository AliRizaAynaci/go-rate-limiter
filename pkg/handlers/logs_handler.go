package handlers

import (
	"rate-limiter/internal/database"
	"rate-limiter/internal/models"

	"github.com/gofiber/fiber/v2"
)

// GetLogsHandler handles requests to retrieve log entries with optional filtering
// Supports filtering by level, endpoint, and message content
// Returns the last 100 log entries ordered by timestamp descending
func GetLogsHandler(c *fiber.Ctx) error {
	db := database.GetDb()
	var logs []models.LogEntry

	query := db

	if level := c.Query("level"); level != "" {
		query = query.Where("level = ?", level)
	}
	if endpoint := c.Query("endpoint"); endpoint != "" {
		query = query.Where("endpoint = ?", endpoint)
	}
	if search := c.Query("search"); search != "" {
		query = query.Where("message LIKE ?", "%"+search+"%")
	}

	query.Order("timestamp desc").Limit(100).Find(&logs)

	return c.JSON(logs)
}
