package handlers

import (
	"rate-limiter/internal/database"
	"rate-limiter/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetLogsHandler(c *fiber.Ctx) error {
	db := database.GetDb()
	var logs []models.LogEntry

	query := db

	// Filtreleme seçenekleri
	if level := c.Query("level"); level != "" {
		query = query.Where("level = ?", level)
	}
	if endpoint := c.Query("endpoint"); endpoint != "" {
		query = query.Where("endpoint = ?", endpoint)
	}
	if search := c.Query("search"); search != "" {
		query = query.Where("message LIKE ?", "%"+search+"%")
	}

	// Son 100 log kaydını getir
	query.Order("timestamp desc").Limit(100).Find(&logs)

	return c.JSON(logs)
}
