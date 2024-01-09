package sqlite

import (
	"errors"
	"jchhay/go-rest-api-gin/internal/models"
	"jchhay/go-rest-api-gin/internal/repository"
	"jchhay/go-rest-api-gin/pkg/db"
)

type SqliteDatabaseFactory struct{}

func (m *SqliteDatabaseFactory) NewBookRepository(dbType string) repository.BookRepository {
	return NewBookRepository(dbType)
}

type bookRepository struct {
}

func NewBookRepository(dbType string) repository.BookRepository {
	db.NewGormClient(dbType)
	return &bookRepository{}
}

func (b *bookRepository) FindAll() (*[]models.Book, error) {
	var books []models.Book
	err := repository.Find(&models.Book{}, &books, "ID asc")
	return &books, err
}

func (b *bookRepository) FindByID(id int) (*models.Book, error) {
	var book models.Book
	err := repository.FirstByID(&book, id)
	if err != nil {
		return nil, errors.New("book not found")
	}
	return &book, nil
}

func (b *bookRepository) Save(book models.Book) (*models.Book, error) {
	err := repository.Save(&book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *bookRepository) DeleteByID(id int) (int64, error) {
	count, err := repository.DeleteByID(&models.Book{}, id)
	if err != nil {
		return 0, err
	}
	return count, nil

}
