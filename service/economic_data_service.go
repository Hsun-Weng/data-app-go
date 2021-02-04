package service

import (
	"data-app-go/dto"
	"data-app-go/model"
	"data-app-go/repository"
)

type EconomicDataService struct {
	repository repository.EconomicDataRepository
}

func NewEconomicDataService(repository repository.EconomicDataRepository) EconomicDataService {
	return EconomicDataService{repository: repository}
}

func (service *EconomicDataService) GetValues(countryCode string, dataCode string) []dto.EconomicDataDTO {
	return ToEconomicDataDTOs(service.repository.FindByCountryCodeAndDataCode(countryCode, dataCode))
}

func ToEconomicDataDTO(economicData model.EconomicData) dto.EconomicDataDTO {
	return dto.EconomicDataDTO{
		Date:  economicData.Date.Time().Format("2021-01-01"),
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
