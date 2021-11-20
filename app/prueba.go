package app

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/castiglionimax/MeliShows-Challenge/domain/performance"
	"github.com/elastic/go-elasticsearch/v8"
)

func ElasticSearchPrueba() {

	var (
		r map[string]interface{}
	)

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Printf("Client: %s", elasticsearch.Version)
	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
	log.Println(strings.Repeat("~", 37))

	// 3. Search for the indexed documents
	//
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"city": "New York",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("performances"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.

	//	performances := make([]performance.Performance, 0)

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		//	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		prueba := hit.(map[string]interface{})["_source"]
		fmt.Print(prueba)

		performance := performance.Performance{}

		arrayByte, _ := json.Marshal(hit.(map[string]interface{})["_source"])
		err := json.Unmarshal(arrayByte, &performance)
		if err != nil {
			log.Print(err)
		}

		/*	performances=append(performances,performance.Performance{
			PerformanceID int64     `json:"performanceID" bson:"performanceID"`
			ShowID        int64     `json:"showID" bson:"showID"`
			ShowName      string    `json:"showName" bson:"showName"`
			TheaterID     int64     `json:"theaterID" bson:"theaterID"`
			City          string    `json:"city" `
			TheaterName   string    `json:"theaterName" bson:"theaterName"`
			Auditorium    string    `json:"auditorium" bson:"auditorium"`
			Sections      []Section `json:"sections" bson:"sections"`
			Date          time.Time `json:"date" bson:"date"`

		})*/
	}

	log.Println(strings.Repeat("=", 37))
}
