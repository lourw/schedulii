package routes

import (
	handlers "schedulii/src/handlers"
	data_handler "schedulii/src/handlers/data_handler"
	google "schedulii/src/handlers/google"
	"schedulii/src/middleware"
	models "schedulii/src/models"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Router struct {
	eventHandler data_handler.EventHandler
	groupHandler data_handler.GroupHandler
}

func (r *Router) SetupRoutes(engine *gin.Engine, env *models.Env) {
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
		login.POST("", handlers.Login)
		login.POST("/validate", handlers.Validate)
	}

	data := engine.Group("/data")
	// data.Use(middleware.CheckAuthenticated)
	{
		data.GET("/readUser", data_handler.ReadUserHandler(env))
		data.GET("/readGroup", r.groupHandler.HandleReadGroup())
		data.GET("/readEvent", r.eventHandler.HandleReadEvent())
	}
}
