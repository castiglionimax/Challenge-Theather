package performance

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/castiglionimax/MeliShows-Challenge/database/elastics"
	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"
)

const (
	Index = "performances"
)

func (p *Performance) Search(query queries.EsQuery) ([]Performance, *errors.RestErr) {
	var r map[string]interface{}

	buf := query.BuildQuery()

	res, err := elastics.Client.Search(buf)

	if err != nil {
		return nil, errors.NewInternalServerError("error when trying to search documents")
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	performances := make([]Performance, int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)))

	for index, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		performance := Performance{}
		arrayByte, _ := json.Marshal(hit.(map[string]interface{})["_source"])
		err := json.Unmarshal(arrayByte, &performance)
		if err != nil {
			log.Print(err)
		}
		//documento full sin _id
		arrayByte, _ = json.Marshal(hit.(map[string]interface{})["_id"])
		err = json.Unmarshal(arrayByte, &performance)
		if err != nil {
			log.Print(err)
		}

		//docuemnto con _id de elastic search
		performance.Date = time.Unix(*performance.DateTimeStamp, 0)
		performance.DateTimeStamp = nil

		//deleting section out of range
		p.ValidatePrice(query)

		fmt.Println(performance)
		performances[index] = performance
	}
	return performances, nil

}

func (p *Performance) Put() *errors.RestErr {

	jsond, _ := json.Marshal(p)
	myString := string(jsond)
	fmt.Println(myString)

	err := elastics.Client.Index(myString, Index, &p.DocumentID)
	if err != nil {
		return errors.NewInternalServerError("error when trying to search documents")
	}

	return nil
}
