package connectors

import (
	"context"
	"github.com/go-redis/redis/v8"
	"library/config"
	"log"
)

func InitRedisClient() *redis.Client {
	RedisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisDevAddr,
		DB:   config.RedisDB,
	})
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf(config.ErrorConnectingRedis, err)
	}
	return RedisClient
}
