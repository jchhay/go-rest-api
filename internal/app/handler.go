package app

import (
	"net/http"
	"strconv"

	"jchhay/go-rest-api-gin/internal/models"

	"github.com/gin-gonic/gin"
)

type BookRequestBody struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func (b *BookRequestBody) ToBookModel() *models.Book {
	return &models.Book{
		Title:    b.Title,
		Author:   b.Author,
		Quantity: b.Quantity,
	}

}

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

type BookHandler interface {
	GetBooks(c *gin.Context)
	BookById(c *gin.Context)
	CreateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
}

type bookHandler struct {
	services BookService
}

func NewBookHandler(services BookService) BookHandler {
	return &bookHandler{services}
}

// GetBooks 	godoc
// @Summary 	Retrieves all books
// @Description Get books
// @Tags 		Books
// @Produce 	application/json
// @Success 	200 {object} []BookResponseBody
// @Router 		/books [get]
func (bc *bookHandler) GetBooks(c *gin.Context) {
	books, err := bc.services.GetBooks()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "books not found"})
		return
	}
	var resp []BookResponseBody
	for _, book := range *books {
		resp = append(resp, *CreateBookResponseBody(&book))
	}
	c.IndentedJSON(http.StatusOK, resp)
}

// BookById 	godoc
// @Summary 	Retrieves users based on book id
// @Description Get book by id
// @Param 		id path int true "Book ID"
// @Tags 		Books
// @Produce 	application/json
// @Success 	200 {object} BookResponseBody
// @Router 		/books/{id} [get]
func (bc *bookHandler) BookById(c *gin.Context) {
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

	c.IndentedJSON(http.StatusOK, CreateBookResponseBody(book))

}

// CreateBook godoc
// @Summary 	Create book
// @Description Save book to database
// @Param 		book body BookRequestBody true "Book to save"
// @Tags 		Books
// @Produce 	application/json
// @Success 	201 {object} BookResponseBody
// @Router 		/books [post]
func (bc *bookHandler) CreateBook(c *gin.Context) {

	var newBook BookRequestBody

	if err := c.BindJSON(&newBook); err != nil {
		// If there is an error, return err defaults from c.BindJSON
		return
	}

	book, err := bc.services.CreateNewBook(*newBook.ToBookModel())

	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, CreateBookResponseBody(book))
}

// DeleteBook	godoc
// @Summary 	Delete book
// @Description Delete book by id
// @Param 		id path int true "Book ID"
// @Tags 		Books
// @Produce 	application/json
// @Success 	200 {object} string
// @Router 		/books/{id} [delete]
func (bc *bookHandler) DeleteBook(c *gin.Context) {
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
