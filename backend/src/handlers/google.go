package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func RunGoogleConnection(c *gin.Context) {
	ctx := context.Background()
	config := readGoogleAPICredentials()

	tok, valid := tokenFromSession(c)
	if !valid {
		authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
		c.Redirect(http.StatusFound, authURL)
		c.Abort()
		return
	}
	
	client := config.Client(context.Background(), tok)
	
	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve calendar client: %v", err)
	}
	
	cal, err := srv.CalendarList.List().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from calendars: %v", err)
	}

	c.JSON(200, cal.Items)
}

func RunGoogleCallback(c *gin.Context) {
	status := c.Writer.Status()
	if status != http.StatusOK {
		fmt.Print("Status not 200 when getting auth callback from Google API")
	}

	config := readGoogleAPICredentials()

	authCode := c.Request.URL.Query().Get("code")
	tok, err := config.Exchange(context.Background(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}

	saveTokenToSession(c, tok)
	c.Redirect(http.StatusFound, "/google/googleAuth")
}

func readGoogleAPICredentials() *oauth2.Config {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/calendar.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	return config
}

func tokenFromSession(c *gin.Context) (*oauth2.Token, bool) {
	session := sessions.Default(c)

	tok, ok := session.Get("token").(oauth2.Token)
	if !ok {
		return nil, false
	}

	return &tok, tok.Valid()
}

func saveTokenToSession(c *gin.Context, token *oauth2.Token) {
	session := sessions.Default(c)
	session.Set("token", *token)
	err := session.Save()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
