package data_srv

import (
	"context"
	"schedulii/src/models/data_model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupService struct {
	db *pgxpool.Pool
}

func NewGroupService(db *pgxpool.Pool) GroupService {
	return GroupService{
		db: db,
	}
}

func (gs *GroupService) CreateGroup(group data_model.Group) error {
	query := `
		INSERT INTO groups 
		VALUES ()
	`
	_, err := gs.db.Exec(
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

func (gs *GroupService) ReadGroup(group data_model.Group) (data_model.Group, error) {
	query := `
		SELECT * FROM groups 
		WHERE group_id = ($1)
	`
	queryResult := gs.db.QueryRow(
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

func (gs *GroupService) UpdateGroup(group data_model.Group) error {
	query := `
		UPDATE groups
		SET group_name = ($2),
			group_url = ($3),
			available_start_hour = ($4),
			available_end_hour = ($5)
		WHERE group_id = ($1)
	`
	_, err := gs.db.Exec(
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
