package memory

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

type bookRepository struct {
	books []models.Book
}

func NewBookRepository() repository.BookRepository {
	return &bookRepository{
		books: books,
	}
}

func (b *bookRepository) FindAll() (*[]models.Book, error) {
	return &b.books, nil
}

func (b *bookRepository) FindByID(id int) (*models.Book, error) {
	for _, book := range b.books {
		if book.ID == id {
			return &book, nil
		}
	}
	return &models.Book{}, errors.New("book not found")
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
	for i, book := range b.books {
		if book.ID == id {
			b.books = append(b.books[:i], b.books[i+1:]...)
			return 0, nil
		}
	}
	return 1, nil
}
