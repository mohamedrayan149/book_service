package main

import (
	"github.com/gin-gonic/gin"
	"library/elastic"
	"library/redis"
	"library/routes"
)

func main() {
	redis.InitRedisClient()
	elastic.InitElasticClient()
	router := gin.Default()
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
