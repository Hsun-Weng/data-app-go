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

func InitFuturesChipController(db *mongo.Database) controller.FuturesChipController {
	wire.Build(repository.NewFuturesChipRepository, service.NewFuturesChipService, controller.NewFuturesChipController)
	return controller.FuturesChipController{}
}

func InitStockChipController(db *mongo.Database) controller.StockChipController {
	wire.Build(repository.NewStockChipRepository, service.NewStockChipService, controller.NewStockChipController)
	return controller.StockChipController{}
}

func InitStockIndexController(db *mongo.Database) controller.StockIndexController {
	wire.Build(repository.NewStockIndexRepository, service.NewStockIndexService, controller.NewStockIndexController)
	return controller.StockIndexController{}
}