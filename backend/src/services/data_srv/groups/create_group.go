package data_srv

import (
	"context"
	
	"schedulii/src/models"
	"schedulii/src/models/data_model"
)

func CreateGroup(env *models.Env, g data_model.Groups) error {
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
