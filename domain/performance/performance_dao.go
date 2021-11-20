package performance

import (
	"encoding/json"
	"log"
	"time"

	"github.com/castiglionimax/MeliShows-Challenge/database/elastics"
	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"
)

/*
func (p *Performance) GetAllbyShowName() ([]*Performance, error) {

	var performanceCollections []*Performance
	client := mongodb.GetSession()

	meliCollection := client.Database("MeliShows")
	//	performacesCollection := meliCollection.Collection("performace")

	performacesCollection := meliCollection.Collection("performace")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := performacesCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var perforquery Performance
		if err = cursor.Decode(&perforquery); err != nil {
			log.Fatal(err)
		}
		performanceCollections = append(performanceCollections, &perforquery)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	return performanceCollections, nil
}

//{createdAt:{$gte:ISODate("2021-01-01"),$lt:ISODate("2022-05-01"}}
*/
func (p *Performance) Search(query queries.EsQuery) ([]Performance, *errors.RestErr) {
	var r map[string]interface{}
	res, err := elastics.Client.Search(query)

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
		performance.Date = time.Unix(*performance.DateTimeStamp, 0)
		performance.DateTimeStamp = nil

		performances[index] = performance
	}
	return performances, nil

}
