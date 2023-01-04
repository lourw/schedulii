package routes

import (
	handlers "schedulii/src/handlers"
	google "schedulii/src/handlers/google"
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
		googleAuth.GET("/googleAuth", google.RunGoogleConnection)
		googleAuth.GET("/googleCallback", google.RunGoogleCallback)
		googleAuth.GET("/calendars", google.UserCalendarListHandler)
		googleAuth.GET("/events", google.UserCalendarEventsHandler)
	}

	authorized := engine.Group("/authorized")
	authorized.Use(middleware.CheckAuthenticated)

	login := engine.Group("/login")
	{
		login.POST("/register", handlers.RegisterUser)
		login.POST("", handlers.Login)
		login.POST("/validate", handlers.Validate)
	}
}
