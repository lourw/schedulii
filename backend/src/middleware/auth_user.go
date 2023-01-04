package middleware

import (
	"net/http"
	"schedulii/src/utils"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckAuthenticated(c *gin.Context) {
	_, err := utils.ExtractUserJWT(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.Next()
}

func CheckGoogleAuthenticated(c *gin.Context) {
	session := sessions.Default(c)

	tok, ok := utils.OauthTokenFromSession(session)
	if !ok || tok.Expiry.Before(time.Now()) {
		c.Redirect(http.StatusFound, "/googleAuth")
		c.Abort()
	}
}
