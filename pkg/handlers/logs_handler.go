package handlers

import (
	"rate-limiter/internal/database"
	"rate-limiter/internal/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get logs
// @Description Get the last 100 log entries with optional filtering
// @Tags logs
// @Accept json
// @Produce json
// @Param level query string false "Filter by log level (DEBUG, INFO, WARN, ERROR)"
// @Param endpoint query string false "Filter by endpoint"
// @Param search query string false "Search in message content"
// @Success 200 {array} models.LogEntry
// @Router /logs [get]
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
