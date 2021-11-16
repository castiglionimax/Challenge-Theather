package services

import (
	"github.com/castiglionimax/MeliShows-Challenge/domain/performance"
)

var PerformancesService performanceServicesInterface = &performancesService{}

type performancesService struct{}

type performanceServicesInterface interface {
	GetAllPerformances() ([]*performance.Performance, error)
}

func (p *performancesService) GetAllPerformances() ([]*performance.Performance, error) {
	dao := &performance.Performance{}

	return nil, nil
}
