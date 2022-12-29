package routes

import (
	handlers "schedulii/src/handlers"
	"schedulii/src/middleware"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	// Serve frontend build
	engine.Use(static.Serve("/", static.LocalFile("../../frontend/build", true)))

	engine.GET("/health", handlers.HealthCheck)

	googleAuth := engine.Group("/google")
	{
		googleAuth.GET("/googleAuth", handlers.RunGoogleConnection)
		googleAuth.GET("/googleCallback", handlers.RunGoogleCallback)
	}

	authorized := engine.Group("/authorized")
	authorized.Use(middleware.CheckAuthenticated) 
	{
		authorized.GET("/", handlers.GetCalendars)
	}

	login := engine.Group("/login")
	{
		login.POST("/register", handlers.RegisterUser)
		login.POST("/login", handlers.Login)
		login.POST("/validate", handlers.Validate)
	}
}
