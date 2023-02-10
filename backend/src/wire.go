//+build wireinject

package main

import (
	"schedulii/src/db"
	"schedulii/src/handlers/data_handler"
	"schedulii/src/routes"
	"schedulii/src/services/data_srv"

	"github.com/google/wire"
	"github.com/gin-gonic/gin"
)

var AppSet = wire.NewSet(
	db.NewDatabaseConnection,
	gin.Default, 

	data_srv.NewEventService,
	data_handler.NewEventHandler,

	data_srv.NewGroupService,
	data_handler.NewGroupHandler,

	data_srv.NewUserService,
	data_handler.NewUserHandler,

	routes.NewRouter,
	NewScheduliiApp, 
)
func InitializeApp() (ScheduliiApp, error) {
	wire.Build(AppSet)
	return ScheduliiApp{}, nil
}
