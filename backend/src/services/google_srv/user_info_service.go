package google_srv

import (
	"context"
	"log"

	"schedulii/src/utils"

	"golang.org/x/oauth2"
	googleOauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func GetGoogleUserData(token *oauth2.Token) *googleOauth2.Userinfo {
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
