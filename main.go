package main

import (
	"github.com/gin-gonic/gin"
	"library/elastic"
	"library/handler"
	"library/middleware"
	"library/redis"
	"library/repository"
	"library/routes"
	"library/service"
)

const ServerPort = ":8080"

func main() {
	redis.InitRedisClient()
	elastic.InitElasticClient()
	router := gin.Default()

	// Initializing Repository layer
	bookRepository := repository.NewElasticsearchBookRepository()
	activityRepository := repository.NewRedisActivityRepository()

	// Initializing Service layer
	bookService := service.NewBookService(bookRepository)
	activityService := service.NewActivityService(activityRepository)

	// Initializing Handler layer
	bookHandler := handler.NewBookHandler(bookService)
	activityHandler := handler.NewActivityHandler(activityService)

	// Initializing Middleware
	logMiddleware := middleware.NewLogUserActionMiddleware(activityService)
	router.Use(logMiddleware.LogUserActionMiddleware())

	routes.SetupRoutes(router, bookHandler, activityHandler)

	router.Run(ServerPort)
}
