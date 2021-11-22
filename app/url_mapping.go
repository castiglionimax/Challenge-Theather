package app

import (
	"github.com/castiglionimax/MeliShows-Challenge/controllers/bookings"
	"github.com/castiglionimax/MeliShows-Challenge/controllers/performances"
	"github.com/castiglionimax/MeliShows-Challenge/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	//	router.GET("/theater/{id}", controllers.GetTheater)

	router.POST("/performances/search", performances.Search)
	router.POST("/bookings/", bookings.Create)

}
