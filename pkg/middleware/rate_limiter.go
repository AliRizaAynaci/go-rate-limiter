package middleware

import (
	"fmt"
	"rate-limiter/internal/database"
	"rate-limiter/internal/models"
	"rate-limiter/internal/redis"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	windowSize  = 60 * time.Second
	maxRequests = 5
)

// SlidingWindowRateLimiter implements a sliding window rate limiting algorithm
// It limits the number of requests that can be made within a specified time window
// Uses Redis to track request counts and enforces rate limits per API key or IP address
func SlidingWindowRateLimiter(c *fiber.Ctx) error {
	identifier := c.Locals("apiKey")
	if identifier == nil {
		identifier = c.IP()
	}

	if identifier == "0.0.0.0" {
		identifier = "127.0.0.1"
	}

	key := "rate_limit:" + identifier.(string)
	now := time.Now().Unix()

	err := redis.RemoveOldRequests(key, now-int64(windowSize.Seconds()))
	if err != nil {
		return c.Status(500).SendString("Redis hatası!")
	}

	count, err := redis.GetRequestsCount(key)
	if err != nil {
		return c.Status(500).SendString("Redis hatası!")
	}

	if count >= maxRequests {
		logEntry := models.LogEntry{
			Level:     "WARNING",
			Timestamp: time.Now(),
			Message:   fmt.Sprintf("Rate limit exceeded for API key: %s", identifier.(string)),
			Endpoint:  c.Path(),
		}

		db := database.GetDb()
		if err := db.Create(&logEntry).Error; err != nil {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Rate limit exceeded",
			})
		}

		return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
			"error": "Rate limit exceeded",
		})
	}

	logEntry := models.LogEntry{
		Level:     "INFO",
		Timestamp: time.Now(),
		Message:   fmt.Sprintf("Request processed for API key: %s (Count: %d/%d)", identifier.(string), count, maxRequests),
		Endpoint:  c.Path(),
	}

	db := database.GetDb()
	if err := db.Create(&logEntry).Error; err != nil {
		fmt.Printf("Log kaydı oluşturulurken hata: %v\n", err)
	}

	err = redis.AddRequest(key, now)
	if err != nil {
		return c.Status(500).SendString("Redis hatası!")
	}

	return c.Next()
}
