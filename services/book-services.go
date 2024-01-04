package services

import (
	"errors"

	"jchhay/go-rest-api-gin/models"
	"jchhay/go-rest-api-gin/repository"
)

func GetBooks() []models.Book {
	repo := repository.GetBookRepositoryInstance()
	books := repository.GetBooks(repo)

	return books
}

func GetBookById(id string) (*models.Book, error) {
	repo := repository.GetBookRepositoryInstance()

	book, err := repository.GetBookById(repo, id)
	if err != nil {
		return nil, errors.New("book not found")
	}

	return book, nil
}

func AddBook(book models.Book) (*models.Book, error) {
	repo := repository.GetBookRepositoryInstance()
	newBook, err := repository.AddBook(repo, book)
	if err != nil {
		return &book, err
	}

	return newBook, nil
}
