package events

import (
	"context"

	models "schedulii/src/models"
)

func ReadEvent(env *models.Env, event models.Event) (models.Event, error) {
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
