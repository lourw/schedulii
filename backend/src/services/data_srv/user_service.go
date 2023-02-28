package data_srv

import (
	"schedulii/src/models"
	"schedulii/src/models/data_model"
)

type UserService struct {
	repository models.Repository[data_model.User]
}

func NewUserService(repository models.Repository[data_model.User]) UserService {
	return UserService{
		repository: repository,
	}
}

func (us *UserService) CreateUser(user data_model.User) error {
	err := us.repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) ReadUser(user data_model.User) (data_model.User, error) {
	result, err := us.repository.Read(user)
    if err != nil {
        return result, err
    }
    return result, nil
}


func (us *UserService) UpdateUser(user data_model.User) error {
	return nil;
}
