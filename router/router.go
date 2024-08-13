package router

import (
	"github.com/KainNhantumbo/books-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func BookRouter(app *fiber.App) {
	route := app.Group("/api/v1/books")
	route.Get("/", book.FindAll)
	route.Post("/", book.Create)
	route.Patch("/:id", book.UpdateOne)
	route.Delete("/:id", book.DeleteOne)
}
