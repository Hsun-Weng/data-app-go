package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	economicDataRouter := InitEconomicDataController(db)
	router.GET("/economic/:countryCode/:dataCode", economicDataRouter.GetValues)

	return router
}
