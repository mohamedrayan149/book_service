package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"library/service"
)

const (
	UsernameParam    = "username"
	ActivityEndpoint = "/activity"
	EmptyString      = ""
	Space            = " "
)

type LogUserActionMiddleware struct {
	activityService *service.ActivityService
}

func NewLogUserActionMiddleware(activityService *service.ActivityService) *LogUserActionMiddleware {
	return &LogUserActionMiddleware{activityService: activityService}
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
		if username == EmptyString || c.FullPath() == ActivityEndpoint || c.FullPath() == EmptyString {
			c.Next()
			return
		}
		action := c.Request.Method + Space + c.FullPath()
		err := middleware.activityService.LogUserAction(username, action)
		if err != nil {
			return
		}
		c.Next()
	}
}
