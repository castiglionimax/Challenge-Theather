package queries

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

const (
	searchMatch = `
"query": {
  "bool": {
	"must": [
		  {"match": 
			{"showName": "Aladdin"}
		  }]
	  }
  }
}`
	searchFull = `
	"query": {
		"bool": {
		  "must": [
				{"match": 
				  {"showName": "Aladdin"}
				},
				{"range": {
					"date": {
					  "gte": 1637877540,
					  "lte": 1637877620
							}     
						  }
				},
	  
				{"nested": {
					"path": "sections",
					"query": {
					  "bool": {
						"must": [
						  {"range": {
							"sections.price": {
							  "gte": 10,
							  "lte": 200
							}
						  }}
						]
					  }
					}
				  }}
		  
				
				]
			}
		}
	  }`
)

func (q EsQuery) Build() bytes.Buffer {

	var buf bytes.Buffer
	var query map[string]interface{}
	//var match []map[string]interface{}

	mp2 := map[string]interface{}{
		"alexis": 3,
	}

	//if q.Equals != nil {
	//			match = q.getEquals()
	//	}

	query = map[string]interface{}{
		"query": mp2,
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	return buf
}

/*
GET performances/_search
{
  "query": {
    "bool": {
      "must": [
        {"fuzzy": {
          "showName": "Aladdin"
        }},
        {"nested": {
          "path": "sections",
          "query": {
            "bool": {
              "must": [
                {"range": {
                  "sections.price": {
                    "gte": 10,
                    "lte": 200
                  }
                }}
              ]
            }
          }
        }},
          {"range": {
          "date": {
            "gte": 1637877540,
            "lte": 1637877620
          }
        }
        }
      ],
      "boost": 1.0
  }
  }
  }
*/
/*


func (q EsQuery) Build() bytes.Buffer {

	var buf bytes.Buffer
	var query map[string]interface{}

	for _, eq := range q.Equals {

		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					eq.Field: eq.Value,
				},
			},
		}
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	return buf
}
*/
func (q EsQuery) BuildQuery() io.Reader {
	var b strings.Builder
	/*
		"query": {
			"bool": {
			  "must": [
	*/
	b.WriteString("{\n")
	b.WriteString(`"query": {`)
	b.WriteString("\n")
	b.WriteString(`"bool": {`)
	b.WriteString("\n")
	b.WriteString(`"must": [`)
	b.WriteString("\n")
	b.WriteString(q.BuildQueryFuzzy())
	b.WriteString("]\n")
	b.WriteString("}\n")
	b.WriteString("}\n")
	b.WriteString("}\n")

	fmt.Printf(b.String())
	return strings.NewReader(b.String())
}

func (q EsQuery) BuildQueryFuzzy() string {
	var b strings.Builder

	/*
					{"fuzzy": {
		          "showName": "Aladdin"
		        }},
	*/

	for index, v := range q.Equals {
		b.WriteString("{\n")
		b.WriteString(`"match":{`)
		b.WriteString("\n")
		b.WriteString(fmt.Sprintf(`"%s": "%s"`, v.Field, v.Value))
		if index != len(q.Equals)-1 {
			b.WriteString("}},\n")
		} else {
			b.WriteString("}}\n")

		}
	}

	return b.String()
}
