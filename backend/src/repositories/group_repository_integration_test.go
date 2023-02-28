//go:build integration
// +build integration

package repositories

import (
	"context"
	"log"
	"schedulii/src/db"
	"schedulii/src/models"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

var groupDb *pgxpool.Pool
var groupRepository GroupRepository
var groupCtx context.Context

var groupTestData []models.Group

func initGroupTest() {
	var err error

	groupDb, err = db.NewDatabaseConnection()
	if err != nil {
		log.Fatalf("Error initializing integration test: %s", err)
	}

	groupRepository = GroupRepository{
		db: groupDb,
	}
	groupCtx = context.Background()

	groupTestData = []models.Group{
		models.Group{
			GroupID:            1234,
			GroupName:          "DummyGroup",
			GroupURL:           "www.schedulii.test.com",
			AvailableStartHour: 9,
			AvailableEndHour:   22,
		},
		models.Group{
			GroupID:            1235,
			GroupName:          "DummyGroup2",
			GroupURL:           "www.schedulii.test2.com",
			AvailableStartHour: 10,
			AvailableEndHour:   24,
		},
	}
}

func afterEachGroupTest() {
	_, err := groupDb.Exec(groupCtx, "DELETE FROM groups")
	if err != nil {
		log.Fatalf("Error on groups table teardown: %s", err)
	}
}

func TestGroupRepository_Read(t *testing.T) {
	initGroupTest()

	t.Run(
		"Create group returns correctly",
		func(t *testing.T) {
			var rowCount int
			var err error

			groupDb.QueryRow(context.Background(), "SELECT COUNT(*) FROM groups").Scan(&rowCount)
			assert.Equal(t, 0, rowCount)

			for i, data := range groupTestData {
				err = groupRepository.Create(data)
				if err != nil {
					log.Fatalf("Error: %s", err)
					t.Fail()
				}

				groupDb.QueryRow(context.Background(), "SELECT COUNT(*) FROM groups").Scan(&rowCount)
				assert.Equal(t, i+1, rowCount)
			}

			afterEachGroupTest()
		},
	)

	t.Run(
		"Read group returns correctly",
		func(t *testing.T) {
			for _, data := range groupTestData {
				err := groupRepository.Create(data)
				if err != nil {
					log.Fatalf("Error: %s", err)
					t.Fail()
				}

				group, readErr := groupRepository.Read(data)
				if readErr != nil {
					t.Fail()
				}
				assert.Equal(t, data, group)
			}

			afterEachGroupTest()
		},
	)

	t.Run(
		"Update group returns correctly",
		func(t *testing.T) {
			for _, data := range groupTestData {
				err := groupRepository.Create(data)
				if err != nil {
					log.Fatalf("Error: %s", err)
					t.Fail()
				}

				updatedData := data
				updatedData.GroupName = "NewName"

				err = groupRepository.Update(updatedData)
				if err != nil {
					log.Fatalf("Error: %s", err)
					t.Fail()
				}

				group, readErr := groupRepository.Read(updatedData)
				if readErr != nil {
					t.Fail()
				}
				assert.Equal(t, updatedData, group)
			}

			afterEachGroupTest()
		},
	)

	t.Cleanup(func() {
		afterEachGroupTest()
		groupDb.Close()
	})
}
