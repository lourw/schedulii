package database

import (
	"fmt"
	"context"
	"net/http"
	models "schedulii/src/models"
	// data "schedulii/src/services/data"

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
        // data.ReadGroup(env, g)
		query := "SELECT * FROM Groups WHERE GroupID = ($1)"
		readError := env.DB.QueryRow(context.Background(), query, g.GroupID).Scan(&g.GroupID, &g.GroupName, &g.GroupURL, &g.AvailableStartHour, &g.AvailableEndHour)
		if readError != nil {
			fmt.Printf("Unable to retrieve user info: %v", err)
		}
        c.JSON(200, g)
    }
    return gin.HandlerFunc(fn)
}
