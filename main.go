package main

import (
	"github.com/gin-gonic/gin"
	"library/middleware"
	"library/repository"
	"library/service"
	"log"
)

const ServerPort = ":8080"

func main() {

	router := gin.Default()

	bookSearchStore := repository.NewBookElasticRepo()
	userActivityCache := repository.NewActivityRedisRepo()

	handler := service.NewHandler(bookSearchStore, userActivityCache)

	logMiddleware := middleware.NewLogUserActionMiddleware(userActivityCache)
	router.Use(logMiddleware.LogUserActionMiddleware())

	service.SetupRoutes(router, handler)

	err := router.Run(ServerPort)
	if err != nil {
		log.Fatal(err)
		return
	}
}
