package database

import (
	"context"
	"fmt"

	models "schedulii/src/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(env *models.Env) gin.HandlerFunc {

	fn := func(c *gin.Context)  {
		query := "INSERT INTO Users VALUES ('gogopher@gmail.com')";
		row := env.DB.QueryRow(context.Background(), query)
		fmt.Println(row)
		c.String(200, "Query successful")
	}

	return gin.HandlerFunc(fn)
}
