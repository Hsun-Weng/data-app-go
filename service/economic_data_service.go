package service

import (
	"data-app-go/dto"
	"data-app-go/model"
	"data-app-go/repository"
	"fmt"
)

type EconomicDataService struct {
	repository repository.EconomicDataRepository
}

func NewEconomicDataService(repository repository.EconomicDataRepository) EconomicDataService {
	return EconomicDataService{repository: repository}
}

func (service *EconomicDataService) GetValues(countryCode string, dataCode string) []dto.EconomicDataDTO {
	return ToEconomicDataDTOs(service.repository.FindEconomicValuesByCountryCodeAndDataCode(countryCode, dataCode))
}

func ToEconomicDataDTO(economicData model.EconomicData) dto.EconomicDataDTO {
	return dto.EconomicDataDTO{
		Date:  fmt.Sprintf("%d-%02d-%02d", economicData.Date.Year(), economicData.Date.Month(), economicData.Date.Day()),
		Value: economicData.Value,
	}
}

func ToEconomicDataDTOs(economicDataArrays []model.EconomicData) []dto.EconomicDataDTO {
	economicDataDTOs := make([]dto.EconomicDataDTO, len(economicDataArrays))

	for i, item := range economicDataArrays {
		economicDataDTOs[i] = ToEconomicDataDTO(item)
	}

	return economicDataDTOs
}
