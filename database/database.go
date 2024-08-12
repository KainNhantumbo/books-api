package database

import (
	"fmt"

	"github.com/KainNhantumbo/books-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Declare the variable for the database
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	connection := config.Config("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(connection))

	if err != nil {
		panic("Failed to connect database.")
	}

	fmt.Println("Database connected!")
}
