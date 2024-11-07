package service

import (
	"github.com/gin-gonic/gin"
)

const (
	BooksBaseRoute   = "/books"
	Slash            = "/"
	StoreStatsRoute  = "/store"
	SearchBooksRoute = "/search"
	ActivityRoute    = "/activity"
)

func SetupRoutes(router *gin.Engine, handler *Handler) {

	bookRoutes := router.Group(BooksBaseRoute)
	{
		bookRoutes.GET(Slash, handler.GetBook)
		bookRoutes.DELETE(Slash, handler.DeleteBook)
		bookRoutes.PUT(Slash, handler.UpdateBook)
		bookRoutes.POST(Slash, handler.AddBook)
	}

	router.GET(StoreStatsRoute, handler.StoreStats)
	router.GET(SearchBooksRoute, handler.SearchBooks)
	router.GET(ActivityRoute, handler.Activity)
}
