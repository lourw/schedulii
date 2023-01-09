package data_srv

import (
	"context"

	"schedulii/src/models"
	"schedulii/src/models/data_model"
)

func ReadEvent(env *models.Env, event data_model.Event) (data_model.Event, error) {
	query := "SELECT * FROM Events WHERE eventid = ($1)"
	queryResult := env.DB.QueryRow(
		context.Background(),
		query,
		event.EventId,
	)
	err := queryResult.Scan(
		&event.EventId,
		&event.GroupId,
		&event.EventName,
		&event.StartTime,
		&event.EndTime,
	)
	if err != nil {
		return event, err
	}
	return event, nil
}
