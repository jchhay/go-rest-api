package dto

import "jchhay/go-rest-api-gin/models"

type BookResponseBody struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func CreateBookResponseBody(book *models.Book) *BookResponseBody {
	return &BookResponseBody{
		ID:       book.ID,
		Title:    book.Title,
		Author:   book.Author,
		Quantity: book.Quantity,
	}
}
