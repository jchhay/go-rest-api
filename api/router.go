package api

import (
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	router := gin.Default()

	// ================== Book Routes
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)

	return router
}
