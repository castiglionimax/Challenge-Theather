package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/castiglionimax/MeliShows-Challenge/services"
	"github.com/gin-gonic/gin"
)

// /performances?showname=Aladdin&start=2021-11-01T00:30&end=2021-11-30T00
func GetAllPerformances(c *gin.Context) {
	//c.Param("show-name")
	performacesFound, err := services.PerformancesService.GetAllPerformances()

	_ = performacesFound
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	showname := c.Query("showname")
	startAt := time.Now().UTC().Format(c.Query("start"))
	endAt := time.Now().UTC().Format(c.Query("end"))

	fmt.Print(startAt)
	c.String(http.StatusOK, "Hello %s %s %s", showname, startAt, endAt)
}
