package google

import (
	"fmt"
	"log"
	"net/http"
	"schedulii/src/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func GoogleOauthLoginHandler(c *gin.Context) {
	session := sessions.Default(c)
	config := utils.ReadGoogleAPICredentials()

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	session.Set("redirect", c.Request.URL)
	c.Redirect(http.StatusFound, authURL)
	c.Abort()
}

func GoogleCallbackHandler(c *gin.Context) {
	session := sessions.Default(c)
	status := c.Writer.Status()
	if status != http.StatusOK {
		fmt.Print("Status not 200 when getting auth callback from Google API")
	}

	config := utils.ReadGoogleAPICredentials()

	authCode := c.Request.URL.Query().Get("code")
	tok, err := config.Exchange(c, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	utils.SaveOauthTokenToSession(session, tok)

	redirect, ok := session.Get("redirect").(string)
	if !ok {
		redirect = "/"
	}
	c.Redirect(http.StatusFound, redirect)
}
