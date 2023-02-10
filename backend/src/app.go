package main

import (
	"encoding/gob"
	"schedulii/src/middleware"
	"schedulii/src/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/oauth2"
)

type ScheduliiApp struct {
	db *pgxpool.Pool
	ginEngine *gin.Engine
	router routes.Router
}

func NewScheduliiApp(db *pgxpool.Pool, ginEngine *gin.Engine, router routes.Router) ScheduliiApp {
	app := ScheduliiApp{
		db: db,
		ginEngine: ginEngine,
		router: router,
	}
	app.setup()
	return app
}

func (sa *ScheduliiApp) setup() {
	// needed for the Google Oauth process. Not sure where else to register this.
	gob.Register(oauth2.Token{})
	store := cookie.NewStore([]byte("secret"))
	sa.ginEngine.Use(sessions.Sessions("schedulii", store))
	sa.ginEngine.Use(gin.Logger())
	sa.ginEngine.Use(middleware.CORSMiddleware)
	sa.router.SetupRoutes()
}

func (sa *ScheduliiApp) Run() error {
	panic(sa.ginEngine.Run(":8080"))
}

func (sa *ScheduliiApp) Teardown() {
	sa.db.Close()
}
