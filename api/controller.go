package api

import (
	"net/http"

	"jchhay/go-rest-api-gin/models"
	"jchhay/go-rest-api-gin/services"

	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	books := services.GetBooks()
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := services.GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)

}

func createBook(c *gin.Context) {

	var newBook models.Book
	// Call BindJSON to bind the received JSON to
	// newBook.
	if err := c.BindJSON(&newBook); err != nil {
		// If there is an error, return err defaults from c.BindJSON
		return
	}

	book, err := services.AddBook(newBook)

	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, book)
}
