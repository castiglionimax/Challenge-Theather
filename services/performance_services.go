package services

import (
	"github.com/castiglionimax/MeliShows-Challenge/domain/performance"
	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"
)

var PerformancesService performanceServicesInterface = &performancesService{}

type performancesService struct{}

type performanceServicesInterface interface {
	GetAll() ([]*performance.Performance, error)
	Search(query queries.EsQuery) ([]performance.Performance, *errors.RestErr)
}

func (p *performancesService) GetAll() ([]*performance.Performance, error) {
	dao := &performance.Performance{}
	_ = dao
	return nil, nil
	//	return dao.GetAllbyShowName()
}

func (p *performancesService) Search(query queries.EsQuery) ([]performance.Performance, *errors.RestErr) {
	dao := &performance.Performance{}
	return dao.Search(query)
}
