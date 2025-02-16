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
	dbPath := "/data/logs.db" // Container içinde sabit bir yol
	if os.Getenv("DB_PATH") != "" {
		dbPath = os.Getenv("DB_PATH")
	}

	database, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
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

	// Default API Key oluştur
	var count int64
	database.Model(&models.APIKey{}).Count(&count)
	if count == 0 {
		defaultKey := models.APIKey{
			Key:   "test-api-key-123",
			Limit: 100,
		}
		if err := database.Create(&defaultKey).Error; err != nil {
			log.Printf("Default API Key oluşturulamadı: %v\n", err)
		} else {
			log.Printf("Default API Key oluşturuldu: %s\n", defaultKey.Key)
		}
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
