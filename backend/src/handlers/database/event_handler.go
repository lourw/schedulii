package database

import (
	"net/http"
	"schedulii/src/models"
	"schedulii/src/services/data/events"

	"github.com/gin-gonic/gin"
)

func ReadEventHandler(env *models.Env) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var event models.Event
		err := c.ShouldBindQuery(&event)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"query error": err.Error()})
			return
		}
		event, err = events.ReadEvent(env, event)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"read error": err.Error()})
			return
		}
		c.JSON(200, event)
	}
	return gin.HandlerFunc(fn)
}
