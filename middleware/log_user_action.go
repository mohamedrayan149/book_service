package middleware

import (
	"github.com/gin-gonic/gin"
	"library/config"
	"library/datastore"
	"net/http"
)

type LogUserActionMiddleware struct {
	activityRepository datastore.UserActivity
}

func NewLogUserActionMiddleware(activityRepository datastore.UserActivity) *LogUserActionMiddleware {
	return &LogUserActionMiddleware{activityRepository: activityRepository}
}

func (middleware *LogUserActionMiddleware) LogUserActionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query(config.UsernameParam)
		if username == config.EmptyString {
			c.JSON(http.StatusBadRequest, gin.H{config.ErrorField: config.ErrUserNameRequired})
			c.Abort() // Prevents further handler/middleware processing
			return
		}
		if c.FullPath() == config.ActivityRoute || c.FullPath() == config.EmptyString {
			c.Next()
			return
		}
		action := c.Request.Method + config.Space + c.FullPath()
		err := middleware.activityRepository.LogUserAction(username, action)
		if err != nil {
			return
		}
		c.Next()
	}
}
