package controller

import (
	"data-app-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StockMarginController struct {
	service service.StockMarginService
}

func NewStockMarginController(service service.StockMarginService) StockMarginController {
	return StockMarginController{service: service}
}

func (controller *StockMarginController) GetStockMargins(context *gin.Context) {
	stockCode := context.Param("stockCode")
	startDateStr := context.Query("startDate")
	endDateStr := context.Query("endDate")

	context.JSON(http.StatusOK, controller.service.GetStockMargins(stockCode, startDateStr, endDateStr))
}
