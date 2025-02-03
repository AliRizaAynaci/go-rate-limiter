// internal/middleware/middleware_test.go
package middleware_test

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	"rate-limiter/internal/redis"
	"rate-limiter/pkg/middleware"
	"testing"
)

func setupTestApp(redisDB int) *fiber.App {
	config := redis.RedisConfig{
		Addr: "localhost:6379",
		DB:   redisDB,
	}
	client := redis.NewRedisClient(config)
	client.FlushDB(context.Background())

	app := fiber.New()
	app.Use(middleware.RateLimiter)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Merhaba, dünya!")
	})
	return app
}

func TestRateLimiter_AllowsUnderLimit(t *testing.T) {
	app := setupTestApp(2)

	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "1.2.3.4:12345"

	for i := 0; i < 5; i++ {
		resp, err := app.Test(req)
		if err != nil {
			t.Fatalf("Test isteği sırasında hata: %v", err)
		}
		if resp.StatusCode != 200 {
			t.Fatalf("Beklenen 200 status kodu, fakat %d geldi", resp.StatusCode)
		}
	}
}

func TestRateLimiter_BlocksOverLimit(t *testing.T) {
	app := setupTestApp(2)

	req := httptest.NewRequest("GET", "/", nil)
	req.RemoteAddr = "5.6.7.8:56789"

	var resp *httptest.ResponseRecorder
	for i := 0; i < 6; i++ {
		response, err := app.Test(req)
		if err != nil {
			t.Fatalf("Test isteği sırasında hata: %v", err)
		}
		resp = httptest.NewRecorder()
		resp.Code = response.StatusCode
	}

	if resp.Code != 429 {
		t.Fatalf("Limit aşıldığında beklenen status kodu 429, fakat %d geldi", resp.Code)
	}
}
