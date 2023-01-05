package database

import (
	"context"
	"fmt"
	"log"
	"net/http"

	models "schedulii/src/models"

	"github.com/gin-gonic/gin"
)

func CreateEvent(env *models.Env) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		var user models.User

		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "validated"})
		query := "INSERT INTO Users VALUES ($1)"
		row, err := env.DB.Exec(context.Background(), query, user.Username)
		if err != nil {
			log.Fatalf("Unable to insert value: %v", err)
		}
		fmt.Println("\nRow inserted successfully!")
		c.JSON(http.StatusOK, row)
	}

	return gin.HandlerFunc(fn)
}
