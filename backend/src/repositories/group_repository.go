package repositories

import (
	"context"
	"schedulii/src/models/data_model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepository struct {
	db *pgxpool.Pool
}

func NewGroupRepository(db *pgxpool.Pool) *GroupRepository {
	return &GroupRepository{
		db: db,
	}
}

func (gr *GroupRepository) Create(group data_model.Group) error {
	query := `
		INSERT INTO groups
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := gr.db.Exec(
		context.Background(),
		query,
		group.GroupID,
		group.GroupName,
		group.GroupURL,
		group.AvailableStartHour,
		group.AvailableEndHour,
	)
	if err != nil {
		return err
	}
	return nil
}


func (gr *GroupRepository) Read(group data_model.Group) (data_model.Group, error) {
	query := `
		SELECT * FROM groups 
		WHERE group_id = ($1)
	`
	queryResult := gr.db.QueryRow(
		context.Background(),
		query,
		group.GroupID,
	)
	err := queryResult.Scan(
		&group.GroupID,
		&group.GroupURL,
		&group.GroupName,
		&group.AvailableStartHour,
		&group.AvailableEndHour,
	)
	if err != nil {
		return group, err
	}
	return group, nil
}

func (gr *GroupRepository) Update(group data_model.Group) error {
	query := `
		UPDATE groups
		SET group_name = ($2),
			group_url = ($3),
			available_start_hour = ($4),
			available_end_hour = ($5)
		WHERE group_id = ($1)
	`
	_, err := gr.db.Exec(
		context.Background(),
		query,
		group.GroupID,
		group.GroupName,
		group.GroupURL,
		group.AvailableStartHour,
		group.AvailableEndHour,
	)
	if err != nil {
		return err
	}
	return nil
}
