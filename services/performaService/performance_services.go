package performaService

import (
	"fmt"

	"github.com/castiglionimax/MeliShows-Challenge/domain/pagination"
	"github.com/castiglionimax/MeliShows-Challenge/domain/performance"
	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"
)

var PerformancesService performanceServicesInterface = &performancesService{}

type performancesService struct{}

type performanceServicesInterface interface {
	Get(id int64) (*performance.Performance, *errors.RestErr)
	Search(query queries.EsQuery, pagination pagination.Pagination) ([]performance.Performance, *errors.RestErr)
	UpdateES(*performance.Performance) *errors.RestErr
	ValidatePrice(performance *performance.Performance, sectionId int64, numberSeat int) *errors.RestErr
	GetAll() ([]performance.Performance, *errors.RestErr)
}

func (p *performancesService) Get(id int64) (*performance.Performance, *errors.RestErr) {
	dao := &performance.Performance{}

	//	"performanceID": 1

	query := &queries.EsQuery{}

	Id := &queries.FieldValue{
		Field: "performanceID",
		Value: id}

	query.Id = *Id
	buf := query.BuildQueryID()

	rest, err := dao.Search(buf, nil)
	if err != nil {
		return nil, err
	}

	return &rest[0], nil
}

func (p *performancesService) GetAll() ([]performance.Performance, *errors.RestErr) {
	query := &queries.EsQuery{}
	dao := &performance.Performance{}

	buf := query.BuildQueryMatchAll()

	rest, err := dao.Search(buf, nil)
	if err != nil {
		return nil, err
	}

	return rest, nil
}

func (p *performancesService) Search(query queries.EsQuery, pagination pagination.Pagination) ([]performance.Performance, *errors.RestErr) {
	if err := query.ValidateOrderValue(); err != nil {
		return nil, err
	}
	dao := &performance.Performance{}

	resp, err := dao.Search(query.BuildQuery(), &pagination)

	for _, e := range resp {
		//deleting section out of range
		e.ValidatePrice(query)
	}

	return resp, err
}

func (p *performancesService) UpdateES(per *performance.Performance) *errors.RestErr {

	if err := per.Put(); err != nil {
		return err
	}

	return nil
}

func (p *performancesService) ValidatePrice(performance *performance.Performance, sectionId int64, numberSeat int) *errors.RestErr {

	if err := performance.ValidateAvailabilitySeat(sectionId, numberSeat); err != false {
		return errors.NewNotFoundError(fmt.Sprintf("error, seat %d in the section %d is not found ", numberSeat, sectionId))

	}

	return nil
}
