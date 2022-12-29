package middleware

import (
	"net/http"
	"schedulii/src/utils"

	"github.com/gin-gonic/gin"
)

func CheckAuthenticated(c *gin.Context) {
	_, err := utils.ExtractToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}
