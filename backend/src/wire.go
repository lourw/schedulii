//+build wireinject

package main

import (
	"schedulii/src/db"
	"schedulii/src/handlers/data_handler"
	"schedulii/src/repositories"
	"schedulii/src/routes"
	"schedulii/src/services/data_srv"

	"github.com/google/wire"
	"github.com/gin-gonic/gin"
)

var AppSet = wire.NewSet(
	db.NewDatabaseConnection,
	gin.Default, 

	repositories.NewEventRepository,
	data_srv.NewEventService,
	data_handler.NewEventHandler,

	repositories.NewGroupRepository,
	data_srv.NewGroupService,
	data_handler.NewGroupHandler,

	repositories.NewUserRepository,
	data_srv.NewUserService,
	data_handler.NewUserHandler,

	routes.NewRouter,
	NewScheduliiApp, 
)
func InitializeApp() (ScheduliiApp, error) {
	wire.Build(AppSet)
	return ScheduliiApp{}, nil
}
