package logger_test

import (
	"testing"

	"rate-limiter/internal/logger"
)

func TestInfoLogging(t *testing.T) {
	logger.Info("Test Info Mesajı")
}

func TestErrorLogging(t *testing.T) {
	logger.Error("Test Error Mesajı")
}
