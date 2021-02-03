//go:generate wire
//+build !wireinject

package main

import (
	"data-app-go/config"
	"data-app-go/controller"
	"data-app-go/repository"
	"data-app-go/service"

	"github.com/google/wire"
)

func InitEconomicDataController(Database *wire.ProviderSet) controller.EconomicDataController {
	wire.Build(repository.NewEconomicDataRepository, service.NewEconomicDataService, controller.NewEconomicDataController)
	return controller.EconomicDataController{}
}

var DbSet = wire.NewSet(config.ReadConfig(), config.NewMongoClient)
