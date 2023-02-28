package services

import (
	"schedulii/src/models"
)

type EventService struct {
	repository models.Repository[models.Event]
}

func NewEventService(repository models.Repository[models.Event]) EventService {
	return EventService{
		repository: repository,
	}
}

func (es *EventService) CreateEvent(event models.Event) error {
	err := es.repository.Create(event)
	if err != nil {
		return err
	}
	return nil
}

func (es *EventService) ReadEvent(event models.Event) (models.Event, error) {
	event, err := es.repository.Read(event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (es *EventService) UpdateEvent(event models.Event) error {
	err := es.repository.Update(event)
	if err != nil {
		return err
	}
	return nil
}
