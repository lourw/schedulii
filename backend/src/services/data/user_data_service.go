package data

import (
	"context"

	models "schedulii/src/models"
)

func CreateUser(env *models.Env, user models.User) error {
	query := "INSERT INTO Users VALUES ($1)"
	_, err := env.DB.Exec(context.Background(), query, user.Username)
	if err != nil {
		return err
	}
	return nil
}

func ReadUser(env *models.Env, user models.User) (*models.User, error) {
	var u models.User
    query := "SELECT * FROM Users WHERE UserEmail = ($1)"
    err := env.DB.QueryRow(context.Background(), query, user.Username).Scan(&u.Username)
    if err != nil {
        return nil, err
    }
    return &u, nil
}
