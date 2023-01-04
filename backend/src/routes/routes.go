package routes

import (
	handlers "schedulii/src/handlers"
	google "schedulii/src/handlers/google"
	"schedulii/src/middleware"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct { // undeclared name error if struct isn't here
	db	*pgxpool.Pool
}

func SetupRoutes(engine *gin.Engine, env *Env) {
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
		googleAuth.GET("/user", google.GoogleUserDataHandler)
	}

	authorized := engine.Group("/authorized")
	authorized.Use(middleware.CheckAuthenticated)

	login := engine.Group("/login")
	{
		login.POST("/register", handlers.RegisterUser)
		login.POST("", handlers.Login)
		login.POST("/validate", handlers.Validate)
	}

	data := engine.Group("/data")
	{
		data.GET("", handlers.DBConnect)
		data.GET("/createUser", handlers.CreateUser(env))
	}
}
