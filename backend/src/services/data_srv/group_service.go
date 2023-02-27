package data_srv

import (
	"schedulii/src/models"
	"schedulii/src/models/data_model"
)

type GroupService struct {
	repository models.Repository[data_model.Group]
}

func NewGroupService(repository models.Repository[data_model.Group]) GroupService {
	return GroupService{
		repository: repository,
	}
}

func (gs *GroupService) CreateGroup(group data_model.Group) error {
	err := gs.repository.Create(group)
	if err != nil {
		return err
	}
	return nil
}

func (gs *GroupService) ReadGroup(group data_model.Group) (data_model.Group, error) {
	result, err := gs.repository.Read(group)
	if err != nil {
		return group, err
	}
	return result, nil
}

func (gs *GroupService) UpdateGroup(group data_model.Group) error {
	err := gs.repository.Update(group)
	if err != nil {
		return err
	}
	return nil
}
