package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

var currentTime = time.Now

func GenerateToken(user_id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["expiry"] = currentTime().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret_api_key"))
}

func ExtractToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	tokenStrings := strings.Split(bearerToken, " ") 
	if len(tokenStrings) == 2 {
		return tokenStrings[1], nil
	}
	return "", fmt.Errorf("no token detected in request")
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, JwtKeyValidator)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenField(token *jwt.Token, fieldName string) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims[fieldName].(string), nil
	}
	return "", fmt.Errorf("invalid token")
}

func JwtKeyValidator(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("unexpected signing method on token: %v", token.Header["alg"])
	}
	return []byte("secret_api_key"), nil
}
