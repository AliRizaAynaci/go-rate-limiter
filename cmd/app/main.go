package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"log"
	"rate-limiter/internal/database"
	"rate-limiter/internal/prometheus"
	"rate-limiter/internal/redis"
	"rate-limiter/pkg/handlers"
	"rate-limiter/pkg/middleware"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func prometheusHandler() fiber.Handler {
	return adaptor.HTTPHandler(promhttp.Handler())
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		prometheus.RecordRedisRequest("/")
		return c.SendString("API Rate Limiter Çalışıyor!")
	})

	app.Get("/metrics", prometheusHandler())

	app.Get("/logs", handlers.GetLogsHandler)

	api := app.Group("/api", middleware.APIKeyMiddleware, middleware.SlidingWindowRateLimiter)

	api.Get("/protected-endpoint", func(c *fiber.Ctx) error {
		prometheus.RecordRedisRequest("/api/protected-endpoint")
		return c.SendString("Bu endpoint API Key ile korunuyor!")
	})

	api.Use(func(c *fiber.Ctx) error {
		if c.Response().StatusCode() == fiber.StatusTooManyRequests {
			prometheus.RecordRateLimitViolation()
		}
		return c.Next()
	})
}

func main() {
	database.ConnectDb()

	redisConfig := redis.RedisConfig{
		Addr: "redis:6379",
	}
	redis.NewRedisClient(redisConfig)

	prometheus.InitMetrics()

	app := fiber.New()
	setupRoutes(app)

	fmt.Println("Server çalışıyor: http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
