package bookingsService

import (
	"github.com/castiglionimax/MeliShows-Challenge/domain/booking"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"
)

var BookingsService bookingsServicesInterface = &bookingsService{}

type bookingsService struct{}

type bookingsServicesInterface interface {
	Create(booking booking.Booking) (*booking.Booking, *errors.RestErr)
}

func (b *bookingsService) Create(booking booking.Booking) (*booking.Booking, *errors.RestErr) {

	if err := booking.SaveMongo(); err != nil {
		return nil, err
	}
	return &booking, nil
}
