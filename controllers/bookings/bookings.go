package bookings

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/castiglionimax/MeliShows-Challenge/domain/booking"
	"github.com/castiglionimax/MeliShows-Challenge/services/bookingsService"
	"github.com/castiglionimax/MeliShows-Challenge/services/performaService"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid request body")
		c.JSON(restErr.Status, restErr)
		return
	}

	var booking booking.Booking
	if err := json.Unmarshal(bytes, &booking); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	perfor, errSaver := performaService.PerformancesService.Get(booking.PerformanceID)
	if errSaver != nil {
		c.JSON(errSaver.Status, errSaver)
		return
	}

	for _, s := range booking.Sold {

		errSaver = performaService.PerformancesService.ValidatePrice(perfor, s.SectionID, s.Seat)
		if errSaver != nil {
			c.JSON(errSaver.Status, errSaver)
			return
		}
	}
	var totalPrice float32
	totalPrice = 0
	for i := 0; i < len(booking.Sold); i++ {
		price := perfor.UpdateSeats(booking.Sold[i].SectionID, booking.Sold[i].Seat)
		booking.Sold[i].SetPrice(price)
		totalPrice += price
	}

	errUpdate := performaService.PerformancesService.UpdateES(perfor)
	if errUpdate != nil {
		c.JSON(errUpdate.Status, errUpdate)
		return
	}
	booking.TotalPrice = totalPrice

	result, saveErr := bookingsService.BookingsService.Create(booking)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}
