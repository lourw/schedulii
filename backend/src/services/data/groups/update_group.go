package groups
 
import (
	"context"

	"schedulii/src/models"
)

func UpdateGroup(env *models.Env, g models.Groups) error {
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
