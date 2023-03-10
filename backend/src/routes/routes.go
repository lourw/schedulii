package routes

import (
	"schedulii/src/handlers"
	"schedulii/src/handlers/google"
	"schedulii/src/middleware"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type Router struct {
	ginEngine    *gin.Engine
	userHandler  handlers.UserHandler
	eventHandler handlers.EventHandler
	groupHandler handlers.GroupHandler
}

func NewRouter(
	ginEngine *gin.Engine,
	userHandler handlers.UserHandler,
	eventHandler handlers.EventHandler,
	groupHandler handlers.GroupHandler,
) Router {
	return Router{
		ginEngine:    ginEngine,
		userHandler:  userHandler,
		eventHandler: eventHandler,
		groupHandler: groupHandler,
	}
}

func (r *Router) SetupRoutes() {
	// Serve frontend build
	r.ginEngine.Use(static.Serve("/", static.LocalFile("../../frontend/build", true)))

	r.ginEngine.GET("/health", handlers.HealthCheck)

	r.ginEngine.GET("/googleAuth", google.GoogleOauthLoginHandler)
	r.ginEngine.GET("/googleCallback", google.GoogleCallbackHandler)
	googleAuth := r.ginEngine.Group("/google")
	googleAuth.Use(middleware.CheckGoogleAuthenticated)
	{
		googleAuth.GET("/calendars", google.UserCalendarListHandler)
		googleAuth.GET("/events", google.UserCalendarEventsHandler)
		googleAuth.GET("/userInfo", google.UserInfoHandler)
	}

	login := r.ginEngine.Group("/login")
	{
		login.POST("", handlers.Login)
		login.POST("/validate", handlers.Validate)
	}

	data := r.ginEngine.Group("/data")
	// data.Use(middleware.CheckAuthenticated)
	{
		data.GET("/readUser", r.userHandler.HandleReadUser())
		data.GET("/readGroup", r.groupHandler.HandleReadGroup())
		data.GET("/readEvent", r.eventHandler.HandleReadEvent())
	}
}
