package sqlite

import (
	"errors"
	"jchhay/go-rest-api-gin/internal/models"
	"jchhay/go-rest-api-gin/internal/repository"
)

// Dummy Data
var books = []models.Book{
	{ID: 1, Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: 2, Title: "Ulysses", Author: "James Joyce", Quantity: 5},
	{ID: 3, Title: "Don Quixote", Author: "Miguel de Cervantes", Quantity: 10},
	{ID: 4, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 8},
}

type PostGresDatabaseFactory struct{}

func (m *PostGresDatabaseFactory) NewBookRepository() repository.BookRepository {
	return NewBookRepository()
}

type bookRepository struct {
	books []models.Book
}

func NewBookRepository() repository.BookRepository {
	return &bookRepository{
		books: books,
	}
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
	for _, b := range b.books {
		if book.ID == b.ID {
			return &book, errors.New("book id already exists")
		}
	}
	b.books = append(b.books, book)
	return &book, nil
}

func (b *bookRepository) DeleteByID(id int) (int64, error) {
	count, err := repository.DeleteByID(&models.Book{}, id)
	if err != nil {
		return 0, err
	}
	return count, nil

}
