package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"library/repository"
	"net/http"
)

const (
	UsernameParam       = "username"
	ActivityEndpoint    = "/activity"
	EmptyString         = ""
	Space               = " "
	ErrorField          = "error"
	ErrUserNameRequired = "Username is a required field."
)

type LogUserActionMiddleware struct {
	activityRepository repository.ActivityRepository
}

func NewLogUserActionMiddleware(activityRepository repository.ActivityRepository) *LogUserActionMiddleware {
	return &LogUserActionMiddleware{activityRepository: activityRepository}
}

func (middleware *LogUserActionMiddleware) LogUserActionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// trying to get username from query Params
		username := c.Query(UsernameParam)
		// if not found try to get from JSON body
		if username == EmptyString {
			//this is important because I cant get body json twice
			var bodyBytes []byte
			if c.Request.Body != nil {
				bodyBytes, _ = io.ReadAll(c.Request.Body)
			}
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			var requestBody struct {
				Username string `json:"username"`
			}
			if err := json.Unmarshal(bodyBytes, &requestBody); err == nil {
				username = requestBody.Username
			}
		}
		if username == EmptyString {
			c.JSON(http.StatusBadRequest, gin.H{ErrorField: ErrUserNameRequired})
			c.Abort() // Prevents further handler/middleware processing
			return
		}
		if c.FullPath() == ActivityEndpoint || c.FullPath() == EmptyString {
			c.Next()
			return
		}
		action := c.Request.Method + Space + c.FullPath()
		err := middleware.activityRepository.LogUserAction(username, action)
		if err != nil {
			return
		}
		c.Next()
	}
}
