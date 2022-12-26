package main

import (
	"encoding/gob"
	"log"
	router "schedulii/src/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func main() {
	ginEngine := setUpEngine()

	// needed for the Google Oauth process. Not sure where else to register this.
	gob.Register(oauth2.Token{})

	err := ginEngine.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func setUpEngine() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("schedulii", store))
	r.Use(gin.Logger())
	router.SetupRoutes(r)
	return r
}
