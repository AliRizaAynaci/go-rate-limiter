package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"rate-limiter/internal/database"
	"rate-limiter/internal/redis"
	"rate-limiter/pkg/handlers"
	"rate-limiter/pkg/middleware"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Rate Limiter Çalışıyor!")
	})

	app.Get("/logs", handlers.GetLogsHandler)

	api := app.Group("/api", middleware.RateLimiter)
	api.Get("/protected-endpoint", func(c *fiber.Ctx) error {
		return c.SendString("Bu endpoint rate limiter tarafından korunuyor!")
	})
}

func main() {

	database.ConnectDb()

	redisConfig := redis.RedisConfig{
		Addr: "redis:6379",
	}

	redis.NewRedisClient(redisConfig)

	app := fiber.New()

	setupRoutes(app)

	fmt.Println("Server çalışıyor: http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
