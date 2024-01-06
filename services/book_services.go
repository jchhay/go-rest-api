package services

import (
	"errors"

	"jchhay/go-rest-api-gin/dto"
	"jchhay/go-rest-api-gin/models"
	"jchhay/go-rest-api-gin/repository"
)

type BookService interface {
	GetBooks() []models.Book
	GetBookById(id int) (*models.Book, error)
	CreateNewBook(dto.BookRequestBody) (*dto.BookResponseBody, error)
}

type bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}

func (b *bookService) GetBooks() []models.Book {
	return b.repo.FindAll()
}

func (b *bookService) GetBookById(id int) (*models.Book, error) {

	book, err := b.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("book not found")
	}

	return book, nil
}

func (b *bookService) CreateNewBook(book dto.BookRequestBody) (*dto.BookResponseBody, error) {

	newBook, err := b.repo.Save(*book.ToBookModel())
	if err != nil {
		return dto.CreateBookResponseBody(newBook), err
	}

	return dto.CreateBookResponseBody(newBook), nil
}
