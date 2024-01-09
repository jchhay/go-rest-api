package app

import (
	"errors"

	"jchhay/go-rest-api-gin/internal/models"
	"jchhay/go-rest-api-gin/internal/repository"
)

type BookService interface {
	GetBooks() (*[]models.Book, error)
	GetBookById(id int) (*models.Book, error)
	CreateNewBook(models.Book) (*models.Book, error)
	DeleteBook(id int) (int64, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}

func (b *bookService) GetBooks() (*[]models.Book, error) {
	return b.repo.FindAll()
}

func (b *bookService) GetBookById(id int) (*models.Book, error) {

	book, err := b.repo.FindByID(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return book, nil
}

func (b *bookService) CreateNewBook(book models.Book) (*models.Book, error) {

	newBook, err := b.repo.Save(book)
	if err != nil {
		return newBook, err
	}

	return newBook, nil
}

func (b *bookService) DeleteBook(id int) (int64, error) {
	return b.repo.DeleteByID(id)
}
