package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schedulii/src/models"
	"schedulii/src/services"
)

type EventHandler struct {
	es services.EventService
}

func NewEventHandler(es services.EventService) EventHandler {
	return EventHandler{
		es: es,
	}
}

func (eh *EventHandler) HandleReadEvent() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var e models.Event
		err := c.ShouldBindQuery(&e)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"query error": err.Error()})
			return
		}
		event, err := eh.es.ReadEvent(e)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"read error": err.Error()})
			return
		}
		c.JSON(200, event)
	}
	return gin.HandlerFunc(fn)
}
