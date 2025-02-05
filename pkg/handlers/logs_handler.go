package handlers

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/internal/database"
	"rate-limiter/internal/models"
)

func GetLogsHandler(c *fiber.Ctx) error {
	db := database.GetDb()
	var logs []models.LogEntry

	level := c.Query("level")
	query := db
	if level != "" {
		query = query.Where("level = ?", level)
	}

	query.Order("timestamp desc").Find(&logs)

	return c.JSON(logs)
}
