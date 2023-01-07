package database

import (
	"net/http"
	models "schedulii/src/models"
	data "schedulii/src/services/data"

	"github.com/gin-gonic/gin"
)

func ReadGroupHandler(env *models.Env) gin.HandlerFunc {
    fn := func(c *gin.Context) {
        var g models.Groups
        err := c.ShouldBindQuery(&g)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        group, err := data.ReadGroup(env, g.GroupID)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        }
        c.JSON(200, group)
    }
    return gin.HandlerFunc(fn)
}
