package service

import (
	"data-app-go/dto"
	"data-app-go/model"
	"data-app-go/repository"
	"fmt"
	"log"
	"time"
)

type StockMarginService struct {
	repository repository.StockMarginRepository
}

func NewStockMarginService(repository repository.StockMarginRepository) StockMarginService {
	return StockMarginService{repository: repository}
}

func (service *StockMarginService) GetStockMargins(stockCode string, startDateStr string, endDateStr string) []dto.StockMarginDTO {
	startDate, startDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", startDateStr))
	endDate, endDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59Z", endDateStr))
	if startDateParseErr != nil{
		log.Fatal(startDateParseErr)
	}
	if endDateParseErr != nil{
		log.Fatal(endDateParseErr)
	}
	return ToStockMarginDTOs(service.repository.FindStockMarginsByStockCodeAndDateBetween(stockCode, startDate, endDate))
}

func ToStockMarginDTO(stockMargin model.StockMargin) dto.StockMarginDTO {
	return dto.StockMarginDTO{
		Date : fmt.Sprintf("%d-%02d-%02d", stockMargin.Date.Year(), stockMargin.Date.Month(), stockMargin.Date.Day()),
		StockCode :stockMargin.StockCode,
		LongShare :stockMargin.LongShare,
		TotalLongShare :stockMargin.TotalLongShare,
		ShortShare :stockMargin.ShortShare,
		TotalShortShare :stockMargin.TotalShortShare,
		DayShare :stockMargin.DayShare,
	}
}

func ToStockMarginDTOs(stockMarginArray []model.StockMargin) []dto.StockMarginDTO {
	stockMarginDTOs := make([]dto.StockMarginDTO, len(stockMarginArray))
	for i, item := range stockMarginArray {
		stockMarginDTOs[i] = ToStockMarginDTO(item)
	}

	return stockMarginDTOs
}
