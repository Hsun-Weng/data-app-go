package service

import (
	"data-app-go/dto"
	"data-app-go/model"
	"data-app-go/repository"
	"fmt"
	"log"
	"math"
	"time"
)

type StockIndexService struct {
	repository repository.StockIndexRepository
}

func NewStockIndexService(repository repository.StockIndexRepository) StockIndexService {
	return StockIndexService{repository: repository}
}

func (service *StockIndexService) GetStockIndexPrices(indexCode string, startDateStr string, endDateStr string) []dto.StockIndexDTO {
	startDate, startDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", startDateStr))
	endDate, endDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59Z", endDateStr))
	if startDateParseErr != nil {
		log.Fatal(startDateParseErr)
	}
	if endDateParseErr != nil {
		log.Fatal(endDateParseErr)
	}
	return ToStockIndexDTOs(service.repository.FindStockIndexesByIndexCodeAndDateBetween(indexCode, startDate, endDate))
}

func (service *StockIndexService) GetStockIndexLatest(indexCode string) dto.StockIndexDTO {
	return ToStockIndexDTO(service.repository.FindStockIndexLatestByIndexCode(indexCode))
}

func (service *StockIndexService) GetStockIndexesLatest(indexCodes []string) []dto.StockIndexDTO {
	latestStockIndex, err := service.repository.FindStockIndexLatest()
	if err != nil {
		return nil
	}
	dateStr := fmt.Sprintf("%d-%02d-%02d", latestStockIndex.Date.Year(), latestStockIndex.Date.Month(), latestStockIndex.Date.Day())
	startDate, startDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", dateStr))
	endDate, endDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59Z", dateStr))
	if startDateParseErr != nil {
		log.Fatal(startDateParseErr)
	}
	if endDateParseErr != nil {
		log.Fatal(endDateParseErr)
	}
	return ToStockIndexDTOs(service.repository.FindStockIndexesByIndexCodesAndDateBetween(indexCodes, startDate, endDate))
}

func ToStockIndexDTO(stockIndex *model.StockIndex) dto.StockIndexDTO {
	return dto.StockIndexDTO{
		Date:          fmt.Sprintf("%d-%02d-%02d", stockIndex.Date.Year(), stockIndex.Date.Month(), stockIndex.Date.Day()),
		IndexCode:     stockIndex.IndexCode,
		Open:          stockIndex.Open,
		Low:           stockIndex.Low,
		High:          stockIndex.High,
		Close:         stockIndex.Close,
		Volume:        stockIndex.Volume,
		Change:        stockIndex.Change,
		ChangePercent: math.Trunc(stockIndex.ChangePercent*100*1e2+0.5) * 1e-2,
	}
}

func ToStockIndexDTOs(stockIndexArray []*model.StockIndex) []dto.StockIndexDTO {
	stockIndexDTOs := make([]dto.StockIndexDTO, len(stockIndexArray))
	for i, item := range stockIndexArray {
		stockIndexDTOs[i] = ToStockIndexDTO(item)
	}

	return stockIndexDTOs
}
