package repository

import (
	"errors"
	"jchhay/go-rest-api-gin/models"
	"sync"
)

// Dummy Data
var books = []models.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "Ulysses", Author: "James Joyce", Quantity: 5},
	{ID: "3", Title: "Don Quixote", Author: "Miguel de Cervantes", Quantity: 10},
	{ID: "4", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 8},
}

type BookRepository interface {
	GetBookRepository() []models.Book
	GetBookRepositoryById(id string) (*models.Book, error)
	AddBookRepository(book models.Book) (*models.Book, error)
}

var bookInstance *bookRepository
var once sync.Once

// GetBookRepositoryInstance returns a singleton instance of the BookRepository interface.
func GetBookRepositoryInstance() BookRepository {
	once.Do(func() {
		bookInstance = &bookRepository{}
	})
	return bookInstance
}

func GetBooks(repo BookRepository) []models.Book {
	return repo.GetBookRepository()
}

func GetBookById(repo BookRepository, id string) (*models.Book, error) {
	return repo.GetBookRepositoryById(id)
}

func AddBook(repo BookRepository, book models.Book) (*models.Book, error) {
	return repo.AddBookRepository(book)
}

// bookRepository implements the BookRepository interface.
type bookRepository struct{}

func (r *bookRepository) GetBookRepository() []models.Book {
	return books
}

func (r *bookRepository) GetBookRepositoryById(id string) (*models.Book, error) {
	for _, book := range books {
		if book.ID == id {
			return &book, nil
		}
	}
	return &models.Book{}, errors.New("book not found")
}

func (r *bookRepository) AddBookRepository(book models.Book) (*models.Book, error) {
	for _, b := range books {
		if book.ID == b.ID {
			return &book, errors.New("book id already exists")
		}
	}
	books = append(books, book)
	return &book, nil
}
