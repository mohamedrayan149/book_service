package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var Client *redis.Client

func InitRedisClient() {
	Client = redis.NewClient(&redis.Options{
		Addr: "redis-search.fiverrdev.com:6382",
		DB:   0,
	})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
}
