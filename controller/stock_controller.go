package controller

import (
	"data-app-go/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StockController struct {
	service service.StockService
}

func NewStockController(service service.StockService) StockController {
	return StockController{service: service}
}

func (controller *StockController) GetStockPrices(context *gin.Context) {
	stockCode := context.Param("stockCode")
	startDateStr := context.Query("startDate")
	endDateStr := context.Query("endDate")

	context.JSON(http.StatusOK, controller.service.GetStockPrices(stockCode, startDateStr, endDateStr))
}

func (controller *StockController) GetStockPriceLatest(context *gin.Context) {
	stockCode := context.Param("stockCode")

	context.JSON(http.StatusOK, controller.service.GetStockLatest(stockCode))
}

func (controller *StockController) GetStocksLatest(context *gin.Context) {
	var stockCodes []string
	err := context.BindJSON(&stockCodes)
	if err != nil {
		log.Fatalln(err)
		context.Status(http.StatusBadRequest)
		return
	}
	context.JSON(http.StatusOK, controller.service.GetStocksLatest(stockCodes))
}

func (controller *StockController) GetStocksPageLatest(context *gin.Context) {
	sortColumn := context.Param("sortColumn")
	page, _ := strconv.ParseInt(context.Param("page"), 10, 32)
	size, _ := strconv.ParseInt(context.Param("size"), 10, 32)
	direction := context.Param("direction")
	context.JSON(http.StatusOK, controller.service.GetStocksPageLatest(sortColumn, int(page), size, direction))
}
