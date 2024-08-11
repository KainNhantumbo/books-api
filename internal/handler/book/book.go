package book

import "github.com/gofiber/fiber/v2"

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
	return c.JSON(&data)
}

func Create(c *fiber.Ctx) error {
	book := new(Book)
	err := c.BodyParser(book)

	if err != nil {
		return err
	}

	data := append(data, *book)
	return c.JSON(data)
}

func UpdateOne(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(400).SendString("Invalid request Id.")
	}

	book := new(Book)

	err := c.BodyParser(book)

	if err != nil {
		return err
	}

	for i, item := range data {
		if item.Id == id {
			data[i] = *book
			break
		}
	}
	return c.Status(200).JSON(data)
}

func DeleteOne(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(400).SendString("Invalid request Id.")
	}

	var updatedBooks []Book = []Book{}

	// Iterate over the original slice
	for _, book := range data {
		if book.Id != id {
			updatedBooks = append(updatedBooks, book)
		}
	}

	data = updatedBooks
	return c.JSON(updatedBooks)

}
