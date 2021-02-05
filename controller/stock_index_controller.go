package controller

import (
	"data-app-go/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type StockIndexController struct {
	service service.StockIndexService
}

func NewStockIndexController(service service.StockIndexService) StockIndexController {
	return StockIndexController{service: service}
}

func (controller *StockIndexController) GetStockIndexPrices(context *gin.Context) {
	indexCode := context.Param("indexCode")
	startDateStr := context.Query("startDate")
	endDateStr := context.Query("endDate")

	context.JSON(http.StatusOK, controller.service.GetStockIndexPrices(indexCode, startDateStr, endDateStr))
}

func (controller *StockIndexController) GetStockIndexLatest(context *gin.Context) {
	indexCode := context.Param("indexCode")

	context.JSON(http.StatusOK, controller.service.GetStockIndexLatest(indexCode))
}

func (controller *StockIndexController) GetStockIndexesLatest(context *gin.Context) {
	var indexCodes []string
	err := context.BindJSON(&indexCodes)
	if err != nil {
		log.Fatalln(err)
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusOK, controller.service.GetStockIndexesLatest(indexCodes))
}