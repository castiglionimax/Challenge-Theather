package performance

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/castiglionimax/MeliShows-Challenge/database/elastics"
	"github.com/castiglionimax/MeliShows-Challenge/domain/pagination"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"
)

const (
	Index = "performances"
)

func (p *Performance) Search(query io.Reader, pagination *pagination.Pagination) ([]Performance, *errors.RestErr) {
	var r map[string]interface{}

	res, err := elastics.Client.Search(query, pagination)

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

	if len(performances) == 0 {
		return nil, errors.NewNotFoundError("document not found")
	}

	for index, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		performance := Performance{}

		arrayByte, _ := json.Marshal(hit.(map[string]interface{})["_source"])

		err := json.Unmarshal(arrayByte, &performance)
		if err != nil {
			log.Print(err)
		}
		//documento full sin _id

		j := fmt.Sprintf("%s", hit.(map[string]interface{})["_id"])
		log.Printf("_id=%s", j)
		performance.DocumentID = &j

		//docuemnto con _id de elastic search

		performances[index] = performance
	}
	return performances, nil

}

func (p *Performance) Put() *errors.RestErr {

	Id := *p.DocumentID
	p.DocumentID = nil
	jsond, _ := json.Marshal(p)
	myString := string(jsond)
	//fmt.Println(myString)

	err := elastics.Client.Index(myString, Index, &Id)
	if err != nil {
		return errors.NewInternalServerError("error when trying to search documents")
	}

	return nil
}
