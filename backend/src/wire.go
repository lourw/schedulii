package main

import (
	"schedulii/src/services/data_srv"
	"schedulii/src/handlers/data_handler"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabaseConnection() *pgxpool.Pool {
	return setupDatabaseConnection()
}

func InitializeHandlers() {
	wire.Build(NewDatabaseConnection(), data_srv.NewEventService, data_handler.NewEventHandler)
}
