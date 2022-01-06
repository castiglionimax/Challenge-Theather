# Challenge
# Overview

This project was made as a challenge to the company that I applied to.
I’ve kept the company name anonymous in order to protect confidential information. So, the following case is not the real one but an appreciation of it.
# Introduction

The goal is to build an API that contains all RESTful fundamentals according to the RFC 2616 HTTP/1.1 spec.
The API is about managing theaters and booking different performances
The API has two type of users:

1. Who add data to the platform like new theaters, seats or even rooms
2. From the front page, it’s used to book performance for customers.

Note: A performance might have different prices for one or more sections. So, the same performance might have different prices.

The portal must have the follow functionality:

- Getting all shows and performances available.
- Getting all seats and prices available for a specific performance.
- Booking a performance by passport ID, first name and last name as a reserve.

Note: It’s important to manage the use of cases when someone book a performance from who book the same performance, the API mustn’t allow it.
Record data into the database might be omitted. So, It’s not necessary to have a public endpoint API with that functionality.

Main goals were evaluated:

- Add a functionality for search performances by dates, prices and the results should be ordered by descent as default but with the option of ascending according to some attribute.
- Implement some functionality to improve the search base of most recently.
- Host the API into a cloud provider (Google App Engine, Amazon AWS, Azure).

## Decisions I took
- The code was written in Golang using GIN [Gin](https://github.com/gin-gonic/gin).
- The database selected was  MongoDB and the connector for Go was [MongoDB Golang Driver](https://github.com/mongodb/mongo-go-driver).
- The functionality of searching was developed using Elasticseach and the connector for Golang was the official one. [go-elasticsearch](https://github.com/elastic/go-elasticsearch).
-  In-memory cache was used to improve searches [TTLCache](https://github.com/ReneKroon/ttlcache).

## To considered
- Even though every single search showed performance using elaticsearch, the documents were saved on MondoDB as well.
- When a customer wants to book a show, the functionality works as a transaction so if one is false the whole transaction is aborted.
- The endpoint for search performance published had the possibility to manage enhanced search by attribute.
- Even though I knew implementing in-memory cache was not the right call, I did it to accomplish that requirement (improve searching). If I had more time, I would have developed a distributed cache.


## Endpoint List
- Base URL: XXXXX-Challenge-dev.us-west-1.elasticbeanstalk.com

POST /performaces/search

#### Attributes  
- equals" Array
- field: name of the attribute that you want to get
- Value: value
- "range_price"
- from": initial price
- "to": final price
- "orderby": Array. It could get more than one attribute
- "field": name of the attribute that you want to get
- "value":"asc" to ascended  o "des" to descended
- "Range_date": Timestamp
- "from":1632356338 (seg)
- "to":1637616062
	
#### Examples

- Body
```
{
	"equals":[
		{
		"field":"showName",
		"value":"Aladdin"
		},
		{
		"field":"city",
		"value":"New York"
		}
    	],
	"range_price":{
	    	"from":1,
        	"to":300
	},
      	"orderby":[
        	{
	         "field":"price",
             	 "value":"asc"
	      	}
	    ],
	"Range_date":{
		 "from":1632356338,
		 "to":1637616062
	}	
}

```

#### Errores
```
- Responses:
  - 500: StatusInternalServerError
    {
        "message": "error when trying to search documents",
        "error": "internal_server",
        "status": 500
    }
    
    - 404: StatusNotFound 
    {
           "message": "error when trying to search documents",
            "error": "Not_Found_Error",
            "status": 404
    }
    
    -  400:	StatusBadRequest  

  {
        "message": "error when trying to search documents",
        "error": "bad_request",
        "status": 400

    }
    - 404: StatusNotFound 
    {
        "message": "error, seat %d in the section %d is not found"
        "error": "Not_Found_Error",
        "status": 404
    }
    -  400:	StatusBadRequest  

  {
      
        "message": invalid json body",
        "error": "bad_request",
        "status": 400

    }
  ```
### Bookings

POST /bookings

#### Ejemplos

- Body
```
{
    "performanceID":1,
    "person":{"dni":31233231, "fullname": "Alfred Molina"},
    "sold":[{"seat":3, "sectionId":1},{"seat":3,"sectionId":3}]
}
```

- Responses
    
    ```  
  - 200: Created
  
  {
    "performanceID": 1,
    "person": {
        "dni": 31113821,
        "fullname": "Alfrey Molini"
    },
    "sold": [
        {
            "seat": 5,
            "sectionID": 1,
            "price": 300
        },
        {
            "seat": 5,
            "sectionID": 3,
            "price": 110
        }
    ],
    "total_price": 410
    }

    -  400:	StatusBadRequest  

  {
        "message": "invalid request body",
        "error": "bad_request",
        "status": 400
    }
      
    -  400:	StatusBadRequest  

  {   
        "message": "invalid json body",
        "error": "bad_request",
        "status": 400
    }

    
      - 500: StatusInternalServerError
    {
        "message": "error when trying to search documents",
        "error": "internal_server",
        "status": 500
    }
    
    - 404: StatusNotFound 
    
    {  
           "message": "error when trying to search documents",
            "error": "Not_Found_Error",
            "status": 404
    }
    
    -  400:	StatusBadRequest  

  {
        "message": invalid json body",
        "error": "bad_request",
        "status": 400
    }
    ```  

## Elasticsearch

- Es importante esto porque el array es del tipo nested

    ```
    - Mapping
    {"mappings" :{
    "properties" : {
        "auditorium" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "city" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "date" : {
          "type" : "long"
        },
        "dateShow" : {
          "type" : "date"
        },
        "performanceID" : {
          "type" : "long"
        },
        "sections" : {
          "type" : "nested",
          "properties" : {
            "currency" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "description" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "id" : {
              "type" : "long"
            },
            "name" : {
              "type" : "text",
              "fields" : {
                "keyword" : {
                  "type" : "keyword",
                  "ignore_above" : 256
                }
              }
            },
            "price" : {
              "type" : "long"
            },
            "seats" : {
              "type" : "long"
            }
          }
        },
        "showID" : {
          "type" : "long"
        },
        "showName" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        },
        "theaterID" : {
          "type" : "long"
        },
        "theaterName" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        }
      }
    }
  }
    ```


## Documents
### Peformance EXAMPLES

```
{
     "performanceID":1,
    "showID":1,
    "showName": "Aladdin",
    "theaterID": 1,
    "theaterName": "Richard Rodgers Theatre",
    "city": "New York",
    "auditorium":"New York",
    "sections": [
        { "id":1,"name": "OrchestraA","description":"Row A", "seats": [1,2,3,4,5,6,7,8,9,10],"price":300.00,"currency": "USD"},
        { "id":2,"name": "Mezzanine","description":"General", "seats": [1,2,3,4,5,6,7,8,9,10],"price":110.00,"currency": "USD"},
        { "id":3,"name": "Balcony","description":"General", "seats": [1,2,3,4,5,6,7,8,9,10],"price":50.00,"currency": "USD"}
    ],
    "date": 1636581600
}
```



