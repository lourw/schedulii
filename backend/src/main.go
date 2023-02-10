package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"schedulii/src/middleware"
	"schedulii/src/models"
	"schedulii/src/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/oauth2"
)

type ScheduliiApp struct {
	ginEngine *gin.Engine
	router routes.Router
}

func NewScheduliiApp(ginEngine *gin.Engine, router routes.Router) ScheduliiApp {
	return ScheduliiApp{
		ginEngine: ginEngine,
		router: router,
	}
}

func main() {
	env := &models.Env{DB: setupDatabaseConnection()}
	ginEngine := setUpEngine(env)

	// needed for the Google Oauth process. Not sure where else to register this.
	gob.Register(oauth2.Token{})


	err := ginEngine.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start:", err)
	}
	defer env.DB.Close()
}

func setUpEngine(env *models.Env) *gin.Engine {
	e := gin.Default()
	r := routes.Router{}

	store := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("schedulii", store))
	e.Use(gin.Logger())
	e.Use(middleware.CORSMiddleware)
	r.SetupRoutes(e, env)
	return e
}

func retrieveURL(key string) string {
	url, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("There is no database connection string")
	}
	return url
}

func setupDatabaseConnection() *pgxpool.Pool {
	connectionString := retrieveURL("DATABASE_URL")
	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db
}

