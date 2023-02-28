package services

import (
	"testing"

	"schedulii/src/models"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

var groupService GroupService
var mockGroupRepositority MockGroupRepository
var dummyGroup models.Group

type MockGroupRepository struct{}

var mockCreateGroup func(group models.Group) error
var mockReadGroup func(group models.Group) (models.Group, error)
var mockUpdateGroup func(group models.Group) error

func (mock *MockGroupRepository) Create(group models.Group) error {
	return mockCreateGroup(group)
}

func (mock *MockGroupRepository) Read(group models.Group) (models.Group, error) {
	return mockReadGroup(group)
}

func (mock *MockGroupRepository) Update(group models.Group) error {
	return mockUpdateGroup(group)
}

func initGroupServiceTest() {
	mockGroupRepositority = MockGroupRepository{}
	groupService = NewGroupService(&mockGroupRepositority)

	dummyGroup = models.Group{
		GroupID:            1234,
		GroupName:          "DummyGroup",
		GroupURL:           "www.schedulii.test.com",
		AvailableStartHour: 9,
		AvailableEndHour:   22,
	}
}

func TestGroupService(t *testing.T) {
	initGroupServiceTest()

	t.Run(
		"Create group returns successfully",
		func(t *testing.T) {
			mockCreateGroup = func(group models.Group) error {
				return nil
			}

			err := groupService.CreateGroup(dummyGroup)
			assert.Nil(t, err)
		},
	)

	t.Run(
		"Create group errors",
		func(t *testing.T) {
			mockCreateGroup = func(group models.Group) error {
				return pgx.ErrNoRows
			}

			err := groupService.CreateGroup(dummyGroup)
			assert.Error(t, err)
		},
	)

	t.Run(
		"Read group returns successfully",
		func(t *testing.T) {
			mockReadGroup = func(group models.Group) (models.Group, error) {
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
			mockReadGroup = func(group models.Group) (models.Group, error) {
				return group, pgx.ErrNoRows
			}

			_, err := groupService.ReadGroup(dummyGroup)
			assert.Error(t, err)
		},
	)

	t.Run(
		"Update group returns successfully",
		func(t *testing.T) {
			mockUpdateGroup = func(group models.Group) error {
				return nil
			}

			err := groupService.UpdateGroup(dummyGroup)
			assert.Nil(t, err)
		},
	)

	t.Run(
		"Update group errors",
		func(t *testing.T) {
			mockUpdateGroup = func(group models.Group) error {
				return pgx.ErrNoRows
			}

			err := groupService.UpdateGroup(dummyGroup)
			assert.Error(t, err)
		},
	)
}
