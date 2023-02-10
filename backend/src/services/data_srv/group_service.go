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
	query := "SELECT * FROM Groups WHERE GroupID = ($1)"
	queryResult := gs.db.QueryRow(
		context.Background(),
		query,
		group.GroupID,
	)
	err := queryResult.Scan(
		&group.GroupID,
		&group.GroupURL,
		&group.AvailableStartHour,
		&group.GroupName,
		&group.AvailableEndHour,
	)
	if err != nil {
		return group, err
	}
	return group, nil
}

func (gs *GroupService) UpdateGroup(group data_model.Group) error {
	query := `
		UPDATE Groups
		SET GroupName = ($2),
			GroupURL = ($3),
			AvailableStartHour = ($4),
			AvailableEndHour = ($5)
		WHERE GroupID = ($1)
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
