package main

import (
	"log"

	"github.com/KainNhantumbo/books-api/config"
	"github.com/KainNhantumbo/books-api/database"
	"github.com/KainNhantumbo/books-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func main() {
	// Connect to the Database
	database.ConnectDB()
	var canFork bool = false

	var allowedDomains string = config.Config("ALLOWED_DOMAINS")

	if env := config.Config("GO_ENV"); env == "production" {
		canFork = true
	}

	app := fiber.New(fiber.Config{
		EnableTrustedProxyCheck: false,
		Prefork:                 canFork,
	})

	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedDomains,
		AllowCredentials: true,
		AllowMethods:     "GET,POST,DELETE,PATCH",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	app.Use(swagger.New(swagger.Config{
		Title: "Books Library API Docs",
		URL:   "/docs/swagger.json",
	}))

	router.AppRouter(app)

	port := config.Config("PORT")
	err := app.Listen(":" + port)

	if err != nil {
		log.Fatal(err)
	}
}
