package main

import (
	"github.com/KainNhantumbo/books-api/controllers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: false,
		Prefork:                 true,
	})

	app.Get("/api/v1/books", controllers.GetBooks)

	log.Fatal(app.Listen(":8080"))
}
