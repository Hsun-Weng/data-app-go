package controller

import (
	"data-app-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FuturesController struct {
	service service.FuturesService
}

func NewFuturesController(service service.FuturesService) FuturesController {
	return FuturesController{service: service}
}

func (controller *FuturesController) GetFuturesPrices(context *gin.Context) {
	futuresCode := context.Param("futuresCode")
	contractDate := context.Param("contractDate")
	startDateStr := context.Query("startDate")
	endDateStr := context.Query("endDate")

	context.JSON(http.StatusOK, controller.service.GetFuturesPrices(futuresCode, contractDate, startDateStr, endDateStr))
}
