package performances

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/castiglionimax/MeliShows-Challenge/cacheinmem"

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
		b                 strings.Builder
	)

	pagination.Limit = c.DefaultQuery("limit", "100")
	pagination.Offset = c.DefaultQuery("offset", "0")
	h := sha1.New()

	h.Write(bytes)

	b.WriteString(hex.EncodeToString(h.Sum(nil)))
	b.WriteString(fmt.Sprintf("%s%s", pagination.Limit, pagination.Offset))

	hashCache := b.String()
	if cacheinmem.Client.GetINuser(hashCache) == true {
		asd := cacheinmem.Client.GetOutUser(hashCache)
		c.JSON(http.StatusOK, asd)
		return
	}

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

	cacheinmem.Client.SetINuser(hashCache, query)
	cacheinmem.Client.SettOutUser(hashCache, performancesFound)

	c.JSON(http.StatusOK, performancesFound)

}
