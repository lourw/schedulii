package data_srv

import (
	"context"
	"schedulii/src/models/data_model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EventService struct {
	db *pgxpool.Pool 
}

func NewEventService(db *pgxpool.Pool) EventService {
	return EventService{
		db: db,
	}
}

func (es *EventService) CreateEvent(event data_model.Event) error {
	query := `
		INSERT INTO events 
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := es.db.Exec(
		context.Background(),
		query, 
		event.EventId, 
		event.GroupId, 
		event.EventName, 
		event.StartTime, 
		event.EndTime,
	)
	if err != nil {
		return err
	}
	return nil
}

func (es *EventService) ReadEvent(event data_model.Event) (data_model.Event, error) {
	query := `
		SELECT * FROM events 
		WHERE event_id = ($1)
	`
	result := es.db.QueryRow(
		context.Background(), 
		query, 
		event.EventId,
	)
	err := result.Scan(
		&event.EventId,
		&event.GroupId,
		&event.EventName,
		&event.StartTime,
		&event.EndTime,
	)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (es *EventService) UpdateEvent(event data_model.Event) error {
	query := `
		UPDATE events
		SET event_name = ($2),
			start_time = ($3),
			end_time = ($4)
		WHERE event_id = ($1)
	`
	_, err := es.db.Exec(
		context.Background(),
		query,
		event.EventId,
		event.EventName,
		event.StartTime,
		event.EndTime,
	)
	if err != nil {
		return err
	}
	return nil
}
