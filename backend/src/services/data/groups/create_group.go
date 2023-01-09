package groups

import (
	"context"
	
	"schedulii/src/models"
)

func CreateGroup(env *models.Env, g models.Groups) error {
	query := "INSERT INTO Groups VALUES ($1, $2, $3, $4, $5)"
	_, err := env.DB.Exec(context.Background(),
		query,
		g.GroupID,
		g.GroupName,
		g.GroupURL,
		g.AvailableStartHour,
		g.AvailableEndHour,
	)
	if err != nil {
		return err
	}
	return nil
}
