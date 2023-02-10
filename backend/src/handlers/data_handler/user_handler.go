package data_handler

import (
	"net/http"
	"schedulii/src/models"
	"schedulii/src/models/data_model"
	"schedulii/src/services/data_srv/users"

	"github.com/gin-gonic/gin"
)

func ReadUserHandler(env *models.Env) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        var u data_model.User
        err := c.ShouldBindQuery(&u)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        user, err := data_srv.ReadUser(env, u)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        c.JSON(200, user)
    }
    return gin.HandlerFunc(fn)
}
