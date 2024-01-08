package api

import (
	"net/http"
	"strconv"

	"jchhay/go-rest-api-gin/dto"
	"jchhay/go-rest-api-gin/services"

	"github.com/gin-gonic/gin"
)

type BookController interface {
	getBooks(c *gin.Context)
	bookById(c *gin.Context)
	createBook(c *gin.Context)
	deleteBook(c *gin.Context)
}

type bookController struct {
	services services.BookService
}

func NewBookController(services services.BookService) BookController {
	return &bookController{services}
}

/*
Get All Books
*/
func (bc *bookController) getBooks(c *gin.Context) {
	books, err := bc.services.GetBooks()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "books not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

/*
Get Book By Id
Params: id
*/
func (bc *bookController) bookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	book, err := bc.services.GetBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, book)

}

/*
Create New Book
Params: id, title, author, quantity
*/
func (bc *bookController) createBook(c *gin.Context) {

	var newBook dto.BookRequestBody

	if err := c.BindJSON(&newBook); err != nil {
		// If there is an error, return err defaults from c.BindJSON
		return
	}

	book, err := bc.services.CreateNewBook(newBook)

	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, book)
}

/*
Delete Book By Id
Params: id
*/
func (bc *bookController) deleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	_, err = bc.services.DeleteBook(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "book deleted"})
}
