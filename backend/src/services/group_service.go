package services

import (
	"schedulii/src/models"
)

type GroupService struct {
	repository models.Repository[models.Group]
}

func NewGroupService(repository models.Repository[models.Group]) GroupService {
	return GroupService{
		repository: repository,
	}
}

func (gs *GroupService) CreateGroup(group models.Group) error {
	err := gs.repository.Create(group)
	if err != nil {
		return err
	}
	return nil
}

func (gs *GroupService) ReadGroup(group models.Group) (models.Group, error) {
	result, err := gs.repository.Read(group)
	if err != nil {
		return group, err
	}
	return result, nil
}

func (gs *GroupService) UpdateGroup(group models.Group) error {
	err := gs.repository.Update(group)
	if err != nil {
		return err
	}
	return nil
}
