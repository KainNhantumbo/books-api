package router

import (
	"github.com/KainNhantumbo/books-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func AppRouter(app *fiber.App) {
	router := app.Group("/api/v1")

	// book router
	bookRouter := router.Group("/books")
	bookRouter.Get("/", handlers.FindAllBooks)
	bookRouter.Post("/", handlers.CreateBook)
	bookRouter.Patch("/:id", handlers.UpdateOneBook)
	bookRouter.Delete("/:id", handlers.DeleteOneBook)
}
