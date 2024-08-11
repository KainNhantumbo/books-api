package main

import (
	"github.com/KainNhantumbo/books-api/router"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: false,
		Prefork:                 true,
	})

	router.BookRouter(app)

	err := app.Listen(":8080")

	if err != nil {
		log.Fatal(err)
	}

}
