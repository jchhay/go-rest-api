package api

import (
	"jchhay/go-rest-api-gin/config"
	"jchhay/go-rest-api-gin/docs"
	"jchhay/go-rest-api-gin/internal/app"
	"jchhay/go-rest-api-gin/internal/repository/factory"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title Golang API (jchhay)
// @version 1
// @description This is a sample rest api server using golang and gin framework

// @contact.name API Support
// @contact.email XXXXXXXXXXXXXXXX

// @host localhost:3000
// @BasePath /api/v1
func NewRouter() *gin.Engine {

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Configure Repositories
	dbType := config.GetConfig().Database.Driver
	bookRepository := factory.NewRepositoryFactory(dbType)

	// Configure Services
	bookService := app.NewBookService(bookRepository)

	// Instantiate controller with dependencies
	bc := app.NewBookHandler(bookService)

	// Register Routes

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := router.Group("/api/v1")

	books := v1.Group("/books")
	{
		books.GET("/", bc.GetBooks)
		books.GET("/:id", bc.BookById)
		books.POST("/", bc.CreateBook)
		books.DELETE("/:id", bc.DeleteBook)
	}
	return router
}
