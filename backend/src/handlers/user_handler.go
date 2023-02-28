package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"schedulii/src/models"
	"schedulii/src/services"
)

type UserHandler struct {
	us services.UserService
}

func NewUserHandler(us services.UserService) UserHandler {
	return UserHandler{
		us: us,
	}
}

func (uh *UserHandler) HandleReadUser() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var u models.User
		err := c.ShouldBindQuery(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := uh.us.ReadUser(u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(200, user)
	}
	return gin.HandlerFunc(fn)
}
