package data_srv

import (
	"schedulii/src/models/data_model"
	"schedulii/src/repositories"
)

type EventService struct {
	er repositories.EventRepository
}

func NewEventService(er repositories.EventRepository) EventService {
	return EventService{
		er: er,
	}
}

func (es *EventService) CreateEvent(event data_model.Event) error {
	err := es.er.CreateEvent(event)
	if err != nil {
		return err
	}
	return nil
}

func (es *EventService) GetEvent(event data_model.Event) (data_model.Event, error) {
	event, err := es.er.GetEvent(event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (es *EventService) UpdateEvent(event data_model.Event) error {
	err := es.er.UpdateEvent(event)
	if err != nil {
		return err
	}
	return nil
}
