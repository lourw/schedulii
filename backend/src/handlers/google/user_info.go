package google

import (
	"net/http"
	"schedulii/src/services/google"
	"schedulii/src/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserInfoHandler(c *gin.Context) {
	session := sessions.Default(c)
	tok, ok := utils.OauthTokenFromSession(session)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid oauth token"})
	}

	userInfo := google.GetGoogleUserData(tok)
	c.JSON(http.StatusOK, userInfo)
}
