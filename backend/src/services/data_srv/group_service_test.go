package data_srv

import (
	"schedulii/src/models/data_model"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

var groupService GroupService
var mockGroupRepositority MockGroupRepository
var dummyGroup data_model.Group

type MockGroupRepository struct{}

var mockCreateGroup func(group data_model.Group) error
var mockReadGroup func(group data_model.Group) (data_model.Group, error)
var mockUpdateGroup func(group data_model.Group) error

func (mock *MockGroupRepository) Create(group data_model.Group) error {
	return mockCreateGroup(group)
}

func (mock *MockGroupRepository) Read(group data_model.Group) (data_model.Group, error) {
	return mockReadGroup(group)
}

func (mock *MockGroupRepository) Update(group data_model.Group) error {
	return mockUpdateGroup(group)
}

func initGroupServiceTest() {
	mockGroupRepositority = MockGroupRepository{}
	groupService = NewGroupService(&mockGroupRepositority)

	dummyGroup = data_model.Group{
		GroupID: 1234,
		GroupName: "DummyGroup",
		GroupURL: "www.schedulii.test.com",
		AvailableStartHour: 9,
		AvailableEndHour: 22,
	}
}

func TestGroupService(t *testing.T) {
	initGroupServiceTest()

	t.Run(
		"Create group returns successfully",
		func(t *testing.T) {
			mockCreateGroup = func(group data_model.Group) error {
				return nil
			}

			err := groupService.CreateGroup(dummyGroup)
			assert.Nil(t, err)
		},
	)

	t.Run(
		"Create group errors",
		func(t *testing.T) {
			mockCreateGroup = func(group data_model.Group) error {
				return pgx.ErrNoRows
			}

			err := groupService.CreateGroup(dummyGroup)
			assert.Error(t, err)
		},
	)

	t.Run(
		"Read group returns successfully",
		func(t *testing.T) {
			mockReadGroup = func(group data_model.Group) (data_model.Group, error) {
				return group, nil
			}

			result, err := groupService.ReadGroup(dummyGroup)
			assert.Nil(t, err)
			assert.Equal(t, result, dummyGroup)
		},
	)

	t.Run(
		"Read group errors",
		func(t *testing.T) {
			mockReadGroup = func(group data_model.Group) (data_model.Group, error) {
				return group, pgx.ErrNoRows
			}

			_, err := groupService.ReadGroup(dummyGroup)
			assert.Error(t, err)
		},
	)

	t.Run(
		"Update group returns successfully",
		func(t *testing.T) {
			mockUpdateGroup = func(group data_model.Group) error {
				return nil
			}

			err := groupService.UpdateGroup(dummyGroup)
			assert.Nil(t, err)
		},
	)

	t.Run(
		"Update group errors",
		func(t *testing.T) {
			mockUpdateGroup = func(group data_model.Group) error {
				return pgx.ErrNoRows
			}

			err := groupService.UpdateGroup(dummyGroup)
			assert.Error(t, err)
		},
	)
}
