package service

import (
	"data-app-go/model"
	"data-app-go/repository"
)

type EconomicDataService struct {
	repository repository.EconomicDataRepository
}

func NewEconomicDataService(repository repository.EconomicDataRepository) EconomicDataService {
	return EconomicDataService{repository: repository}
}

func (service *EconomicDataService) GetValues(countryCode string, dataCode string) []*model.EconomicData {
	return service.repository.FindByCountryCodeAndDataCode(countryCode, dataCode)
}
