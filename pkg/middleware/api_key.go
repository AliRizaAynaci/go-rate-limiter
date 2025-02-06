package middleware

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/internal/database"
	"rate-limiter/internal/models"
)

func APIKeyMiddleware(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-KEY")

	if apiKey == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "API Key is required",
		})
	}

	db := database.GetDb()
	var key models.APIKey
	if err := db.Where("key = ?", apiKey).First(&key).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid API Key",
		})
	}

	c.Locals("apiKey", key.Key)
	c.Locals("rateLimit", key.Limit)

	return c.Next()
}
