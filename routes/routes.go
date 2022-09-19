package routes

import (
	"gin_swagger/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/books", handlers.GetBooks)
		v1.GET("/books/:isbn", handlers.GetBookByISBN)
		// router.DELETE("/books/:isbn", handlers.DeleteBookByISBN)
		// router.PUT("/books/:isbn", handlers.UpdateBookByISBN)
		v1.POST("/books", handlers.PostBook)
	}

	return router
}
