package book

import (
	"time"

	"github.com/KainNhantumbo/books-api/database"
	"github.com/KainNhantumbo/books-api/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Book struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAvailable bool   `json:"isAvailable"`
}

var data []Book = []Book{
	{Id: "1", Name: "The Great Gatsby", Description: "A novel written by F. Scott Fitzgerald.", IsAvailable: true},
	{Id: "2", Name: "To Kill a Mockingbird", Description: "A novel by Harper Lee published in 1960..", IsAvailable: false},
}

func FindAll(c *fiber.Ctx) error {
	db := database.DB
	var books = new([]model.Book)

	db.Find(&books)
	return c.JSON(&books)
}

func Create(c *fiber.Ctx) error {
	db := database.DB
	book := new(model.Book)
	err := c.BodyParser(&book)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Data parse error."})
	}

	db.Create(&book)
	return c.Status(201).JSON(fiber.Map{"status": 201, "message": "Book created successfully."})
}

func UpdateOne(c *fiber.Ctx) error {
	type UpdateBook struct {
		Name          string    `json:"name"`
		Description   string    `json:"description"`
		IsAvailable   bool      `json:"isAvailable"`
		Publisher     string    `json:"publisher"`
		Country       string    `json:"country"`
		Category      string    `json:"category"`
		PublishedDate time.Time `json:"publishedDate"`
		PageCount     int32     `json:"pageCount"`
	}

	id := c.Params("id")
	db := database.DB
	book := new(model.Book)

	if book.ID == uuid.Nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Invalid request ID."})
	}
	// check if this book really exists
	result := db.First(&book, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "Book not found."})
	}

	updateBookData := new(UpdateBook)
	err := c.BodyParser(&updateBookData)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Data parse error."})
	}

	book.Name = updateBookData.Name
	book.Description = updateBookData.Description
	book.IsAvailable = updateBookData.IsAvailable
	book.Publisher = updateBookData.Publisher
	book.PublishedDate = updateBookData.PublishedDate
	book.PageCount = updateBookData.PageCount
	book.Country = updateBookData.Country
	book.Category = updateBookData.Category

	db.Save(&book)
	return c.Status(200).JSON(data)
}

func DeleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	book := new(model.Book)

	if book.ID == uuid.Nil {
		return c.Status(400).JSON(fiber.Map{"status": 400, "message": "Invalid request ID."})
	}
	// check if this book really exists
	result := db.First(&book, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"status": 404, "message": "Book not found."})
	}

	result = db.Delete(&model.Book{})

	if result.Error != nil {
		c.Status(500).JSON(fiber.Map{"status": 500, "message": "Failed to delete book."})
	}
	return c.SendStatus(204)

}
