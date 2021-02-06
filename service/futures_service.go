package service

import (
	"data-app-go/dto"
	"data-app-go/model"
	"data-app-go/repository"
	"fmt"
	"log"
	"time"
)

type FuturesService struct {
	repository repository.FuturesRepository
}

func NewFuturesService(repository repository.FuturesRepository) FuturesService {
	return FuturesService{repository: repository}
}

func (service *FuturesService) GetFuturesPrices(futuresCode string, contractDate string, startDateStr string, endDateStr string) []dto.FuturesDTO {
	startDate, startDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT00:00:00Z", startDateStr))
	endDate, endDateParseErr := time.Parse(time.RFC3339, fmt.Sprintf("%sT23:59:59Z", endDateStr))
	if startDateParseErr != nil {
		log.Fatal(startDateParseErr)
	}
	if endDateParseErr != nil {
		log.Fatal(endDateParseErr)
	}
	return ToFuturesDTOs(service.repository.FindFuturesByFuturesCodeAndContractDateAndDateBetween(futuresCode, contractDate, startDate, endDate))
}

func ToFuturesDTO(futures model.Futures) dto.FuturesDTO {
	return dto.FuturesDTO{
		Date:         fmt.Sprintf("%d-%02d-%02d", futures.Date.Year(), futures.Date.Month(), futures.Date.Day()),
		FuturesCode:  futures.FuturesCode,
		ContractDate: futures.ContractDate,
		Open:         futures.Open,
		Low:          futures.Low,
		High:         futures.High,
		Close:        futures.Close,
		Volume:       futures.Volume,
	}
}

func ToFuturesDTOs(futuresArray []model.Futures) []dto.FuturesDTO {
	futuresDTOs := make([]dto.FuturesDTO, len(futuresArray))
	for i, item := range futuresArray {
		futuresDTOs[i] = ToFuturesDTO(item)
	}

	return futuresDTOs
}
