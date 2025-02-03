package middleware

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/internal/logger"
	"rate-limiter/internal/redis"
	"time"
)

func RateLimiter(c *fiber.Ctx) error {
	ip := c.IP()

	count, err := redis.IncrementInRedis(ip)
	if err != nil {
		logger.Error("Redis hata: " + err.Error())
		return c.Status(500).SendString("Redis hata!")
	}

	if count == 1 {
		err = redis.SetExpire(ip, 1*time.Minute)
		if err != nil {
			logger.Error("Redis expire ayarlama hatası: " + err.Error())
			return c.Status(500).SendString("Redis hata!")
		}
	}

	if count > 5 {
		logger.Info("Rate limit aşılmaya çalışıldı: " + ip)
		return c.Status(429).SendString("Çok fazla istek yaptınız!")
	}

	logger.Info("Rate limit kontrolü başarıyla geçti: " + ip)

	return c.Next()
}
