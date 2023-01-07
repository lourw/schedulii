package database

import (
	"net/http"
	models "schedulii/src/models"
	data "schedulii/src/services/data"

	"github.com/gin-gonic/gin"
)

func ReadUserHandler(env *models.Env) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        var user models.User
        err := c.ShouldBindQuery(&user)
            if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }
        data.ReadUser(env, user)
        c.JSON(200, user)
    }
    return gin.HandlerFunc(fn)
}
