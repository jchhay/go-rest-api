package repository

import (
	"jchhay/go-rest-api-gin/internal/models"
)

type BookRepository interface {
	FindAll() (*[]models.Book, error)
	FindByID(id int) (*models.Book, error)
	Save(book models.Book) (*models.Book, error)
	DeleteByID(id int) (int64, error)
}
