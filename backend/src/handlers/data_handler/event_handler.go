package data_handler

import (
	"net/http"
	"schedulii/src/models/data_model"
	"schedulii/src/services/data_srv"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	es data_srv.EventService
}

func NewEventHandler(es data_srv.EventService) EventHandler {
	return EventHandler{
		es: es,
	}
}

func (eh *EventHandler) HandleReadEvent() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var e data_model.Event
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
