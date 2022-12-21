package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HealthCheckController(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}
