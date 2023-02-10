package data_srv

import (
	"context"
	"schedulii/src/models/data_model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) UserService {
	return UserService{
		db: db,
	}
}

func (us *UserService) CreateUser(user data_model.User) error {
	query := `
		INSERT INTO Users 
		VALUES ($1)
	`
	_, err := us.db.Exec(
		context.Background(),
		query, 
		user.Username,
	)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) ReadUser(user data_model.User) (*data_model.User, error) {
	var u data_model.User
    query := `
		SELECT * FROM Users 
		WHERE UserEmail = ($1)
	`
    result := us.db.QueryRow(
		context.Background(), 
		query, 
		user.Username,
	)
	err := result.Scan(
		&u.Username,
	)
    if err != nil {
        return nil, err
    }
    return &u, nil
}
