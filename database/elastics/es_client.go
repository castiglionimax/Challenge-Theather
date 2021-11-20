package elastics

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/castiglionimax/MeliShows-Challenge/domain/queries"
	"github.com/castiglionimax/MeliShows-Challenge/utils/errors"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

//export ELASTIC_HOSTS=asasdaee
const (
	envEsHosts = "ELASTIC_HOSTS"
	Index      = "performances"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(client *elasticsearch.Client)
	//Get(index string, docType string, id string) (*elastic.GetResult, error)
	//Search(string, elastic.Query) (*elastic.SearchResult, error)
	Search(queryIn queries.EsQuery) (*esapi.Response, *errors.RestErr)
}

type esClient struct {
	es *elasticsearch.Client
}

func Init() {
	var (
		r   map[string]interface{}
		err error
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
	Client.setClient(es)

}

func (c *esClient) setClient(client *elasticsearch.Client) {
	c.es = client
}

func (c *esClient) Search(queryIn queries.EsQuery) (*esapi.Response, *errors.RestErr) {

	/*
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
	*/
	// Perform the search request.
	//buf := queryIn.Build()
	buf := queryIn.BuildQuery()

	res, err := c.es.Search(
		c.es.Search.WithContext(context.Background()),
		c.es.Search.WithIndex(Index),
		c.es.Search.WithBody(buf),
		c.es.Search.WithTrackTotalHits(true),
		c.es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	return res, nil
}
