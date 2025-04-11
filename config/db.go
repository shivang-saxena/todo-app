package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// SetupDB initializes the database connection
func SetupDB() *gorm.DB {
	dbURL := GetEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/todo_app?sslmode=disable")
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.LogMode(true)
	fmt.Println("Connected to database")
	
	// Auto migrate models
	// db.AutoMigrate(&models.Todo{})
	
	return db
}
