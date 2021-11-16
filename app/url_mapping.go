package app

import "github.com/castiglionimax/MeliShows-Challenge/controllers"

func mapUrls() {
	//router.POST("/theter")
	router.GET("/ping", controllers.Ping)

	// /performances?showname=Aladdin&start=2021-11-01T00:30&end=2021-11-30T00
	router.GET("/performances", controllers.GetAllPerformances)

}

//router.GET("minesweeper/users/:user_id/games/:game_id/solution",

//	router.GET("minesweeper/users/:user_id/games",

/*
	 ISO 8601 UTC
	YYYY-MM-DD
	YYYY-MM-DDThh:mm<TZDSuffix>
	YYYY-MM-DDThh:mm:ss<TZDSuffix>

		// YYYY-MM-DDThh:mm:ss.sTZD. date
*/
