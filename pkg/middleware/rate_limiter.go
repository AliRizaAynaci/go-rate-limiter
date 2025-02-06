package middleware

import (
	"rate-limiter/internal/logger"
	"rate-limiter/internal/redis"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	windowSize  = 60 * time.Second
	maxRequests = 5
)

func SlidingWindowRateLimiter(c *fiber.Ctx) error {
	identifier := c.Locals("apiKey")
	if identifier == nil {
		identifier = c.IP()
	}

	key := "rate_limit" + identifier.(string)
	now := time.Now().Unix()

	// Clean the old requests from the redis
	err := redis.RemoveOldRequests(key, now-int64(windowSize.Seconds()))
	if err != nil {
		logger.Error("Redis eski istekleri temizleme hatası: "+err.Error(), c.Path())
		return c.Status(500).SendString("Redis hatası!")
	}

	// Get the new requests count
	count, err := redis.GetRequestsCount(key)
	if err != nil {
		logger.Error("Redis istek sayısı alma hatası: "+err.Error(), c.Path())
		return c.Status(500).SendString("Redis hatası!")
	}

	if count >= maxRequests {
		logger.Warn("Rate limit aşıldı: "+identifier.(string), c.Path())
		return c.Status(429).SendString("Çok fazla istek yaptınız!")
	}

	// Add new request into the window
	err = redis.AddRequest(key, now)
	if err != nil {
		logger.Error("Redis yeni istek ekleme hatası: "+err.Error(), c.Path())
		return c.Status(500).SendString("Redis hatası!")
	}

	logger.Info("Rate limit başarılı: "+identifier.(string), c.Path())
	return c.Next()
}
