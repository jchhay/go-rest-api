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

func NewRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Configure Repositories
	dbType := config.GetConfig().Database.Driver
	bookRepository := factory.NewRepositoryFactory(dbType)

	// Configure Services
	bookService := app.NewBookService(bookRepository)

	// Instantiate controller with dependencies
	bc := app.NewBookHandler(bookService)

	// Configure Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"

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
