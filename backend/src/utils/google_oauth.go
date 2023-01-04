package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func ReadGoogleAPICredentials() *oauth2.Config {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("unable to get the current filename")
	}
	dirname := filepath.Dir(filename)

	b, err :=  os.ReadFile(dirname + "/credentials.json")
	if err != nil {
		log.Fatalf("unable to get credentials file at %s/credentials.json", dirname)
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/calendar.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return config
}

func TokenFromSession(s sessions.Session) (*oauth2.Token, bool) {
	tok, ok := s.Get("token").(oauth2.Token)
	if !ok {
		return nil, false
	}

	return &tok, tok.Valid()
}

func SaveTokenToSession(s sessions.Session, token *oauth2.Token) {
	s.Set("token", *token)
	err := s.Save()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func GetGoogleClient(c *gin.Context) (*http.Client, bool) {
	config := ReadGoogleAPICredentials()
	session := sessions.Default(c)
	tok, valid := TokenFromSession(session)
	if !valid {
		return nil, false
	}

	return config.Client(c, tok), true
}
