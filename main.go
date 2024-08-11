package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/KainNhantumbo/go-books-api/controllers"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: false,
	})

	app.Get("/api/v1/books", func(c *fiber.Ctx) error {
		data := controllers.GetBooks()
		return c.JSON(data)
	})

	log.Fatal(app.Listen(":8080"))
}
