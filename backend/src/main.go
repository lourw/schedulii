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
	"github.com/joho/godotenv"
)

func main() {
	ginEngine := setUpEngine()

	// needed for the Google Oauth process. Not sure where else to register this.
	gob.Register(oauth2.Token{})

	err := ginEngine.Run(":8080")
	if err != nil {
		log.Fatal("Unable to start:", err)
	}
}

func setUpEngine() *gin.Engine {
	r := gin.Default()
	connectionString := retrieveURL("DATABASE_URL")
	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to database!")
	// defer db.Close()
	env := &models.Env{DB: db}
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("schedulii", store))
	r.Use(gin.Logger())
	r.Use(middleware.CORSMiddleware)
	router.SetupRoutes(r, env)
	return r
}

func retrieveURL(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}
	return os.Getenv(key)
}
