package dto

import (
	"jchhay/go-rest-api-gin/models"
)

type BookRequestBody struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// Map to Body Model
func (b *BookRequestBody) ToBookModel() *models.Book {
	return &models.Book{
		ID:       b.ID,
		Title:    b.Title,
		Author:   b.Author,
		Quantity: b.Quantity,
	}

}
