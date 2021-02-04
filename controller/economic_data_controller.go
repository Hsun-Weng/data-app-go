package controller

import (
	"data-app-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EconomicDataController struct {
	service service.EconomicDataService
}

func NewEconomicDataController(service service.EconomicDataService) EconomicDataController {
	return EconomicDataController{service: service}
}

func (controller *EconomicDataController) GetValues(context *gin.Context) {
	countryCode := context.Param("countryCode")
	dataCode := context.Param("dataCode")

	context.JSON(http.StatusOK, controller.service.GetValues(countryCode, dataCode))
}
