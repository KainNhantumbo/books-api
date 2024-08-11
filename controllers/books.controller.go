package controllers

type Book struct {
	id          int
	name        string
	description string
	isAvailable bool
}

var data []Book = []Book{
	{id: 1, name: "The Great Gatsby", description: "A novel written by F. Scott Fitzgerald.", isAvailable: true},
	{id: 2, name: "To Kill a Mockingbird", description: "A novel by Harper Lee published in 1960..", isAvailable: false},
}

func GetBooks() []Book {
	return data
}

func CreateBook(newBook Book) []Book {
	return append(data, newBook)
}

func UpdateBook(id int, bookData Book) []Book {
	for i, book := range data {
		if book.id == id {
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
		if book.id != id {
			updatedBooks = append(updatedBooks, book)
		}
	}

	data = updatedBooks
}
