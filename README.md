# MeliShows-Challenge

## Decisiones
- El projecto fue desarrolado en Golang utilizado [Gin](https://github.com/gin-gonic/gin).
- Se persiste los datos en base de datos MongoDB y se utiliza el conector [MongoDB Golang Driver](https://github.com/mongodb/mongo-go-driver).
- Las busquedas se desarrollan en Elasticsearch y se utiliza el conector oficial [go-elasticsearch] (https://github.com/elastic/go-elasticsearch).
- El projecto esta hosteado en una privete subnet de AWS en un nodo de EC(todos son singles nodo en docker), por lo tanto el entorno esta preparado para escalar pero esta opción no está habilitada.

## Observaciones
- Las busquedas se desarrollan en Elasticsearch y se persiste el documento de funciones tanto en Mongo como en elasticsearch.
- Cada reserva se lo toma como una transaccion, por lo tanto si falla la actualización en Elasticserch o en Mongo la reserva no se hace.
- Las busquedas realizan haciendo un metodo POST y creado un body en json para darle mas flexibildad a la busqueda. Dejando la opcion de realizar mas busquedas a nuevos criterios a futuro.

## List of Endpoints
- Base URL: http://xxxxx:50000

Path: /performaces/search
Rest verb: POST

### POST Busquedas
#### Atributos URL
- Para obtener datos entre 1 a 50 hacer:
- - https://xxxxx:50000/performances/search?limit=25&offset=50
- - offset= 1 significa obtener el 1 documento
- - limit= 10 la cantidad a devolver
#### Atributos Body
- "equals" Array
- - field: key con el nombre del atributo en el documento performances 
- - Value: valor buscado
-	"range_price"
- - from": precio inicial
- - "to":  precio final
- "orderby": Array. Puede ordenar por mas de un atributo
- -  "field": key con el nombre del atributo en el documento performances 
- -  "value":"asc" para ascendente o "des" para descendente 
- 	"Range_date":  en Timestamp
		"from":1632356338 (seg)
    "to":1637616062
	
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
    ```
    {
        "message": "error when trying to search documents",
        "error": "internal_server",
        "status": 500
    }
    ```
    - 404: StatusNotFound 
    ```
    {
      
           "message": "error when trying to search documents",
            "error": "Not_Found_Error",
            "status": 404

    }
    ```
    -  400:	StatusBadRequest  

    ```
  {
      
        "message": "error when trying to search documents",
        "error": "bad_request",
        "status": 400

    }
  ```
    - 404: StatusNotFound 
    ```
    {
        "message": "error, seat %d in the section %d is not found"
        "error": "Not_Found_Error",
        "status": 404
    }
    ```
    -  400:	StatusBadRequest  

    ```
  {
      
        "message": invalid json body",
        "error": "bad_request",
        "status": 400

    }
    ```
### Bookings
Path: /bookings
Rest verb: POST


#### Ejemplos

- Body
    ```

{
    "performanceID":1,
    "person":{"dni":32523291, "fullname": "Alexis Castiglioni"},
    "sold":[{"seat":3, "sectionId":1},{"seat":3,"sectionId":3}]
}
    ```
- Responses
    ```
  - 200: Created
{
    "performanceID": 1,
    "person": {
        "dni": 32523291,
        "fullname": "Alexis Castiglioni"
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
    ```
    -  400:	StatusBadRequest  

    ```
  {
      
        "message": "invalid request body",
        "error": "bad_request",
        "status": 400

    }
  ```
      ```
    -  400:	StatusBadRequest  

    ```
  {
      
        "message": "invalid json body",
        "error": "bad_request",
        "status": 400

    }
  ```

    ```
      - 500: StatusInternalServerError
    ```
    {
        "message": "error when trying to search documents",
        "error": "internal_server",
        "status": 500
    }
    ```
    - 404: StatusNotFound 
    ```
    {
      
           "message": "error when trying to search documents",
            "error": "Not_Found_Error",
            "status": 404

    }
    ```

  
    -  400:	StatusBadRequest  

    ```
  {
      
        "message": invalid json body",
        "error": "bad_request",
        "status": 400

    }
    ```
