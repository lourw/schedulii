package data_handler

import (
	"net/http"
	"schedulii/src/models/data_model"
	"schedulii/src/services/data_srv"

	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	gs data_srv.GroupService
}

func NewGroupHandler(gs data_srv.GroupService) GroupHandler {
	return GroupHandler{
		gs: gs,
	}
}

func (gh *GroupHandler) HandleReadGroup() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var g data_model.Group
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
