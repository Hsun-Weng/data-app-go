package controller

import (
	"data-app-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StockChipController struct {
	service service.StockChipService
}

func NewStockChipController(service service.StockChipService) StockChipController {
	return StockChipController{service: service}
}

func (controller *StockChipController) GetStockChips(context *gin.Context) {
	stockCode := context.Param("stockCode")
	startDateStr := context.Query("startDate")
	endDateStr := context.Query("endDate")

	context.JSON(http.StatusOK, controller.service.GetStockChips(stockCode, startDateStr, endDateStr))
}
