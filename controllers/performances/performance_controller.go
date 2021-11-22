package performances

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
	"github.com/castiglionimax/MeliShows-Challenge/services/performaService"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit  string `form:"limit" json:"limit"`
	Offset string `form:"offset" json:"offset"`
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
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	//fmt.Print(query)

	performancesFound, errRest := performaService.PerformancesService.Search(query)
	if errRest != nil {
		c.JSON(http.StatusInternalServerError, errRest)
		return
	}

	c.JSON(http.StatusOK, performancesFound)

}
