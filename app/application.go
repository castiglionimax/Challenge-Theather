package app

import (
	"log"

	"github.com/castiglionimax/MeliShows-Challenge/cacheinmem"
	"github.com/castiglionimax/MeliShows-Challenge/database/elastics"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func StartApplication() {

	errxa := godotenv.Load()
	if errxa != nil {
		log.Fatal("Error loading .env file")
	}

	elastics.Init()
	cacheinmem.Init()
	mapUrls()
	router.Run(":5000")
}
