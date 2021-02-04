// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"data-app-go/controller"
	"data-app-go/repository"
	"data-app-go/service"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/wire"
)

func InitEconomicDataController(db *mongo.Database) controller.EconomicDataController {
	wire.Build(repository.NewEconomicDataRepository, service.NewEconomicDataService, controller.NewEconomicDataController)
	return controller.EconomicDataController{}
}
