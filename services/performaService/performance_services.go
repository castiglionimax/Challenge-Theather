package performaService

import (
	"github.com/castiglionimax/MeliShows-Challenge/domain/performance"
	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"
)

var PerformancesService performanceServicesInterface = &performancesService{}

type performancesService struct{}

type performanceServicesInterface interface {
	Get(id int64) (*performance.Performance, *errors.RestErr)
	Search(query queries.EsQuery) ([]performance.Performance, *errors.RestErr)
	UpdateES(performance *performance.Performance) *errors.RestErr
}

func (p *performancesService) Get(id int64) (*performance.Performance, *errors.RestErr) {
	dao := &performance.Performance{}

	//	"performanceID": 1

	query := queries.EsQuery{
		Id: queries.FieldValue{
			Field: "performanceID",
			Value: id,
		},
	}

	rest, err := dao.Search(query)
	return &rest[1], err
}

func (p *performancesService) Search(query queries.EsQuery) ([]performance.Performance, *errors.RestErr) {
	if err := query.ValidateOrderValue(); err != nil {
		return nil, err
	}
	dao := &performance.Performance{}

	return dao.Search(query)
}

func (p *performancesService) UpdateES(performance *performance.Performance) *errors.RestErr {

	if err := performance.Put(); err != nil {
		return err
	}

	return nil
}
