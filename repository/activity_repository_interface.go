package repository

type ActivityRepository interface {
	LogUserAction(username, action string) error
	GetLastUserActions(username string) ([]string, error)
}
