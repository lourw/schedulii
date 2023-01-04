package middleware

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

var dummyRequest *http.Request
var dummyJWT string
var dummyOauthTok *oauth2.Token

func setup() {
	dummyRequest, _ = http.NewRequest("POST", "/test", nil)
	dummyJWT = "Token"

	dummyOauthTok = &oauth2.Token{
		AccessToken: "Allowed",
		Expiry: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
	}
}

func TestUserIsNotAuthorizedWithJWT(t *testing.T) {
	dummyRequest.Header = http.Header{
		"Authorization": {"Bearer" + "Test"},
	}

	ok := isUserAuthenticated(dummyRequest)
	assert.False(t, ok)
} 

func TestUserIsAuthorizedWithJWT(t *testing.T) {
	dummyRequest.Header = http.Header{
		"Authorization": {"Bearer " + dummyJWT},
	}

	ok := isUserAuthenticated(dummyRequest)
	assert.True(t, ok)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
