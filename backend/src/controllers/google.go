package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	"github.com/gin-gonic/gin"
)

func RunGoogleConnection(c *gin.Context) {
	ctx := context.Background()
	config := readGoogleAPICredentials()

	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline, oauth2.ApprovalForce)
		c.Redirect(http.StatusFound, authURL)
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

	c.String(http.StatusOK, "%v", cal)
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

	saveToken("token.json", tok)
	c.Redirect(http.StatusFound, "/google")
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

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache OAuth token: %v", err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		log.Fatalf("Error with encoding")
	}
}
