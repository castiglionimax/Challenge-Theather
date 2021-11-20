package app

import (
	"github.com/castiglionimax/MeliShows-Challenge/database/elastics"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	elastics.Init()
	//ElasticSearchPrueba()
	mapUrls()
	router.Run(":5000")
}
