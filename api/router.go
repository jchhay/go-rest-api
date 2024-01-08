package api

import (
	"jchhay/go-rest-api-gin/config"
	"jchhay/go-rest-api-gin/pkg/db"
	"jchhay/go-rest-api-gin/repository/factory"
	"jchhay/go-rest-api-gin/services"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	router := gin.Default()

	// Configure Repositories
	db.SetupDB()
	dbType := config.GetConfig().Database.Driver
	bookRepository := factory.NewRepositoryFactory(dbType)

	// Configure Services
	bookService := services.NewBookService(bookRepository)

	// Instantiate controller with dependencies
	bc := NewBookController(bookService)

	// Register Routes
	router.GET("/books", bc.getBooks)
	router.GET("/books/:id", bc.bookById)
	router.POST("/books", bc.createBook)
	router.DELETE("/books/:id", bc.deleteBook)

	return router
}
