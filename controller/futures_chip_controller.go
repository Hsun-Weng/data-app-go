package controller

import (
	"data-app-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FuturesChipController struct {
	service service.FuturesChipService
}

func NewFuturesChipController(service service.FuturesChipService) FuturesChipController {
	return FuturesChipController{service: service}
}

func (controller *FuturesChipController) GetFuturesChips(context *gin.Context) {
	futuresCode := context.Param("futuresCode")
	startDateStr := context.Query("startDate")
	endDateStr := context.Query("endDate")

	context.JSON(http.StatusOK, controller.service.GetFuturesChips(futuresCode, startDateStr, endDateStr))
}
