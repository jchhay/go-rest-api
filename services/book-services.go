package services

import (
	"errors"

	"jchhay/go-rest-api-gin/models"
	"jchhay/go-rest-api-gin/repository"
)

func GetBookById(id string) (*models.Book, error) {

	book, err := repository.GetBookRepositoryById(id)
	if err != nil {
		return nil, errors.New("book not found")
	}

	return book, nil
}
