package database

import (
	"context"
	"fmt"
	"log"
	"net/http"

	models "schedulii/src/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(env *models.Env) gin.HandlerFunc {

	fn := func(c *gin.Context)  {
		query := "INSERT INTO Users VALUES ($1)"
		row, err := env.DB.Exec(context.Background(), query, "testemail2@gmail.com")
			if err != nil {
				log.Fatalf("Unable to insert value: %v", err)
			}
			fmt.Println("\nRow inserted successfully!")
			c.JSON(http.StatusOK, row)
	}

	return gin.HandlerFunc(fn)
}
