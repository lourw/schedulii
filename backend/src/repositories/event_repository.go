package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"schedulii/src/models"
)

type EventRepository struct {
	db *pgxpool.Pool
}

func NewEventRepository(db *pgxpool.Pool) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (er *EventRepository) Create(event models.Event) error {
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

func (er *EventRepository) Read(event models.Event) (models.Event, error) {
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

func (er *EventRepository) Update(event models.Event) error {
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
