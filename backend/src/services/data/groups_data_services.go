package data

import (
	"context"
	"fmt"
	"log"

	models "schedulii/src/models"
)

func CreateGroup(env *models.Env, group models.Groups) {
	query := "INSERT INTO Groups VALUES ($1, $2, $3, $4, $5)"
	_, err := env.DB.Exec(context.Background(), query, group.GroupID, group.GroupName, group.GroupURL, group.AvailabilityStart, group.AvailabilityEnd)
		if err != nil {
			log.Fatalf("Unable to insert value: %v", err)
		}
		fmt.Println("\nRow inserted successfully!")
}
