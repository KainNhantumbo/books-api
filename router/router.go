package router

import (
	"github.com/KainNhantumbo/books-api/internal/handler/book"
	"github.com/gofiber/fiber/v2"
)

func BookRouter(app *fiber.App) {
	app.Get("/books", book.FindAll)
	app.Post("/books", book.Create)
	app.Patch("/books/:id", book.UpdateOne)
	app.Delete("/books/:id", book.DeleteOne)
}
