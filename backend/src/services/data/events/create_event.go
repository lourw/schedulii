package events

import (
	"context"

	models "schedulii/src/models"
)

func CreateEvent(env *models.Env, event models.Event) error {
	query := "INSERT INTO Events VALUES ($1, $2, $3, $4, $5)"
	_, err := env.DB.Exec(
		context.Background(),
		query,
		event.EventId,
		event.GroupId,
		event.EventName,
		event.StartTime,
		event.EndTime,
	)
	if err != nil {
		return err
	}
	return nil
}
