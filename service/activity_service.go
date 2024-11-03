package service

import "library/repository"

type ActivityService struct {
	activityRepository *repository.ActivityRepository
}

func NewActivityService(activityRepository *repository.ActivityRepository) *ActivityService {
	return &ActivityService{activityRepository: activityRepository}
}

func (activityService *ActivityService) LogUserAction(username, action string) error {
	return activityService.activityRepository.LogUserAction(username, action)
}

func (activityService *ActivityService) GetLastUserActions(username string) ([]string, error) {
	return activityService.activityRepository.GetLastUserActions(username)
}
