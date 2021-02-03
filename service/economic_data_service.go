package service

import (
	"data-app-go/repository"
)

type EconomicDataService struct {
	repository *repository.EconomicDataRepository
}

func NewEconomicDataService(repository *repository.EconomicDataRepository) EconomicDataService {
	return EconomicDataService{repository: repository}
}

func (service *EconomicDataService) GetAll() []string {
	return service.repository.FindAll()
}
