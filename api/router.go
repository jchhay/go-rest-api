package api

import (
	"jchhay/go-rest-api-gin/config"
	"jchhay/go-rest-api-gin/repository"
	"jchhay/go-rest-api-gin/repository/memory"
	"jchhay/go-rest-api-gin/services"

	"github.com/gin-gonic/gin"
)

func useDatabase() bool {
	return config.GetConfig().Server.UseDatabase
}

func Setup() *gin.Engine {

	router := gin.Default()

	// Configure Repositories
	var databaseFactory repository.RepositoryFactory

	if useDatabase() {
		// Placeholder for actual db
	} else {
		databaseFactory = &memory.MemoryDatabaseFactory{}
	}

	bookRepository := databaseFactory.NewBookRepository()

	// Configure Services
	bookService := services.NewBookService(bookRepository)

	// Instantiate controller with dependencies
	bc := NewBookController(bookService)

	// Register Routes
	router.GET("/books", bc.getBooks)
	router.GET("/books/:id", bc.bookById)
	router.POST("/books", bc.createBook)

	return router
}
