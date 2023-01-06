package database

import (
	"context"
	"net/http"
	models "schedulii/src/models"

	"github.com/gin-gonic/gin"
)


func ReadUser(env *models.Env) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		var user models.User
		err := c.ShouldBindQuery(&user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error1": err.Error()})
				return
			}
		query := "SELECT * FROM UserEmail WHERE UserEmail = ($1)"
		scanErr := env.DB.QueryRow(context.Background(), query, user.Username).Scan(&user.Username)
			if scanErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error2": err.Error()})
				return
			}
		c.JSON(http.StatusOK, user)
	}
	return gin.HandlerFunc(fn)
}
