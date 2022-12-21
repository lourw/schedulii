package routes

import (
	controllers "schedulii/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func SetupRoutes(router *gin.Engine) {
	// Serve frontend build
	router.Use(static.Serve("/", static.LocalFile("../../frontend/build", true)))

	// GET Routes
	router.GET("/health", controllers.HealthCheckController)
	router.GET("/google", controllers.RunGoogleConnection)
	router.GET("/auth", controllers.RunGoogleCallback)
}
