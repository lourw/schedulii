package data

import (
	"context"
	"fmt"
	"log"

	models "schedulii/src/models"
)

func CreateGroup(env *models.Env, g models.Groups) {
	query := "INSERT INTO Groups VALUES ($1, $2, $3, $4, $5)"
	_, err := env.DB.Exec(context.Background(), query, g.GroupID, g.GroupName, g.GroupURL, g.AvailableStartHour, g.AvailableEndHour)
	if err != nil {
		log.Fatalf("Unable to insert value: %v", err)
	}
	fmt.Println("\nRow inserted successfully!")
}

func UpdateGroup(env *models.Env, g models.Groups) {
	query := "UPDATE Groups SET GroupName = ($2), GroupURL = ($3), AvailableStartHour = ($4), AvailableEndHour = ($5) WHERE GroupID = ($1)"
	_, err := env.DB.Exec(context.Background(), query, g.GroupID, g.GroupName, g.GroupURL, g.AvailableStartHour, g.AvailableEndHour)
	if err != nil {
		log.Fatalf("Unable to insert value: %v", err)
	}
	fmt.Println("\nRow updated successfully!")
}

// func ReadGroup(env *models.Env, g models.Groups) models.Groups {
// 	query := "SELECT * FROM Groups WHERE GroupID = ($1)"
// 	err := env.DB.QueryRow(context.Background(), query, g.GroupID).Scan(&g.GroupID, &g.GroupName, &g.GroupURL, &g.AvailableStartHour, &g.AvailableEndHour)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve user info: %v", err)
// 	}
// 	return g
// }
