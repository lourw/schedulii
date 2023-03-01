package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"schedulii/src/models"
	"schedulii/src/services"
)

type GroupHandler struct {
	gs services.GroupService
}

func NewGroupHandler(gs services.GroupService) GroupHandler {
	return GroupHandler{
		gs: gs,
	}
}

func (gh *GroupHandler) HandleReadGroup() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var g models.Group
		err := c.ShouldBindQuery(&g)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		group, err := gh.gs.ReadGroup(g)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(200, group)
	}
	return gin.HandlerFunc(fn)
}
