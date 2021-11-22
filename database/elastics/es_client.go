package elastics

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

//export ELASTIC_HOSTS=asasdaee
const (
	envEsHosts      = "ELASTIC_HOSTS"
	IndexPerfomance = "performances"
	IndexBooking    = "bookings"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(client *elasticsearch.Client)
	Index(b string, indexInput string, DocumentIDinput *string) error
	Search(buf io.Reader) (*esapi.Response, error)
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

func (c *esClient) Search(buf io.Reader) (*esapi.Response, error) {

	res, err := c.es.Search(
		c.es.Search.WithContext(context.Background()),
		c.es.Search.WithIndex(IndexPerfomance),
		c.es.Search.WithBody(buf),
		c.es.Search.WithTrackTotalHits(true),
		c.es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, err
	}

	return res, nil
}

func (c *esClient) Index(b string, indexInput string, DocumentIDinput *string) error {

	req := esapi.IndexRequest{}
	if DocumentIDinput == nil {
		// Set up the request object.
		req = esapi.IndexRequest{
			Index:   indexInput,
			Body:    strings.NewReader(b),
			Refresh: "true",
		}
	} else {
		// Set up the request object.
		req = esapi.IndexRequest{
			Index:      indexInput,
			DocumentID: *DocumentIDinput,
			Body:       strings.NewReader(b),
			Refresh:    "true",
		}
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), c.es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return err
	}
	defer res.Body.Close()

	return nil
}
