package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"library/service"
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
		username := c.Query("username")
		// if not found try to get from JSON body
		if username == "" {
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
		if username == "" || c.FullPath() == "/activity" || c.FullPath() == "" {
			c.Next()
			return
		}
		action := c.Request.Method + " " + c.FullPath()
		//fmt.Println(action)
		err := middleware.activityService.LogUserAction(username, action)
		//fmt.Println(err)
		if err != nil {
			return
		}
		c.Next()
	}
}
