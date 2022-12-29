package utils

import (
	// "fmt"
	// "os"
	// "strconv"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user_id string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["expiry"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret_api_key"))
}

func ValidateToken(c *gin.Context) error {
	tokenString, err := ExtractToken(c)
	if err != nil {
		return err
	}

	_, err = jwt.Parse(tokenString, jwtKeyValidator)
	if err != nil {
		return err
	}
	return nil
}

func ExtractTokenId(c *gin.Context) (string, error) {
	tokenString, err := ExtractToken(c)
	if err != nil {
		return "", err
	}

	token, err := jwt.Parse(tokenString, jwtKeyValidator)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["user_id"].(string), nil
	}
	return "", nil
}

func ExtractToken(c *gin.Context) (string, error) {
	bearerToken := c.Request.Header.Get("Authorization")
	tokenStrings := strings.Split(bearerToken, " ") 
	if len(tokenStrings) == 2 {
		return tokenStrings[1], nil
	}
	return "", errors.New("no token detected in request")
}


func jwtKeyValidator(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("unexpected signing method on token: %v", token.Header["alg"])
	}
	return []byte("secret_api_key"), nil
}
