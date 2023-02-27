package data_srv

import (
	"log"
	"schedulii/src/models/data_model"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

var eventService EventService
var mockEventRepository MockEventRepository
var dummyEvent data_model.Event

type MockEventRepository struct{}

var createEventMock func(event data_model.Event) error
var readEventMock func(event data_model.Event) (data_model.Event, error)
var updateEventMock func(event data_model.Event) error

func (mock *MockEventRepository) Create(event data_model.Event) error {
	return createEventMock(event)
}

func (mock *MockEventRepository) Read(event data_model.Event) (data_model.Event, error) {
	return readEventMock(event)
}

func (mock *MockEventRepository) Update(event data_model.Event) error {
	return updateEventMock(event)
}

func setup() {
	mockEventRepository = MockEventRepository{}
	eventService = EventService{
		repository: &mockEventRepository,
	}

	dummyEventStartTime, startErr := time.Parse(time.RFC822, "02 Jan 06 15:00 PST")
	dummyEventEndTime, endRrr := time.Parse(time.RFC822, "02 Jan 06 15:30 PST")
	if startErr != nil || endRrr != nil {
		log.Fatalf("Error: invalid times in event service setup")
		return
	}
	dummyEvent = data_model.Event{
		EventId:   1234,
		GroupId:   1234,
		EventName: "Sample",
		StartTime: dummyEventStartTime,
		EndTime:   dummyEventEndTime,
	}
}

func TestEventService(t *testing.T) {
	setup()

	t.Run(
		"Create event returns correctly",
		func(t *testing.T) {
			createEventMock = func(event data_model.Event) error {
				return nil
			}
		
			err := eventService.CreateEvent(dummyEvent)
			assert.Nil(t, err)
		},
	)

	t.Run(
		"Create event throws an error",
		func(t *testing.T) {
			createEventMock = func(event data_model.Event) error {
				return pgx.ErrNoRows
			}
		
			err := eventService.CreateEvent(dummyEvent)
			assert.Error(t, err)
		},
	)

	t.Run(
		"Read event returns correctly",
		func(t *testing.T) {
			readEventMock = func(event data_model.Event) (data_model.Event, error) {
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
			readEventMock = func(event data_model.Event) (data_model.Event, error) {
				return event, pgx.ErrNoRows
			}
		
			_, err := eventService.ReadEvent(dummyEvent)
			assert.Error(t, err)
		},
	)

	t.Run(
		"Update event returns successfully",
		func(t *testing.T) {
			updateEventMock = func(event data_model.Event) error {
				return nil
			}
		
			err := eventService.UpdateEvent(dummyEvent)
			assert.Nil(t, err)
		},
	)

	t.Run(
		"Update event returns an error",
		func(t *testing.T) {
			updateEventMock = func(event data_model.Event) error {
				return pgx.ErrNoRows
			}
		
			err := eventService.UpdateEvent(dummyEvent)
			assert.Error(t, err)
		},
	)
}
