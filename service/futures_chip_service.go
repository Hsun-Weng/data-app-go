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

type FuturesChipService struct {
	repository repository.FuturesChipRepository
}

func NewFuturesChipService(repository repository.FuturesChipRepository) FuturesChipService {
	return FuturesChipService{repository: repository}
}

func (service *FuturesChipService) GetFuturesChips(futuresCode string, startDateStr string, endDateStr string) []dto.FuturesChipDTO {
	startDate, startDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", startDateStr))
	endDate, endDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59Z", endDateStr))
	if startDateParseErr != nil{
		log.Fatal(startDateParseErr)
	}
	if endDateParseErr != nil{
		log.Fatal(endDateParseErr)
	}
	return ToFuturesChipDTOs(service.repository.FindFuturesChipsByFuturesCodeAndDateBetween(futuresCode, startDate, endDate))
}

func ToFuturesChipDTO(futuresChip model.FuturesChip) dto.FuturesChipDTO {
	return dto.FuturesChipDTO{
		Date:  fmt.Sprintf("%d-%02d-%02d", futuresChip.Date.Year(), futuresChip.Date.Month(), futuresChip.Date.Day()),
		FuturesCode: futuresChip.FuturesCode,
		InvestorChipDTOs: ToInvestorFuturesChipDTOs(futuresChip.InvestorChips),
	}
}

func ToFuturesChipDTOs(futuresChipArray []model.FuturesChip) []dto.FuturesChipDTO {
	futuresChipDTOs := make([]dto.FuturesChipDTO, len(futuresChipArray))
	for i, item := range futuresChipArray {
		futuresChipDTOs[i] = ToFuturesChipDTO(item)
	}

	return futuresChipDTOs
}

func ToInvestorFuturesChipDTO(investorFuturesChip model.InvestorFuturesChip) dto.InvestorFuturesChipDTO {
	return dto.InvestorFuturesChipDTO{
		InvestorCode: investorFuturesChip.InvestorCode,
		OpenInterestLongLot: investorFuturesChip.OpenInterestLongLot,
		OpenInterestShortLot: investorFuturesChip.OpenInterestShortLot,
		OpenInterestNetLot: investorFuturesChip.OpenInterestLongLot - investorFuturesChip.OpenInterestShortLot,
		Percent: math.Trunc(investorFuturesChip.Percent*100*1e2 + 0.5)*1e-2,
	}
}

func ToInvestorFuturesChipDTOs(investorFuturesChipArray []model.InvestorFuturesChip) []dto.InvestorFuturesChipDTO {
	investorFuturesChipDTOs := make([]dto.InvestorFuturesChipDTO, len(investorFuturesChipArray))

	for i, item := range investorFuturesChipArray {
		investorFuturesChipDTOs[i] = ToInvestorFuturesChipDTO(item)
	}

	return investorFuturesChipDTOs
}
