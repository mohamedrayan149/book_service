package redis

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

var Client *redis.Client

func InitRedisClient() {
	Client = redis.NewClient(&redis.Options{
		Addr: RedisDevAddr,
		DB:   RedisDB,
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf(ErrorConnectingRedis, err)
	}
}
