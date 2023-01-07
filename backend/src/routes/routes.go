package routes

import (
	handlers 	"schedulii/src/handlers"
	models 		"schedulii/src/models"
	google 		"schedulii/src/handlers/google"
	database	"schedulii/src/handlers/database"
	"schedulii/src/middleware"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine, env *models.Env) {
	// Serve frontend build
	engine.Use(static.Serve("/", static.LocalFile("../../frontend/build", true)))

	engine.GET("/health", handlers.HealthCheck)

	engine.GET("/googleAuth", google.GoogleOauthLoginHandler)
	engine.GET("/googleCallback", google.GoogleCallbackHandler)
	googleAuth := engine.Group("/google")
	googleAuth.Use(middleware.CheckGoogleAuthenticated)
	{
		googleAuth.GET("/calendars", google.UserCalendarListHandler)
		googleAuth.GET("/events", google.UserCalendarEventsHandler)
		googleAuth.GET("/userInfo", google.UserInfoHandler)
	}

	login := engine.Group("/login")
	{
		login.POST("/register", handlers.RegisterUser)
		login.POST("", handlers.Login)
		login.POST("/validate", handlers.Validate)
	}

	data := engine.Group("/data")
	// data.Use(middleware.CheckAuthenticated)
	{
		data.GET("/readUser", database.ReadUser(env))
	}
}
