package services

import (
	"schedulii/src/models"
)

type UserService struct {
	repository models.Repository[models.User]
}

func NewUserService(repository models.Repository[models.User]) UserService {
	return UserService{
		repository: repository,
	}
}

func (us *UserService) CreateUser(user models.User) error {
	err := us.repository.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (us *UserService) ReadUser(user models.User) (models.User, error) {
	result, err := us.repository.Read(user)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (us *UserService) UpdateUser(user models.User) error {
	return nil
}
