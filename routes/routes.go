package routes

import (
	"github.com/gin-gonic/gin"
	"library/handler"
	"library/middleware"
	"library/repository"
	"library/service"
)

func SetupRoutes(router *gin.Engine) {

	bookRepository := repository.NewBookRepository()
	activityRepository := repository.NewActivityRepository()

	bookService := service.NewBookService(bookRepository)
	activityService := service.NewActivityService(activityRepository)

	bookHandler := handler.NewBookHandler(bookService)
	activityHandler := handler.NewActivityHandler(activityService)

	logMiddleware := middleware.NewLogUserActionMiddleware(activityService)
	router.Use(logMiddleware.LogUserActionMiddleware())

	bookRoutes := router.Group("/books")
	{
		bookRoutes.GET("/", bookHandler.GetBookHandler)
		bookRoutes.DELETE("/", bookHandler.DeleteBookHandler)
		bookRoutes.PUT("/", bookHandler.UpdateBookHandler)
		bookRoutes.POST("/", bookHandler.AddBookHandler)
	}
	router.GET("/store", bookHandler.StoreStatsHandler)
	router.GET("/search", bookHandler.SearchBooksHandler)
	router.GET("/activity", activityHandler.ActivityHandler)
}
