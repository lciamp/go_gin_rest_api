# go-gin-rest-api
go-gin-rest-api

```go get .```  to get gin

```go run .``` to start server

```curl localhost:8080/drivers``` result:
```
[
    {
        "id": "1",
        "firstName": "Charles",
        "lastName": "LeClerc",
        "team": "Scuderia Ferrari",
        "number": 16
    },
    {
        "id": "2",
        "firstName": "Carlos",
        "lastName": "Sainz",
        "team": "Scuderia Ferrari",
        "number": 55
    }
]
```

Gin logging:
```[GIN-debug] GET    /drivers                  --> main.getDrivers (3 handlers)
[GIN-debug] Listening and serving HTTP on localhost:8080
[GIN] 2021/10/16 - 22:40:57 | 200 |     141.856µs |       127.0.0.1 | GET      "/drivers"
```