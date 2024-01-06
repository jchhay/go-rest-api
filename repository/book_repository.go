package repository

import (
	"jchhay/go-rest-api-gin/models"
)

type BookRepository interface {
	FindAll() []models.Book
	FindByID(id int) (*models.Book, error)
	Save(book models.Book) (*models.Book, error)
}
