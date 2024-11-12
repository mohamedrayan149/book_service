package datastore

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"library/config"
	"library/connectors"
)

type UserActivityRedis struct {
	redisClient *redis.Client
}

func NewUserActivityRedis() *UserActivityRedis {
	client := connectors.InitRedisClient()
	return &UserActivityRedis{redisClient: client}
}

func (rc *UserActivityRedis) LogUserAction(username, action string) error {
	key := fmt.Sprintf(config.UserActionsKeyPattern, username)
	err := rc.redisClient.RPush(context.Background(), key, action).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc *UserActivityRedis) GetLastUserActions(username string) ([]string, error) {
	key := fmt.Sprintf(config.UserActionsKeyPattern, username)
	actions, err := rc.redisClient.LRange(context.Background(), key, config.ActionHistoryLimit, config.EndOfList).Result()
	if err != nil {
		return nil, err
	}
	rc.redisClient.LTrim(context.Background(), key, config.ActionHistoryLimit, config.EndOfList)
	return actions, nil
}
