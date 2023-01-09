package data_srv

import (
	"context"

	"schedulii/src/models"
	"schedulii/src/models/data_model"
)

func CreateUser(env *models.Env, user data_model.User) error {
	query := "INSERT INTO Users VALUES ($1)"
	_, err := env.DB.Exec(context.Background(), query, user.Username)
	if err != nil {
		return err
	}
	return nil
}
