package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"schedulii/src/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user models.User) error {
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

func (ur *UserRepository) Read(user models.User) (models.User, error) {
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

func (ur *UserRepository) Update(user models.User) error {
	return nil
}
