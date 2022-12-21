package main

import (
	"fmt"
	router "schedulii/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()
	router.SetupRoutes(ginEngine)

	err := ginEngine.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
	}
}
