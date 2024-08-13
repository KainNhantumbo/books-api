package main

import (
	"log"

	"github.com/KainNhantumbo/books-api/config"
	"github.com/KainNhantumbo/books-api/database"
	"github.com/KainNhantumbo/books-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Connect to the Database
	database.ConnectDB()

	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: false,
		Prefork:                 true,
	})

	app.Use(recover.New())

	router.BookRouter(app)

	port := config.Config("PORT")
	err := app.Listen(":" + port)

	if err != nil {
		log.Fatal(err)
	}
}
