package users

import (
	"context"

	"schedulii/src/models"
)

func ReadUser(env *models.Env, user models.User) (*models.User, error) {
	var u models.User
    query := "SELECT * FROM Users WHERE UserEmail = ($1)"
    err := env.DB.QueryRow(context.Background(), query, user.Username).Scan(&u.Username)
    if err != nil {
        return nil, err
    }
    return &u, nil
}
