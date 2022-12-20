package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func (c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	return r
}

func main() {
	r := setUpRouter()
	err := r.Run(":8080");
	if err != nil {
		fmt.Println("Error starting server")
	}
}
