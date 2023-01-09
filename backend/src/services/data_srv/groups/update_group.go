package data_srv
 
import (
	"context"

	"schedulii/src/models"
	"schedulii/src/models/data_model"
)

func UpdateGroup(env *models.Env, g data_model.Groups) error {
	query := `
		UPDATE Groups
		SET GroupName = ($2),
			GroupURL = ($3),
			AvailableStartHour = ($4),
			AvailableEndHour = ($5)
		WHERE GroupID = ($1)
	`
	_, err := env.DB.Exec(
		context.Background(),
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
