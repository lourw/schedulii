package main

import "github.com/gin-gonic/gin"
import "net/http"

func setUpRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func (c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
	return r
}

func main() {
	r := setUpRouter()
	r.Run(":8080");
}
