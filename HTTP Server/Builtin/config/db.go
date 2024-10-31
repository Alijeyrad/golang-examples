package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"library/models"
)

var (
	once sync.Once
)

func NewDB() (*gorm.DB, error) {
	var db *gorm.DB
	once.Do(func() {
		var err error
		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		pass := os.Getenv("DB_PASS")
		name := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")
		ssl := os.Getenv("DB_SSLMODE")
		dsn := fmt.Sprintf(
			"host=%v user=%v password=%v dbname=%v port=%v sslmode=%v",
			host,
			user,
			pass,
			name,
			port,
			ssl,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}
		// Migrate the schema
		db.AutoMigrate(&models.Book{})
	})
	return db, nil
}
