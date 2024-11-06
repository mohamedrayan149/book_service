package repository

import (
	"context"
	"fmt"
	"library/redis"
)

// Constants for Redis key patterns and action limits
const (
	UserActionsKeyPattern = "user:%s:actions"
	ActionHistoryLimit    = -3
	EndOfList             = -1
)

type RedisActivityRepository struct {
}

func NewRedisActivityRepository() *RedisActivityRepository {
	return &RedisActivityRepository{}
}

func (redisActivityRepository *RedisActivityRepository) LogUserAction(username, action string) error {
	key := fmt.Sprintf(UserActionsKeyPattern, username)
	err := redis.Client.RPush(context.Background(), key, action).Err()
	if err != nil {
		return err
	}
	return nil
}

func (redisActivityRepository *RedisActivityRepository) GetLastUserActions(username string) ([]string, error) {
	key := fmt.Sprintf(UserActionsKeyPattern, username)
	actions, err := redis.Client.LRange(context.Background(), key, ActionHistoryLimit, EndOfList).Result()
	if err != nil {
		return nil, err
	}
	redis.Client.LTrim(context.Background(), key, ActionHistoryLimit, EndOfList)
	return actions, nil
}
