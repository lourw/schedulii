package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func ReadGoogleAPICredentials() *oauth2.Config {
	credentials, ok := os.LookupEnv("GOOGLE_APP_CREDENTIALS")
	if !ok {
		log.Fatalf("unable to get credentials")
	}

	config, err := google.ConfigFromJSON([]byte(credentials),
		"https://www.googleapis.com/auth/calendar.readonly",
		"https://www.googleapis.com/auth/userinfo.profile",
		"openid email profile",
	)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return config
}

func OauthTokenFromSession(s sessions.Session) (*oauth2.Token, bool) {
	tok, ok := s.Get("token").(oauth2.Token)
	if !ok {
		return nil, false
	}

	return &tok, tok.Valid()
}

func SaveOauthTokenToSession(s sessions.Session, token *oauth2.Token) {
	s.Set("token", *token)
	err := s.Save()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func GetGoogleClient(c *gin.Context) (*http.Client, bool) {
	config := ReadGoogleAPICredentials()
	session := sessions.Default(c)
	tok, valid := OauthTokenFromSession(session)
	if !valid {
		return nil, false
	}

	return config.Client(c, tok), true
}
