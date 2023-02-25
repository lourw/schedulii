package data_srv

import (
	"schedulii/src/models/data_model"
	"schedulii/src/repositories"
)

type UserService struct {
	ur repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) UserService {
	return UserService{
		ur: ur,
	}
}

func (us *UserService) CreateUser(user data_model.User) error {
	err := us.ur.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) ReadUser(user data_model.User) (data_model.User, error) {
	result, err := us.ur.ReadUser(user)
    if err != nil {
        return result, err
    }
    return result, nil
}
