package routes

import (
	"github.com/gin-gonic/gin"
	"library/handler"
)

const (
	BooksBaseRoute   = "/books"
	Slash            = "/"
	StoreStatsRoute  = "/store"
	SearchBooksRoute = "/search"
	ActivityRoute    = "/activity"
)

func SetupRoutes(router *gin.Engine, bookHandler *handler.BookHandler, activityHandler *handler.ActivityHandler) {

	bookRoutes := router.Group(BooksBaseRoute)
	{
		bookRoutes.GET(Slash, bookHandler.GetBookHandler)
		bookRoutes.DELETE(Slash, bookHandler.DeleteBookHandler)
		bookRoutes.PUT(Slash, bookHandler.UpdateBookHandler)
		bookRoutes.POST(Slash, bookHandler.AddBookHandler)
	}

	router.GET(StoreStatsRoute, bookHandler.StoreStatsHandler)
	router.GET(SearchBooksRoute, bookHandler.SearchBooksHandler)
	router.GET(ActivityRoute, activityHandler.ActivityHandler)
}
