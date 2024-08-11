package controllers

import "github.com/gofiber/fiber/v2"

type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsAvailable bool   `json:"isAvailable"`
}

var data []Book = []Book{
	{Id: 1, Name: "The Great Gatsby", Description: "A novel written by F. Scott Fitzgerald.", IsAvailable: true},
	{Id: 2, Name: "To Kill a Mockingbird", Description: "A novel by Harper Lee published in 1960..", IsAvailable: false},
}

func GetBooks(c *fiber.Ctx) error {
	return c.JSON(data)
}

func CreateBook(c *fiber.Ctx) error {
	body := c.Request().Body()
	err := c.BodyParser(body)

	if err != nil {
		return err
	}

	data := append(data, )
	return c.JSON(data)
}

func UpdateBook(id int, bookData Book) []Book {
	for i, book := range data {
		if book.Id == id {
			data[i] = bookData
			break // Exit the loop after updating the book
		}
	}
	return data
}

func DeleteBook(id int) {
	var updatedBooks []Book

	// Iterate over the original slice
	for _, book := range data {
		if book.Id != id {
			updatedBooks = append(updatedBooks, book)
		}
	}

	data = updatedBooks
}
