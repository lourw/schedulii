package services

import (
	"testing"
	"time"

	"schedulii/src/models"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

var eventService EventService
var mockEventRepository MockEventRepository
var dummyEvent models.Event

type MockEventRepository struct{}

var mockCreateEvent func(event models.Event) error
var mockReadEvent func(event models.Event) (models.Event, error)
var mockUpdateEvent func(event models.Event) error

func (mock *MockEventRepository) Create(event models.Event) error {
	return mockCreateEvent(event)
}

func (mock *MockEventRepository) Read(event models.Event) (models.Event, error) {
	return mockReadEvent(event)
}

func (mock *MockEventRepository) Update(event models.Event) error {
	return mockUpdateEvent(event)
}

func initEventServiceTest() {
	mockEventRepository = MockEventRepository{}
	eventService = NewEventService(&mockEventRepository)

	dummyEvent = models.Event{
		EventId:   1234,
		GroupId:   1234,
		EventName: "Sample",
		StartTime: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2023, time.January, 1, 1, 20, 0, 0, time.UTC),
	}
}

func TestEventService(t *testing.T) {
	initEventServiceTest()

	t.Run(
		"Create event returns correctly",
		func(t *testing.T) {
			mockCreateEvent = func(event models.Event) error {
				return nil
			}

			err := eventService.CreateEvent(dummyEvent)
			assert.Nil(t, err)
		},
	)

	t.Run(
		"Create event throws an error",
		func(t *testing.T) {
			mockCreateEvent = func(event models.Event) error {
				return pgx.ErrNoRows
			}

			err := eventService.CreateEvent(dummyEvent)
			assert.Error(t, err)
		},
	)

	t.Run(
		"Read event returns correctly",
		func(t *testing.T) {
			mockReadEvent = func(event models.Event) (models.Event, error) {
				return event, nil
			}

			result, err := eventService.ReadEvent(dummyEvent)
			assert.Nil(t, err)
			assert.Equal(t, result, dummyEvent)
		},
	)

	t.Run(
		"Read event returns an error",
		func(t *testing.T) {
			mockReadEvent = func(event models.Event) (models.Event, error) {
				return event, pgx.ErrNoRows
			}

			_, err := eventService.ReadEvent(dummyEvent)
			assert.Error(t, err)
		},
	)

	t.Run(
		"Update event returns successfully",
		func(t *testing.T) {
			mockUpdateEvent = func(event models.Event) error {
				return nil
			}

			err := eventService.UpdateEvent(dummyEvent)
			assert.Nil(t, err)
		},
	)

	t.Run(
		"Update event returns an error",
		func(t *testing.T) {
			mockUpdateEvent = func(event models.Event) error {
				return pgx.ErrNoRows
			}

			err := eventService.UpdateEvent(dummyEvent)
			assert.Error(t, err)
		},
	)
}
