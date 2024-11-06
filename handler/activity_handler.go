package handler

import (
	"github.com/gin-gonic/gin"
	"library/service"
	"net/http"
)

// Constants for parameter names and response keys
const (
	UsernameParam         = "username"
	ActionsField          = "actions"
	ErrorUsernameRequired = "username is required"
)

type ActivityHandler struct {
	activityService *service.ActivityService
}

func NewActivityHandler(activityService *service.ActivityService) *ActivityHandler {
	return &ActivityHandler{activityService: activityService}
}

func (activityHandler *ActivityHandler) ActivityHandler(c *gin.Context) {
	username := c.Query(UsernameParam)
	if username == EmptyString {
		c.JSON(http.StatusBadRequest, gin.H{ErrorField: ErrorUsernameRequired})
		return
	}

	actions, err := activityHandler.activityService.GetLastUserActions(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{ErrorField: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{ActionsField: actions})
}
