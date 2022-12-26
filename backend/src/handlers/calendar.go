package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCalendars(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
