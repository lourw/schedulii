package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckAuthenticated(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")
	if username == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}
