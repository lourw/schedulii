package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabaseConnection() (*pgxpool.Pool, error) {
	connectionString := retrieveDatabaseURL("DATABASE_URL")
	db, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return db, nil
}

func retrieveDatabaseURL(key string) string {
	url, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("There is no database connection string")
	}
	return url
}
