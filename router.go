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
	futuresController := InitFuturesController(db)
	stockChipController := InitStockChipController(db)
	stockController := InitStockController(db)
	stockIndexController := InitStockIndexController(db)
	stockMarginController := InitStockMarginController(db)

	router.GET("/economic/:countryCode/:dataCode", economicDataController.GetValues)
	router.GET("/futures/:futuresCode/chip", futuresChipController.GetFuturesChips)
	router.GET("/futures/:futuresCode", futuresController.GetFuturesPrices)
	router.GET("/stock/:stockCode/chip", stockChipController.GetStockChips)
	router.GET("/stock/:stockCode/prices", stockController.GetStockPrices)
	router.GET("/stock/:stockCode/price/latest", stockController.GetStockPriceLatest)
	router.POST("/stocks/price/latest", stockController.GetStocksLatest)
	router.GET("/stocks/rank/price/latest", stockController.GetStocksPageLatest)
	router.GET("/index/:indexCode/prices", stockIndexController.GetStockIndexPrices)
	router.GET("/index/:indexCode/price/latest", stockIndexController.GetStockIndexLatest)
	router.POST("/indexes/price/latest", stockIndexController.GetStockIndexesLatest)
	router.GET("/stock/:stockCode/margin", stockMarginController.GetStockMargins)

	return router
}
