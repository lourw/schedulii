package services

import (
	"schedulii/src/models"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

var userService UserService
var mockUserRepository MockUserRepository
var dummyUser models.User

type MockUserRepository struct{}

var mockCreateUser func(user models.User) error
var mockReadUser func(user models.User) (models.User, error)
var mockUpdateUser func(user models.User) error

func (mock *MockUserRepository) Create(user models.User) error {
	return mockCreateUser(user)
}

func (mock *MockUserRepository) Read(user models.User) (models.User, error) {
	return mockReadUser(user)
}

func (mock *MockUserRepository) Update(user models.User) error {
	return mockUpdateUser(user)
}

func initUserServiceTest() {
	mockUserRepository = MockUserRepository{}
	userService = NewUserService(&mockUserRepository)

	dummyUser = models.User{
		Username: "User",
	}
}

func TestUserService(t *testing.T) {
	initUserServiceTest()

	t.Run(
		"Create user returns correctly",
		func(t *testing.T) {
			mockCreateUser = func(user models.User) error {
				return nil
			}

			err := userService.CreateUser(dummyUser)
			assert.Nil(t, err)
		},
	)

	t.Run(
		"Create user throws an error",
		func(t *testing.T) {
			mockCreateUser = func(user models.User) error {
				return pgx.ErrNoRows
			}

			err := userService.CreateUser(dummyUser)
			assert.Error(t, err)
		},
	)

	t.Run(
		"Read user returns correctly",
		func(t *testing.T) {
			mockReadUser = func(user models.User) (models.User, error) {
				return user, nil
			}

			result, err := userService.ReadUser(dummyUser)
			assert.Nil(t, err)
			assert.Equal(t, result, dummyUser)
		},
	)

	t.Run(
		"Read user returns an error",
		func(t *testing.T) {
			mockReadUser = func(user models.User) (models.User, error) {
				return user, pgx.ErrNoRows
			}

			_, err := userService.ReadUser(dummyUser)
			assert.Error(t, err)
		},
	)
}
