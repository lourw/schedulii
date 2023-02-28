//go:build integration
// +build integration

package repositories

import (
	"context"
	"fmt"
	"log"
	"schedulii/src/db"
	"schedulii/src/models"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

var eventDb *pgxpool.Pool
var eventRepository EventRepository
var eventCtx context.Context

var testGroupId int
var testEventData []models.Event

func initEventTests() {
	var err error

	eventDb, err = db.NewDatabaseConnection()
	if err != nil {
		log.Fatalf("Error initializing integration test: %s", err)
	}

	eventRepository = EventRepository{
		db: eventDb,
	}
	eventCtx = context.Background()

	testGroupId = 123
	testEventData = []models.Event{
		models.Event{
			EventId:   1234,
			EventName: "TestEvent",
			GroupId:   testGroupId,
			StartTime: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2023, time.January, 1, 1, 30, 0, 0, time.UTC),
		},
		models.Event{
			EventId:   1235,
			EventName: "TestEvent2",
			GroupId:   testGroupId,
			StartTime: time.Date(2023, time.February, 2, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2023, time.February, 2, 1, 30, 0, 0, time.UTC),
		},
	}
}

func beforeEachEventTest() {
	// Events has a reliance on a group existing on database
	_, err := eventDb.Exec(
		eventCtx,
		fmt.Sprintf("INSERT INTO groups "+
			"VALUES (%d, 'test', 'test@email.com', 9, 10)",
			testGroupId,
		),
	)
	if err != nil {
		log.Fatalf("Error on setup: %s", err)
	}
}

func afterEachEventTest() {
	_, err := eventDb.Exec(eventCtx, "DELETE FROM events")
	if err != nil {
		log.Fatalf("Error on events table teardown: %s", err)
	}

	_, err = eventDb.Exec(eventCtx, "DELETE FROM groups")
	if err != nil {
		log.Fatalf("Error on groups table teardown: %s", err)
	}
}

func TestEventRepository(t *testing.T) {
	initEventTests()

	t.Run(
		"Create event returns correctly",
		func(t *testing.T) {
			beforeEachEventTest()

			var rowCount int
			var err error

			eventDb.QueryRow(context.Background(), "SELECT COUNT(*) FROM events").Scan(&rowCount)
			assert.Equal(t, 0, rowCount)

			for i, data := range testEventData {
				err = eventRepository.Create(data)
				if err != nil {
					log.Fatalf("Error: %s", err)
					t.Fail()
				}

				eventDb.QueryRow(context.Background(), "SELECT COUNT(*) FROM events").Scan(&rowCount)
				assert.Equal(t, i+1, rowCount)
			}

			afterEachEventTest()
		},
	)

	t.Run(
		"Read event returns correctly",
		func(t *testing.T) {
			beforeEachEventTest()

			for _, data := range testEventData {
				err := eventRepository.Create(data)
				if err != nil {
					log.Fatalf("Error: %s", err)
					t.Fail()
				}

				event, readErr := eventRepository.Read(data)
				if readErr != nil {
					t.Fail()
				}
				assert.Equal(t, data, event)
			}

			afterEachEventTest()
		},
	)

	t.Run(
		"Update event returns correctly",
		func(t *testing.T) {
			beforeEachEventTest()

			for _, data := range testEventData {
				err := eventRepository.Create(data)
				if err != nil {
					log.Fatalf("Error: %s", err)
					t.Fail()
				}

				updatedData := data
				updatedData.EventName = "NewName"

				err = eventRepository.Update(updatedData)
				if err != nil {
					log.Fatalf("Error: %s", err)
					t.Fail()
				}

				event, readErr := eventRepository.Read(updatedData)
				if readErr != nil {
					t.Fail()
				}
				assert.Equal(t, updatedData, event)
			}

			afterEachEventTest()
		},
	)

	t.Cleanup(func() {
		afterEachEventTest()
		eventDb.Close()
	})
}
