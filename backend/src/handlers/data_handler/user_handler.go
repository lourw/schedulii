package data_handler

import (
	"net/http"
	"schedulii/src/models/data_model"
	"schedulii/src/services/data_srv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	us data_srv.UserService
}

func NewUserHandler(us data_srv.UserService) UserHandler {
	return UserHandler{
		us: us,
	}
}

func (uh *UserHandler) HandleReadUser() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var u data_model.User
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
