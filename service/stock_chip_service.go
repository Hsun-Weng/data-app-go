package service

import (
	"data-app-go/dto"
	"data-app-go/model"
	"data-app-go/repository"
	"fmt"
	"log"
	"time"
)

type StockChipService struct {
	repository repository.StockChipRepository
}

func NewStockChipService(repository repository.StockChipRepository) StockChipService {
	return StockChipService{repository: repository}
}

func (service *StockChipService) GetStockChips(stockCode string, startDateStr string, endDateStr string) []dto.StockChipDTO {
	startDate, startDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", startDateStr))
	endDate, endDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59Z", endDateStr))
	if startDateParseErr != nil {
		log.Fatal(startDateParseErr)
	}
	if endDateParseErr != nil {
		log.Fatal(endDateParseErr)
	}
	return ToStockChipDTOs(service.repository.FindStockChipsByStockCodeAndDateBetween(stockCode, startDate, endDate))
}

func ToStockChipDTO(stockChip *model.StockChip) dto.StockChipDTO {
	return dto.StockChipDTO{
		Date:             fmt.Sprintf("%d-%02d-%02d", stockChip.Date.Year(), stockChip.Date.Month(), stockChip.Date.Day()),
		StockCode:        stockChip.StockCode,
		InvestorChipDTOs: ToInvestorStockChipDTOs(stockChip.InvestorChips),
	}
}

func ToStockChipDTOs(stockChipArray []*model.StockChip) []dto.StockChipDTO {
	stockChipDTOs := make([]dto.StockChipDTO, len(stockChipArray))
	for i, item := range stockChipArray {
		stockChipDTOs[i] = ToStockChipDTO(item)
	}

	return stockChipDTOs
}

func ToInvestorStockChipDTO(investorStockChip model.InvestorStockChip) dto.InvestorStockChipDTO {
	return dto.InvestorStockChipDTO{
		InvestorCode: investorStockChip.InvestorCode,
		LongShare:    investorStockChip.LongShare,
		ShortShare:   investorStockChip.ShortShare,
	}
}

func ToInvestorStockChipDTOs(investorStockChipArray []model.InvestorStockChip) []dto.InvestorStockChipDTO {
	investorStockChipDTOs := make([]dto.InvestorStockChipDTO, len(investorStockChipArray))

	for i, item := range investorStockChipArray {
		investorStockChipDTOs[i] = ToInvestorStockChipDTO(item)
	}

	return investorStockChipDTOs
}
