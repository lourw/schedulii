package data

import (
	"context"
	"fmt"
	"log"

	models "schedulii/src/models"
)

// Helper function to add a user to the database
func CreateUser(env *models.Env, user models.User) {
	query := "INSERT INTO Users VALUES ($1)"
	_, err := env.DB.Exec(context.Background(), query, user.Username)
	if err != nil {
		log.Fatalf("Unable to insert value: %v", err)
	}
	fmt.Println("\nRow inserted successfully!")
}

func ReadUser(env *models.Env, user models.User) models.User {
    query := "SELECT * FROM Users WHERE UserEmail = ($1)"
    err := env.DB.QueryRow(context.Background(), query, user.Username).Scan(&user.Username)
        if err != nil {
            log.Fatalf("Unable to retrieve user info: %v", err)
        }
    return user
}
