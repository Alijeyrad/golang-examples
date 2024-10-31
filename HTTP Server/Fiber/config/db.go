package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"library/models"
)

var db *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Migrate the schema
	db.AutoMigrate(&models.Book{})
}

func GetDB() *gorm.DB {
	return db
}
