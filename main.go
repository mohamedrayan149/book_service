package main

import (
	"github.com/gin-gonic/gin"
	"library/elastic"
	"library/handler"
	"library/redis"
)

func main() {

	// Initialize Elasticsearch and Redis clients
	redis.InitRedisClient()
	elastic.InitElasticClient()
	r := gin.Default()

	// Book routes
	r.POST("/books", handler.AddBookHandler)
	r.PUT("/books", handler.UpdateBookHandler)
	r.GET("/books", handler.GetBookHandler)
	r.DELETE("/books", handler.DeleteBookHandler)
	r.GET("/search", handler.SearchBooksHandler)
	r.GET("/store", handler.StoreStatsHandler)

	//Activity route
	r.GET("/activity", handler.ActivityHandler)

	// Start the server
	r.Run(":8080")
}
