package database

import (
	"fmt"

	"github.com/KainNhantumbo/books-api/config"
	"github.com/KainNhantumbo/books-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	DATABASE_URL := config.GetEnvValue("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database.")
	}

	fmt.Println("Database connected!")

	// Migrate the database
	DB.AutoMigrate(&model.Book{})
	fmt.Println("Database Migrated.")
}
