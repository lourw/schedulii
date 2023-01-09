package users

import (
	"context"

	"schedulii/src/models"
)

func CreateUser(env *models.Env, user models.User) error {
	query := "INSERT INTO Users VALUES ($1)"
	_, err := env.DB.Exec(context.Background(), query, user.Username)
	if err != nil {
		return err
	}
	return nil
}
