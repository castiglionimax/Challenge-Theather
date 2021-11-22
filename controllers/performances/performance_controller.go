package performances

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/castiglionimax/MeliShows-Challenge/domain/pagination"
	"github.com/castiglionimax/MeliShows-Challenge/domain/performance"
	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
	"github.com/castiglionimax/MeliShows-Challenge/services/performaService"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid request body")
		c.JSON(restErr.Status, restErr)
		return
	}

	var (
		query             queries.EsQuery
		performancesFound []performance.Performance
		errRest           *errors.RestErr
		pagination        pagination.Pagination
	)

	pagination.Limit = c.DefaultQuery("limit", "100")
	pagination.Offset = c.DefaultQuery("offset", "0")

	if len(bytes) > 0 {

		if err := json.Unmarshal(bytes, &query); err != nil {
			restErr := errors.NewBadRequestError("invalid json body")
			c.JSON(restErr.Status, restErr)
			return
		}

		performancesFound, errRest = performaService.PerformancesService.Search(query, pagination)
		if errRest != nil {
			c.JSON(http.StatusInternalServerError, errRest)
			return
		}

	} else {
		performancesFound, errRest = performaService.PerformancesService.GetAll()
		if errRest != nil {
			c.JSON(http.StatusInternalServerError, errRest)
			return
		}
	}
	//

	c.JSON(http.StatusOK, performancesFound)

}
