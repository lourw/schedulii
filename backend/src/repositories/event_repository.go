package repositories

import (
	"context"
	"schedulii/src/models/data_model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EventRepository struct {
	db *pgxpool.Pool
}

func NewEventRepository(db *pgxpool.Pool) EventRepository {
	return EventRepository{
		db: db,
	}
}

func (er *EventRepository) CreateEvent(event data_model.Event) error {
	query := `
		INSERT INTO events 
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := er.db.Exec(
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

func (er *EventRepository) GetEvent(event data_model.Event) (data_model.Event, error) {
	query := `
		SELECT * FROM events 
		WHERE event_id = ($1)
	`
	result := er.db.QueryRow(
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

func (er *EventRepository) UpdateEvent(event data_model.Event) error{
	query := `
		UPDATE events
		SET event_name = ($2),
			start_time = ($3),
			end_time = ($4)
		WHERE event_id = ($1)
	`
	_, err := er.db.Exec(
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
