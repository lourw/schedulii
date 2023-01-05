package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"context"
	"os"
	"schedulii/src/middleware"
	router "schedulii/src/routes"
	models "schedulii/src/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"github.com/jackc/pgx/v5/pgxpool"
)

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
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("schedulii", store))
	r.Use(gin.Logger())
	r.Use(middleware.CORSMiddleware)
	router.SetupRoutes(r, env)
	return r
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
	fmt.Println("Successfully connected to database!")

	return db
}
