package routes

import (
	"github.com/gin-gonic/gin"
	"library/handler"
	"library/repository"
	"library/service"
)

func SetupRoutes(router *gin.Engine) {

	bookRepository := repository.NewBookRepository()

	bookService := service.NewBookService(bookRepository)

	bookHandler := handler.NewBookHandler(bookService)

	router.GET("/books", bookHandler.GetBookHandler)
	router.DELETE("/books", bookHandler.DeleteBookHandler)
}
