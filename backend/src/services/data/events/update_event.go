package events

import (
	"context"

	"schedulii/src/models"
)

func UpdateEvent(env *models.Env, event models.Event) error {
	query := `
		UPDATE Events
			SET EventName = ($2),
				StartTime = ($3),
				EndTime = ($4)
			WHERE EventID = ($1)
	`
	_, err := env.DB.Exec(
		context.Background(),
		query,
		event.EventId,
		event.EventName,
		event.StartTime,
		event.EndTime,
	)
	if err != nil {
		return err
	}
	return nil
}
