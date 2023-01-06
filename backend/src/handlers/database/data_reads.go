package database

import (
	"context"
	"net/http"
	models 	"schedulii/src/models"
	data	"schedulii/src/services/data"
	"github.com/gin-gonic/gin"
)


func ReadUser(env *models.Env) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		var user models.User

		c.ShouldBindQuery(&user)
		err := env.DB.QueryRow(context.Background(), data.SelectUser(), user.Username).Scan(&user.Username)
			if err != nil {
				c.JSON(http.StatusBadRequest, "Could not find user.")
				return
			}
		c.JSON(http.StatusOK, user)
	}
	return gin.HandlerFunc(fn)
}
