package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"clubApi/models"
)

var database *gorm.DB

// DatabaseInit initializes the database connection and runs migrations.
func DatabaseInit() error {
	// Database connection credentials
	credentials := "root:admin@tcp(localhost:3306)/tti"

	// Connect to the database
	var err error
	database, err = gorm.Open(mysql.Open(credentials), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	fmt.Println("Connected to the database successfully")

	// AutoMigrate the User model
	if err := database.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("failed to migrate table: %w", err)
	}
	// AutoMigrate the User model
	if err := database.AutoMigrate(&models.Club{}); err != nil {
		return fmt.Errorf("failed to migrate table: %w", err)
	}

	return nil
}

// DB returns the database connection instance.
func DB() *gorm.DB {
	if database == nil {
		log.Fatal("Database connection is not initialized")
	}
	return database
}
