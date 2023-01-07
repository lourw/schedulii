package database

import (
	"net/http"
	models "schedulii/src/models"
	data "schedulii/src/services/data"

	"github.com/gin-gonic/gin"
)

func ReadUserHandler(env *models.Env) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        var u models.User
        err := c.ShouldBindQuery(&u)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        user, err := data.ReadUser(env, u)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        c.JSON(200, user)
    }
    return gin.HandlerFunc(fn)
}
