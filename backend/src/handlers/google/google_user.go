package google

import (
	"context"
	"log"
	"net/http"
	utils "schedulii/src/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	oauth2 "golang.org/x/oauth2"
	googleOauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func GoogleUserDataHandler(c *gin.Context) {
	session := sessions.Default(c)
	tok, ok := utils.OauthTokenFromSession(session)
	if !ok {
		log.Fatalf("No token saved")
	}

	user := getGoogleUserData(tok)
	c.JSON(http.StatusOK, user)
}

func getGoogleUserData(token *oauth2.Token) *googleOauth2.Userinfo {
	config := utils.ReadGoogleAPICredentials()
	context := context.Background()

	srv, err := googleOauth2.NewService(
		context, 
		option.WithTokenSource(
			config.TokenSource(context, token)))
	if err != nil {
		log.Fatalf("Error creating service for oauth2")
	}

	user, err := srv.Userinfo.Get().Do()
	if err != nil {
		log.Fatalf("Error fetching user info %v", err)
	}

	return user
}	
