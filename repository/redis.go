package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"library/connectors"
)

// Constants for Redis key patterns and action limits
const (
	UserActionsKeyPattern = "user:%s:actions"
	ActionHistoryLimit    = -3
	EndOfList             = -1
)

type ActivityRedisRepo struct {
	redisClient *redis.Client
}

func NewActivityRedisRepo() *ActivityRedisRepo {
	client := connectors.InitRedisClient()
	return &ActivityRedisRepo{redisClient: client}
}

func (rc *ActivityRedisRepo) LogUserAction(username, action string) error {
	key := fmt.Sprintf(UserActionsKeyPattern, username)
	err := rc.redisClient.RPush(context.Background(), key, action).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc *ActivityRedisRepo) GetLastUserActions(username string) ([]string, error) {
	key := fmt.Sprintf(UserActionsKeyPattern, username)
	actions, err := rc.redisClient.LRange(context.Background(), key, ActionHistoryLimit, EndOfList).Result()
	if err != nil {
		return nil, err
	}
	rc.redisClient.LTrim(context.Background(), key, ActionHistoryLimit, EndOfList)
	return actions, nil
}
