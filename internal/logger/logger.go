package logger

import (
	"log"
	"os"
	"path/filepath"
)

const (
	INFO  = "INFO"
	ERROR = "ERROR"
)

func LogMessage(level, message string) {
	logDir := "/logs"
	if _, err := os.Stat("/logs"); os.IsNotExist(err) {
		logDir = "logs"
	}

	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		log.Fatalf("Log dizini oluşturulamadı: %v", err)
	}

	logPath := filepath.Join(logDir, "app.log")
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)
	logger.SetPrefix("[" + level + "] ")
	logger.Println(message)
}

func Info(message string) {
	LogMessage(INFO, message)
}

func Error(message string) {
	LogMessage(ERROR, message)
}
