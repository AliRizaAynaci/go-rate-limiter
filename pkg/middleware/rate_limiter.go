package middleware

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/internal/logger"
	"rate-limiter/internal/redis"
	"time"
)

func RateLimiter(c *fiber.Ctx) error {
	ip := c.IP()
	endpoint := c.Path()

	count, err := redis.IncrementInRedis(ip)
	if err != nil {
		logger.Error("Redis hata: "+err.Error(), endpoint)
		return c.Status(500).SendString("Redis hata!")
	}

	if count == 1 {
		err = redis.SetExpire(ip, 1*time.Minute)
		if err != nil {
			logger.Error("Redis expire ayarlama hatası: "+err.Error(), endpoint)
			return c.Status(500).SendString("Redis hata!")
		}
	}

	if count > 5 {
		logger.Warn("Rate limit aşılmaya çalışıldı: "+ip, endpoint)
		return c.Status(429).SendString("Çok fazla istek yaptınız!")
	}

	logger.Info("Rate limit kontrolü başarıyla geçti: "+ip, endpoint)
	return c.Next()
}
