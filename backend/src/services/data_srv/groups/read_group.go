package data_srv

import (
	"context"

	"schedulii/src/models"
	"schedulii/src/models/data_model"
)

func ReadGroup(env *models.Env, groupID int) (*data_model.Groups, error) {
	var g data_model.Groups
    query := "SELECT * FROM Groups WHERE GroupID = ($1)"
    queryResult := env.DB.QueryRow(
		context.Background(),
		query,
		groupID,
	)
	err := queryResult.Scan(
		&g.GroupID,
		&g.GroupName,
		&g.GroupURL,
		&g.AvailableStartHour,
		&g.AvailableEndHour,
	)
    if err != nil {
        return nil, err
    }
    return &g, nil
}
