package queries

import (
	"fmt"
	"io"
	"strings"
)

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
	if q.Equals != nil {
		b.WriteString(q.buildQueryMatch())
		if q.Range_date != nil || q.Range_price != nil || q.Range_date != nil {
			b.WriteString(",\n")
		}
	}
	if q.Range_date != nil {
		b.WriteString(q.buildQueryDate())
		if q.Range_price != nil || q.Range_date != nil {
			b.WriteString(",\n")
		}
	}
	if q.Range_date != nil {
		b.WriteString(q.buildQueryDate())
		if q.Range_price != nil {
			b.WriteString(",\n")
		}
	}
	if q.Range_price != nil {
		b.WriteString(q.buildQueryPrice())
	}

	b.WriteString("]")
	b.WriteString("}\n")

	if q.Orderby == nil {
		b.WriteString("}\n")
	} else {
		b.WriteString("},\n")
		b.WriteString(q.buildQueryOrderby())
	}

	b.WriteString("}\n")

	//fmt.Printf(b.String())
	return strings.NewReader(b.String())
}

func (q EsQuery) buildQueryMatch() string {
	var b strings.Builder

	/*
					{"match": {
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

func (q EsQuery) buildQueryPrice() string {
	var b strings.Builder

	/*
			   {"nested": {
			          "path": "sections",
			          "query": {
			            "bool": {
			              "must": [
			                {"range": {
			                  "sections.price": {
			                    "gte": 10,
			                    "lte": 200
		}}}]}}}}
	*/
	b.WriteString("{\n")
	b.WriteString(`"nested":{`)
	b.WriteString("\n")
	b.WriteString(`"path": "sections",`)
	b.WriteString("\n")
	b.WriteString(`"query": {`)
	b.WriteString("\n")
	b.WriteString(`"bool": {`)
	b.WriteString("\n")
	b.WriteString(`"must": [`)
	b.WriteString("\n")
	b.WriteString(`{"range": {`)
	b.WriteString("\n")
	b.WriteString(`"sections.price": {`)
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf(`"gte":"%d","lte":"%d"`, q.Range_price.From, q.Range_price.To))

	b.WriteString("}}}]}}}}\n")
	return b.String()
}

func (q EsQuery) buildQueryDate() string {
	var b strings.Builder

	/*
	   {
	   	"range": {
	   	"date": {
	   	  "gte": 1637877540,
	   	  "lte": 1637877620
	   			}
	   		  }
	   }
	*/

	b.WriteString("{\n")
	b.WriteString(`"range": {`)
	b.WriteString("\n")

	b.WriteString(`"date": {`)
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf(`"gte":"%d","lte":"%d"`, q.Range_date.From, q.Range_date.To))
	b.WriteString("}}}\n")

	return b.String()
}

func (q EsQuery) buildQueryOrderby() string {
	var b strings.Builder

	/*
			 "sort" : [
		      {"date" : {"order" : "asc"}},
		      {"showName.keyword" : {"order" : "desc"}}
		   ]
	*/

	b.WriteString(`"sort": [`)
	b.WriteString("\n")

	for index, v := range q.Orderby {
		if v.Field == "date" || v.Field == "performanceID" || v.Field == "theaterID" || v.Field == "showID" {
			b.WriteString(fmt.Sprintf(`{"%s" : {"order" : "%s"`, v.Field, v.Value))
		}
		if v.Field == "price" {
			b.WriteString(buildQueryOrderbyPrice(v))
		} else {
			b.WriteString(fmt.Sprintf(`{"%s.keyword" : {"order" : "%s"`, v.Field, v.Value))
		}

		if index != len(q.Orderby)-1 {
			b.WriteString("}},\n")
		} else {

			b.WriteString("}}\n")
		}
	}

	b.WriteString(`]`)
	b.WriteString("\n")

	return b.String()
}

func buildQueryOrderbyPrice(q FieldValueOrder) string {
	var b strings.Builder
	/*
	   {
	           "sections.price" : {
	              "order" : "asc",
	              "nested": {
	                 "path": "sections"
	              }
	           }
	        },
	*/

	b.WriteString("{\n")
	b.WriteString(`"sections.price" : {`)
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf(`"order":"%s",`, q.Value))
	b.WriteString("\n")
	b.WriteString(`"nested": {`)
	b.WriteString("\n")
	b.WriteString(`"path": "sections"`)
	b.WriteString("\n")
	b.WriteString("}")
	return b.String()

}

func (q EsQuery) BuildQueryID() io.Reader {
	var b strings.Builder
	b.WriteString("\n")
	/*	{"query": {  "match": {
			"performanceID": 1
		  }}}

	*/

	b.WriteString(`{"query": {  "match": {`)
	b.WriteString(fmt.Sprintf(`"%s": %d`, q.Id.Field, q.Id.Value))
	b.WriteString(`}}}`)

	return strings.NewReader(b.String())
}

func (q EsQuery) BuildQueryMatchAll() io.Reader {

	var b strings.Builder
	b.WriteString("\n")
	/*
	   {"query": {"match_all": {}}}
	*/

	b.WriteString(`{"query": {"match_all": {}}}`)

	return strings.NewReader(b.String())
}
