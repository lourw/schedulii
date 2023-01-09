package data_srv

import (
	"context"

	"schedulii/src/models"
	"schedulii/src/models/data_model"
)

func ReadUser(env *models.Env, user data_model.User) (*data_model.User, error) {
	var u data_model.User
    query := "SELECT * FROM Users WHERE UserEmail = ($1)"
    err := env.DB.QueryRow(context.Background(), query, user.Username).Scan(&u.Username)
    if err != nil {
        return nil, err
    }
    return &u, nil
}
