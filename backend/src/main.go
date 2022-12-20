package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func serveFrontend(router *gin.Engine) {
	router.Use(static.Serve("/", static.LocalFile("../../frontend/build", true)))
}	

func setUpRouter(router *gin.Engine) {
	router.GET("/health", func (c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
}

func main() {
	r := gin.Default()
	setUpRouter(r)
	serveFrontend(r)
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server")
	}
}
