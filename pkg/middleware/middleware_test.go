package middleware_test

import (
	"github.com/gofiber/fiber/v2"
	"net/http/httptest"
	"rate-limiter/internal/database"
	"rate-limiter/pkg/middleware"
	"testing"
)

func TestMain(m *testing.M) {
	database.ConnectDb()
	m.Run()
}

func TestAPIKeyMiddleware(t *testing.T) {
	app := fiber.New()
	app.Use(middleware.APIKeyMiddleware)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("X-API-KEY", "gecersiz")
	resp, _ := app.Test(req)

	if resp.StatusCode != 401 {
		t.Fatalf("Geçersiz API Key için 401 bekleniyordu, ancak %d geldi", resp.StatusCode)
	}
}
