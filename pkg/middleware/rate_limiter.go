package middleware

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/internal/logger"
	"rate-limiter/internal/redis"
	"time"
)

func RateLimiter(c *fiber.Ctx) error {
	identifier := c.Locals("apiKey")
	if identifier == nil {
		identifier = c.IP()
	}

	count, err := redis.IncrementInRedis(identifier.(string))
	if err != nil {
		logger.Error("Redis hata: "+err.Error(), c.Path())
		return c.Status(500).SendString("Redis hata!")
	}

	rateLimit := c.Locals("rateLimit")
	limit := 5
	if rateLimit != nil {
		limit = rateLimit.(int)
	}

	if count == 1 {
		err = redis.SetExpire(identifier.(string), 1*time.Minute)
		if err != nil {
			logger.Error("Redis expire ayarlama hatası: "+err.Error(), c.Path())
			return c.Status(500).SendString("Redis hata!")
		}
	}

	if count > int64(limit) {
		logger.Warn("Rate limit aşılmaya çalışıldı: "+identifier.(string), c.Path())
		return c.Status(429).SendString("Çok fazla istek yaptınız!")
	}

	logger.Info("Rate limit kontrolü başarıyla geçti: "+identifier.(string), c.Path())
	return c.Next()
}
