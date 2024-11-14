package main

import (
	"github.com/gin-gonic/gin"
	"library/config"
	"library/datastore"
	"library/middleware"
	"library/service"
	"log"
)

func main() {

	bookSearchStore := datastore.NewBookStoreElastic()
	userActivityCache := datastore.NewUserActivityRedis()

	handler := service.NewHandler(bookSearchStore, userActivityCache)

	logMiddleware := middleware.NewLogUserActionMiddleware(userActivityCache)

	router := gin.Default()
	router.Use(logMiddleware.LogUserActionMiddleware())
	service.Routes(router, handler)

	err := router.Run(config.ServerPort)
	if err != nil {
		log.Fatal(err)
		return
	}
}
