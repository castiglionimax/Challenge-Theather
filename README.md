# MeliShows-Challenge

## Decisiones
- El projecto fue desarrolado en Golang utilizado [Gin](https://github.com/gin-gonic/gin).
- Se persiste los datos en base de datos MongoDB y se utiliza el conector [MongoDB Golang Driver](https://github.com/mongodb/mongo-go-driver).
- Las busquedas se desarrollan en Elasticsearch y se utiliza el conector oficial [go-elasticsearch](https://github.com/elastic/go-elasticsearch).
- Se realizó un in-memory cache con  [TTLCache](https://github.com/ReneKroon/ttlcache).

## Observaciones
- Las busquedas se desarrollan en Elasticsearch y se persiste el documento de funciones en MongoDB.
- Cada reserva se lo toma como una transaccion, por lo tanto si falla la actualización en Elasticserch o en MongoDB la reserva no se hace.
- Se expone un endpoint que brinda mas flexibildad a la busqueda. Dejando la opcion de realizar más busquedas a futuro por nuevos criterios.
- Las busquedas serán guardas en un in-memory cache con un hash en sha1 para posteriormente compararlas y obtener inequivocamente el resultado correcto, el cual tambien está cacheado.

## Listado de Endpoint
- Base URL: MeliShows-Challenge-dev.us-west-1.elasticbeanstalk.com

POST /performaces/search

#### Atributos 
- Para obtener datos entre 1 a 50 hacer:
- https://..../performances/search?limit=25&offset=50
- offset= 1 ->  obtener el primer documento
- limit= 50 ->  cantidad a devolver
#### Atributos Body
- "equals" Array
- field: key con el nombre del atributo en el documento performances 
- Value: valor buscado
- "range_price"
- from": precio inicial
- "to":  precio final
- "orderby": Array. Puede ordenar por mas de un atributo
- "field": key con el nombre del atributo en el documento performances 
- "value":"asc" para ascendente o "des" para descendente 
- "Range_date":  en Timestamp
- "from":1632356338 (seg)
- "to":1637616062
	
#### Ejemplos

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
### Mapping
- Es importante esto porque el array es del tipo nested
```
{
    "mappings" : {
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

## Documentos
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



