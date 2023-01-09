package groups

import (
	"context"

	"schedulii/src/models"
)

func ReadGroup(env *models.Env, groupID int) (*models.Groups, error) {
	var g models.Groups
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
