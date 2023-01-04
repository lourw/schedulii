package utils

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

var dummyUserId string
var dummyNow time.Time
var dummyTokenSecret string
var dummyRequest *http.Request
var dummyJWT = "token"

func setup() {
	dummyUserId = "testUser"
	dummyTokenSecret = "token_secret"
	dummyNow = time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)

	currentTime = func() time.Time {
		return dummyNow
	}

	dummyRequest, _ = http.NewRequest("POST", "/test", nil)
}

func TestGenerateUser(t *testing.T) {
	generatedToken, _ := GenerateUserJWT(dummyUserId)
	token, err := jwt.Parse(generatedToken, JwtKeyValidator)
	if err != nil {
		t.Fail()
	}
	generatedClaims := token.Claims.(jwt.MapClaims)

	expectedExpiryTime := dummyNow.Add(time.Hour * 24).Unix()

	assert.True(t, generatedClaims["authorized"].(bool))
	assert.Equal(t, generatedClaims["user_id"].(string), dummyUserId)
	assert.Equal(t, int64(generatedClaims["expiry"].(float64)), expectedExpiryTime)
}

func TestValidJWTInRequest(t *testing.T) {
	dummyRequest.Header = http.Header{
		"Authorization": {"Bearer " + dummyJWT},
	}

	jwt, ok := ExtractUserJWT(dummyRequest)
	if !ok {
		t.Fail()
	}

	assert.Equal(t, "token", jwt)
}

func TestInvalidJWTInRequest(t *testing.T) {
	dummyRequest.Header = http.Header{
		"Authorization": {"Bearer"},
	}

	_, ok := ExtractUserJWT(dummyRequest)
	if !ok {
		return
	}
	t.Fail()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
