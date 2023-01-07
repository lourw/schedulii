package data

import (
	"context"

	models "schedulii/src/models"
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

func UpdateGroup(env *models.Env, g models.Groups) error {
	query := `
		UPDATE Groups
		SET GroupName = ($2),
			GroupURL = ($3),
			AvailableStartHour = ($4),
			AvailableEndHour = ($5)
		WHERE GroupID = ($1)
	`
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

func ReadGroup(env *models.Env, groupID int) (*models.Groups, error) {
	var g models.Groups
    query := "SELECT * FROM Groups WHERE GroupID = ($1)"
    queryResult := env.DB.QueryRow(context.Background(),
		query,
		groupID,
	)
	err := queryResult.Scan(&g.GroupID,
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
