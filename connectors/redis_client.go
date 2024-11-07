package connectors

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

const (
	RedisDevAddr         = "redis-search.fiverrdev.com:6382"
	RedisDB              = 0
	ErrorConnectingRedis = "Error connecting to Redis: %v"
)

func InitRedisClient() *redis.Client {
	RedisClient := redis.NewClient(&redis.Options{
		Addr: RedisDevAddr,
		DB:   RedisDB,
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf(ErrorConnectingRedis, err)
	}
	return RedisClient
}
