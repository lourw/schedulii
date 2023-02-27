package data_srv

import (
	"schedulii/src/models/data_model"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

var userService UserService
var mockUserRepository MockUserRepository
var dummyUser data_model.User

type MockUserRepository struct{}

var mockCreateUser func(user data_model.User) error
var mockReadUser func(user data_model.User) (data_model.User, error)
var mockUpdateUser func(user data_model.User) error

func (mock *MockUserRepository) Create(user data_model.User) error {
	return mockCreateUser(user)
}

func (mock *MockUserRepository) Read(user data_model.User) (data_model.User, error) {
	return mockReadUser(user)
}

func (mock *MockUserRepository) Update(user data_model.User) error {
	return mockUpdateUser(user)
}

func initUserServiceTest() {
	mockUserRepository = MockUserRepository{}
	userService = NewUserService(&mockUserRepository)

	dummyUser = data_model.User{
		Username: "User",
	}
}

func TestUserService(t *testing.T) {
	initUserServiceTest()

	t.Run(
		"Create user returns correctly",
		func(t *testing.T) {
			mockCreateUser = func(user data_model.User) error {
				return nil
			}

			err := userService.CreateUser(dummyUser)
			assert.Nil(t, err)
		},
	)

	t.Run(
		"Create user throws an error",
		func(t *testing.T) {
			mockCreateUser = func(user data_model.User) error {
				return pgx.ErrNoRows
			}

			err := userService.CreateUser(dummyUser)
			assert.Error(t, err)
		},
	)

	t.Run(
		"Read user returns correctly",
		func(t *testing.T) {
			mockReadUser = func(user data_model.User) (data_model.User, error) {
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
			mockReadUser = func(user data_model.User) (data_model.User, error) {
				return user, pgx.ErrNoRows
			}

			_, err := userService.ReadUser(dummyUser)
			assert.Error(t, err)
		},
	)
}
