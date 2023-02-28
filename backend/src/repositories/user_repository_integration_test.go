//go:build integration
// +build integration

package repositories

import (
	"context"
	"log"
	"schedulii/src/db"
	"schedulii/src/models"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

var userDb *pgxpool.Pool
var userRepository UserRepository
var userCtx context.Context

var userTestData []models.User

func initUserTest() {
	var err error

	userDb, err = db.NewDatabaseConnection()
	if err != nil {
		log.Fatalf("Error initializing integration test: %s", err)
	}

	userRepository = UserRepository{
		db: userDb,
	}
	userCtx = context.Background()

	userTestData = []models.User{
		models.User{
			Username: "User@gmail.com",
		},
		models.User{
			Username: "User2@gmail.com",
		},
	}
}

func afterEachUserTest() {
	_, err := userDb.Exec(userCtx, "DELETE FROM users")
	if err != nil {
		log.Fatalf("Error on users table teardown: %s", err)
	}
}

func TestUserRepository(t *testing.T) {
	initUserTest()

	t.Run(
		"Create user returns correctly",
		func(t *testing.T) {
			var rowCount int
			var err error

			userDb.QueryRow(context.Background(), "SELECT COUNT(*) FROM users").Scan(&rowCount)
			assert.Equal(t, 0, rowCount)

			for i, data := range userTestData {
				err = userRepository.Create(data)
				if err != nil {
					log.Fatalf("Error: %s", err)
					t.Fail()
				}

				userDb.QueryRow(context.Background(), "SELECT COUNT(*) FROM users").Scan(&rowCount)
				assert.Equal(t, i+1, rowCount)
			}

			afterEachUserTest()
		},
	)

	t.Run(
		"Read user returns correctly",
		func(t *testing.T) {
			for _, data := range userTestData {
				err := userRepository.Create(data)
				if err != nil {
					log.Fatalf("Error: %s", err)
					t.Fail()
				}

				user, readErr := userRepository.Read(data)
				if readErr != nil {
					t.Fail()
				}
				assert.Equal(t, data, user)
			}

			afterEachUserTest()
		},
	)

	t.Cleanup(func() {
		afterEachUserTest()
		userDb.Close()
	})
}
