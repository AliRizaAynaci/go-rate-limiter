package middleware

import (
	"fmt"
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

	fmt.Printf("🟡 Gelen IP: %s\n", c.IP()) // Debug için ekleyelim

	if identifier == "0.0.0.0" {
		fmt.Println("⚠️  IP 0.0.0.0 olarak algılandı, düzeltme yapılıyor...")
		identifier = "127.0.0.1" // Test sırasında düzgün çalışması için
	}

	key := "rate_limit:" + identifier.(string)
	now := time.Now().Unix()

	fmt.Printf("🟡 Rate limiter çalışıyor. Key: %s\n", key)

	err := redis.RemoveOldRequests(key, now-int64(windowSize.Seconds()))
	if err != nil {
		fmt.Println("❌ Redis eski istekleri temizleyemedi!")
		return c.Status(500).SendString("Redis hatası!")
	}

	count, err := redis.GetRequestsCount(key)
	if err != nil {
		fmt.Println("❌ Redis'ten istek sayısı alınamadı!")
		return c.Status(500).SendString("Redis hatası!")
	}

	fmt.Printf("🟢 Redis'teki mevcut istek sayısı: %d\n", count)

	if count >= maxRequests {
		fmt.Println("🚫 Rate limit aşıldı, 429 dönülüyor!")
		return c.Status(429).SendString("Çok fazla istek yaptınız!")
	}

	err = redis.AddRequest(key, now)
	if err != nil {
		fmt.Println("❌ Redis'e yeni istek eklenemedi!")
		return c.Status(500).SendString("Redis hatası!")
	}

	fmt.Println("✅ Redis'e yeni istek başarıyla eklendi!")

	return c.Next()
}
