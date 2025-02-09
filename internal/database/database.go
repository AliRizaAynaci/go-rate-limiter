package database

import (
	"log"
	"os"
	"rate-limiter/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB

func ConnectDb() {
	var err error
	database, err = gorm.Open(sqlite.Open("logs.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database! \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the database successfully")
	database.Logger = logger.Default.LogMode(logger.Error)

	log.Println("Running Migrations")
	err = database.AutoMigrate(&models.LogEntry{}, &models.APIKey{})
	if err != nil {
		log.Fatal("Failed to migrate database! \n", err.Error())
		os.Exit(2)
	}
}

func GetDb() *gorm.DB {
	if database == nil {
		log.Fatal("Database connection is nil! Application is exiting.")
		os.Exit(1)
	}
	return database
}

func SetDb(db *gorm.DB) {
	database = db
}
