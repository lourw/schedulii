package main

import (
	"fmt"
	router "schedulii/src/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	ginEngine.Use(sessions.Sessions("schedulii", store))
	router.SetupRoutes(ginEngine)

	err := ginEngine.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
	}
}
