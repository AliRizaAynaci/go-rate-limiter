package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"rate-limiter/internal/database"
	"rate-limiter/internal/prometheus"
	"rate-limiter/internal/redis"
	"rate-limiter/pkg/handlers"
	"rate-limiter/pkg/middleware"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// prometheusHandler returns a Fiber handler that serves Prometheus metrics
func prometheusHandler() fiber.Handler {
	return adaptor.HTTPHandler(promhttp.Handler())
}

// setupRoutes configures all the routes for the Fiber application
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

// gracefulShutdown handles the graceful shutdown process of the application
// It closes Redis and Database connections, and shuts down the server properly
func gracefulShutdown(app *fiber.App) {
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM)

	<-shutdownChan
	fmt.Println("\nGraceful shutdown başlatılıyor...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := redis.GetClient().Close(); err != nil {
		log.Printf("Redis bağlantısı kapatılırken hata: %v", err)
	}

	sqlDB, err := database.GetDb().DB()
	if err != nil {
		log.Printf("Database SQL bağlantısı alınırken hata: %v", err)
	} else {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Database bağlantısı kapatılırken hata: %v", err)
		}
	}

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server kapatılırken hata: %v", err)
	}

	fmt.Println("Server güvenli bir şekilde kapatıldı")
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

	go func() {
		fmt.Println("Server çalışıyor: http://localhost:3000")
		if err := app.Listen(":3000"); err != nil {
			log.Printf("Server hatası: %v", err)
		}
	}()

	gracefulShutdown(app)
}
