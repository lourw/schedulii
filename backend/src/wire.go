//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"schedulii/src/db"
	"schedulii/src/handlers"
	"schedulii/src/models"
	"schedulii/src/repositories"
	"schedulii/src/routes"
	"schedulii/src/services"
)

var AppSet = wire.NewSet(
	db.NewDatabaseConnection,
	gin.Default,

	repositories.NewEventRepository,
	wire.Bind(new(models.Repository[models.Event]), new(*repositories.EventRepository)),
	services.NewEventService,
	handlers.NewEventHandler,

	repositories.NewGroupRepository,
	wire.Bind(new(models.Repository[models.Group]), new(*repositories.GroupRepository)),
	services.NewGroupService,
	handlers.NewGroupHandler,

	repositories.NewUserRepository,
	wire.Bind(new(models.Repository[models.User]), new(*repositories.UserRepository)),
	services.NewUserService,
	handlers.NewUserHandler,

	routes.NewRouter,
	NewScheduliiApp,
)

func InitializeApp() (ScheduliiApp, error) {
	wire.Build(AppSet)
	return ScheduliiApp{}, nil
}
