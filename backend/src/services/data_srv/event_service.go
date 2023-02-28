package data_srv

import (
	"schedulii/src/models"
	"schedulii/src/models/data_model"
)

type EventService struct {
	repository models.Repository[data_model.Event]
}

func NewEventService(repository models.Repository[data_model.Event]) EventService {
	return EventService{
		repository: repository,
	}
}

func (es *EventService) CreateEvent(event data_model.Event) error {
	err := es.repository.Create(event)
	if err != nil {
		return err
	}
	return nil
}

func (es *EventService) ReadEvent(event data_model.Event) (data_model.Event, error) {
	event, err := es.repository.Read(event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (es *EventService) UpdateEvent(event data_model.Event) error {
	err := es.repository.Update(event)
	if err != nil {
		return err
	}
	return nil
}
