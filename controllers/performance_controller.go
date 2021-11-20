package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
	"github.com/castiglionimax/MeliShows-Challenge/services"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Showname string `form:"showname" json:"showname"`
	Start    string `form:"start" json:"start"`
	End      string `form:"end" json:"end"`
}

type Pagination struct {
	Limit  string `form:"limit" json:"limit"`
	Offset string `form:"offset" json:"offset"`
}

func GetAllPerformances(c *gin.Context) {
	/*
	   	//performacesFound, err := services.PerformancesService.GetAllPerformances()

	   	_ = performacesFound
	   	if err != nil {
	   		c.JSON(http.StatusInternalServerError, err)
	   	}

	   	var incomingQuery Query
	   	if c.Bind(&incomingQuery) == nil {
	   		log.Println("====== Bind By Query String ======")
	   		log.Println(incomingQuery.Showname)
	   		log.Println(time.Now().UTC().Format(c.Query(incomingQuery.Start)))
	   		log.Println(time.Now().UTC().Format(c.Query(incomingQuery.End)))
	   	}

	   //	c.JSON(http.StatusOK, performacesFound)
	*/
}

/*
 /performances?limit=20&offset=100. This query would return the 20 rows starting with the 100th row.

…/defects	Returns the first 25 defects (the default limit is 25).
…/defects?limit=10	Returns the first 10 defects.
…/defects?offset=5&limit=5	Returns defects 6..10.
…/defects?offset=10	Returns defects 11..36 (the default number of the returned defects is 25).
/*/
func GetPerformances(c *gin.Context) {
	/*
		var incomingQuery Query
		var pagingQuery Pagination

		c.Bind(&pagingQuery)
		parameter1, err := time.Parse("2006-01-02T15:04", incomingQuery.Start)
		if err != nil {
			fmt.Println(err)
		}
		parameter2, err := time.Parse("2006-01-02T15:04", incomingQuery.End)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("My Date Reformatted:\t", parameter2.Format(time.RFC822))

		performacesFound, err := services.PerformancesService.GetQuery(incomingQuery.Showname, parameter1, parameter2)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, performacesFound)
	*/
}

func Search(c *gin.Context) {

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid request body")
		c.JSON(restErr.Status, restErr)
		return
	}

	var query queries.EsQuery

	if err := json.Unmarshal(bytes, &query); err != nil {
		restErr := errors.NewBadRequestError("invalid request body")
		c.JSON(restErr.Status, restErr)
		return
	}
	fmt.Print(query)

	performancesFound, Searcherr := services.PerformancesService.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Searcherr)
		return
	}

	c.JSON(http.StatusOK, performancesFound)

}
