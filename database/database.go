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
	connection := config.GetEnvValue("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database.")
	}

	// sqlDB, err := DB.DB()

	// if err != nil {
	// 	log.Fatalf("Failed to get *sql.DB: %v", err)
	// }

	// if err := sqlDB.Close(); err != nil {
	// 	log.Fatalf("failed to close the database connection: %v", err)
	// }

	fmt.Println("Database connected!")

	// Migrate the database
	DB.AutoMigrate(&model.Book{})
	fmt.Println("Database Migrated.")
}
