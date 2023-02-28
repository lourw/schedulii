package handlers

import (
	"net/http"
	"schedulii/src/models"
	"schedulii/src/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateUserJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Validate(c *gin.Context) {
	tokenString, ok := utils.ExtractUserJWT(c.Request)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JWT"})
	}

	token, err := utils.DecryptJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userId, err := utils.ExtractJWTField(token, "user_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"userId": userId})
}
