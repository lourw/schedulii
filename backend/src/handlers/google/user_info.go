package google

import (
	"net/http"
	"schedulii/src/services/google_srv"
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

	userInfo := google_srv.GetGoogleUserData(tok)
	c.JSON(http.StatusOK, userInfo)
}
