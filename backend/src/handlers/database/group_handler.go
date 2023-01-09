package database

import (
	"net/http"
	"schedulii/src/models"
    "schedulii/src/models/data_model"
	"schedulii/src/services/data_srv/groups"

	"github.com/gin-gonic/gin"
)

func ReadGroupHandler(env *models.Env) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        var g data_model.Groups
        err := c.ShouldBindQuery(&g)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        group, err := data_srv.ReadGroup(env, g.GroupID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        c.JSON(200, group)
    }
    return gin.HandlerFunc(fn)
}
