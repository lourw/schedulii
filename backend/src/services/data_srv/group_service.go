package data_srv

import (
	"schedulii/src/models/data_model"
	"schedulii/src/repositories"
)

type GroupService struct {
	gr repositories.GroupRepository
}

func NewGroupService(gr repositories.GroupRepository) GroupService {
	return GroupService{
		gr: gr,
	}
}

func (gs *GroupService) CreateGroup(group data_model.Group) error {
	err := gs.gr.CreateGroup(group)
	if err != nil {
		return err
	}
	return nil
}

func (gs *GroupService) ReadGroup(group data_model.Group) (data_model.Group, error) {
	result, err := gs.gr.ReadGroup(group)
	if err != nil {
		return group, err
	}
	return result, nil
}

func (gs *GroupService) UpdateGroup(group data_model.Group) error {
	err := gs.gr.UpdateGroup(group)
	if err != nil {
		return err
	}
	return nil
}
