package repository

import (
	"errors"
	"jchhay/go-rest-api-gin/models"
)

var books = []models.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "Ulysses", Author: "James Joyce", Quantity: 5},
	{ID: "3", Title: "Don Quixote", Author: "Miguel de Cervantes", Quantity: 10},
	{ID: "4", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 8},
}

func GetBookRepository() []models.Book {
	return books
}

func GetBookRepositoryById(id string) (*models.Book, error) {
	for _, book := range books {
		if book.ID == id {
			return &book, nil
		}
	}
	return &models.Book{}, errors.New("book not found")

}
