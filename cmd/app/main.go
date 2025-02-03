package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"rate-limiter/internal/redis"
	"rate-limiter/pkg/middleware"
)

func main() {
	redisConfig := redis.RedisConfig{
		Addr: "redis:6379",
	}

	redis.NewRedisClient(redisConfig)

	app := fiber.New()

	app.Use(middleware.RateLimiter)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Rate Limiter Çalışıyor!")
	})

	fmt.Println("Server çalışıyor: http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
