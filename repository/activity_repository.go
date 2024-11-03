package repository

import (
	"context"
	"fmt"
	"library/redis"
)

type ActivityRepository struct {
}

func NewActivityRepository() *ActivityRepository {
	return &ActivityRepository{}
}

func (activityRepository *ActivityRepository) LogUserAction(username, action string) error {
	key := fmt.Sprintf("user:%s:actions", username)
	err := redis.Client.RPush(context.Background(), key, action).Err()
	if err != nil {
		return err
	}
	return nil
}

func (activityRepository *ActivityRepository) GetLastUserActions(username string) ([]string, error) {
	key := fmt.Sprintf("user:%s:actions", username)
	actions, err := redis.Client.LRange(context.Background(), key, -3, -1).Result()
	if err != nil {
		return nil, err
	}
	redis.Client.LTrim(context.Background(), key, -3, -1)
	return actions, nil
}
