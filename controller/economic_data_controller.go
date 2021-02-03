package controller

import (
	"data-app-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EconomicDataController struct {
	service *service.EconomicDataService
}

func NewEconomicDataController(service *service.EconomicDataService) EconomicDataController {
	return EconomicDataController{service: service}
}

func Test(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{})
}
