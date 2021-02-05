package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	economicDataController := InitEconomicDataController(db)
	futuresChipController := InitFuturesChipController(db)
	stockChipController := InitStockChipController(db)

	router.GET("/economic/:countryCode/:dataCode", economicDataController.GetValues)
	router.GET("/futures/:futuresCode/chip", futuresChipController.GetFuturesChips)
	router.GET("/stock/:stockCode/chip", stockChipController.GetStockChips)

	return router
}
