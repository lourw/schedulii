package data_srv

import (
	"context"

	"schedulii/src/models"
	"schedulii/src/models/data_model"
)

func CreateEvent(env *models.Env, event data_model.Event) error {
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
