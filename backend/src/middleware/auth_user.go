package middleware

import (
	"net/http"

	"schedulii/src/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func CheckAuthenticated(c *gin.Context) {
	if isUserAuthenticated(c.Request) {
		c.Next()
	}
	c.Redirect(http.StatusFound, "/")
}

func isUserAuthenticated(req *http.Request) bool {
	_, ok := utils.ExtractUserJWT(req)
	return ok
}

func CheckGoogleAuthenticated(c *gin.Context) {
	session := sessions.Default(c)
	tok, ok := utils.OauthTokenFromSession(session)

	if ok {
		if isUserGoogleAuthenticated(tok) {
			c.Next()
		}
	}

	c.Redirect(http.StatusFound, "/googleAuth")
}

func isUserGoogleAuthenticated(tok *oauth2.Token) bool {
	return tok.Valid()
}
