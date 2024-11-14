package service

import (
	"github.com/gin-gonic/gin"
	"library/config"
)

func Routes(router *gin.Engine, handler *Handler) {

	bookRoutes := router.Group(config.BooksBaseRoute)
	{
		bookRoutes.GET(config.Slash, handler.GetBook)
		bookRoutes.DELETE(config.Slash, handler.DeleteBook)
		bookRoutes.PUT(config.Slash, handler.UpdateBook)
		bookRoutes.POST(config.Slash, handler.AddBook)
	}

	router.GET(config.StoreStatsRoute, handler.StoreStats)
	router.GET(config.SearchBooksRoute, handler.SearchBooks)
	router.GET(config.ActivityRoute, handler.Activity)
}
