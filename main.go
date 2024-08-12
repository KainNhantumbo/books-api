package main

import (
	"log"

	"github.com/KainNhantumbo/books-api/config"
	"github.com/KainNhantumbo/books-api/database"
	"github.com/KainNhantumbo/books-api/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: false,
		Prefork:                 true,
	})

	router.BookRouter(app)

	// Connect to the Database
	database.ConnectDB()

	port := config.Config("PORT")
	err := app.Listen(":" + port)

	if err != nil {
		log.Fatal(err)
	}

}
