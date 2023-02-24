package repositories

import (
	"context"
	"schedulii/src/models/data_model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user data_model.User) error {
	query := `
		INSERT INTO users 
		VALUES ($1)
	`
	_, err := ur.db.Exec(
		context.Background(),
		query, 
		user.Username,
	)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) ReadUser(user data_model.User) (data_model.User, error) {
    query := `
		SELECT * FROM users 
		WHERE user_email = ($1)
	`
    result := ur.db.QueryRow(
		context.Background(), 
		query, 
		user.Username,
	)
	err := result.Scan(
		&user.Username,
	)
    if err != nil {
        return user, err
    }
    return user, nil
}

