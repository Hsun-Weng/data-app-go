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

type StockService struct {
	repository repository.StockRepository
}

func NewStockService(repository repository.StockRepository) StockService {
	return StockService{repository: repository}
}

func (service *StockService) GetStockPrices(stockCode string, startDateStr string, endDateStr string) []dto.StockDTO {
	startDate, startDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", startDateStr))
	endDate, endDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59Z", endDateStr))
	if startDateParseErr != nil {
		log.Fatal(startDateParseErr)
	}
	if endDateParseErr != nil {
		log.Fatal(endDateParseErr)
	}
	return ToStockDTOs(service.repository.FindStocksByStockCodeAndDateBetween(stockCode, startDate, endDate))
}

func (service *StockService) GetStockLatest(stockCode string) dto.StockDTO {
	return ToStockDTO(service.repository.FindStockLatestByStockCode(stockCode))
}

func (service *StockService) GetStocksLatest(stockCodes []string) []dto.StockDTO {
	latestStock, err := service.repository.FindStockLatest()
	if err != nil {
		return nil
	}
	dateStr := fmt.Sprintf("%d-%02d-%02d", latestStock.Date.Year(), latestStock.Date.Month(), latestStock.Date.Day())
	startDate, startDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", dateStr))
	endDate, endDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59Z", dateStr))
	if startDateParseErr != nil {
		log.Fatal(startDateParseErr)
	}
	if endDateParseErr != nil {
		log.Fatal(endDateParseErr)
	}
	return ToStockDTOs(service.repository.FindStocksByStockCodesAndDateBetween(stockCodes, startDate, endDate))
}

func (service *StockService) GetStocksPageLatest(sortColumn string, page int, size int64, directionStr string) dto.StockPageDTO {
	latestStock, err := service.repository.FindStockLatest()
	if err != nil {
		return dto.StockPageDTO{}
	}
	dateStr := fmt.Sprintf("%d-%02d-%02d", latestStock.Date.Year(), latestStock.Date.Month(), latestStock.Date.Day())
	startDate, startDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", dateStr))
	endDate, endDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59Z", dateStr))
	if startDateParseErr != nil {
		log.Fatal(startDateParseErr)
	}
	if endDateParseErr != nil {
		log.Fatal(endDateParseErr)
	}
	direction := -1 //預設倒序
	if directionStr == "ASC" {
		direction = 1
	}
	totalSize := service.repository.CountStocksByDateBetween(startDate, endDate)

	return dto.StockPageDTO{
		TotalPage: totalSize / size,
		Page:      page,
		Size:      size,
		Content:   ToStockDTOs(service.repository.FindStocksByDateBetweenAndSortByAndLimit(startDate, endDate, sortColumn, direction, size)),
	}
}

func ToStockDTO(Stock *model.Stock) dto.StockDTO {
	return dto.StockDTO{
		Date:          fmt.Sprintf("%d-%02d-%02d", Stock.Date.Year(), Stock.Date.Month(), Stock.Date.Day()),
		StockCode:     Stock.StockCode,
		Open:          Stock.Open,
		Low:           Stock.Low,
		High:          Stock.High,
		Close:         Stock.Close,
		Volume:        Stock.Volume,
		Change:        Stock.Change,
		ChangePercent: math.Trunc(Stock.ChangePercent*100*1e2+0.5) * 1e-2,
	}
}

func ToStockDTOs(StockArray []*model.Stock) []dto.StockDTO {
	StockDTOs := make([]dto.StockDTO, len(StockArray))
	for i, item := range StockArray {
		StockDTOs[i] = ToStockDTO(item)
	}

	return StockDTOs
}
