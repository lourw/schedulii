package handlers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func retrieveURL(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}
	return os.Getenv(key)
  }

func DBConnect(c *gin.Context) {
	// Load the env file to access database url for connection
	connectionString := retrieveURL("DATABASE_URL")

	// Connect to database
	dbpool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully connected to database!")

	defer dbpool.Close()
}

func (env *Env) createUser(email string) (Users, error) {
	query := 'INSERT INTO Users VALUES ("gogopher@gmail.com")';

	// result := env.Pool.QueryRow(context.Background(), query, UserEmail)
}
