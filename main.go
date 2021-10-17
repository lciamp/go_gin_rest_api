package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// driver data struct
type driver struct{
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Team      string `json:"team"`
	Number    int16  `json:"number"`
}

// slice to populate driver data
var drivers = []driver{
	{ID: "1", FirstName: "Charles", LastName: "LeClerc", Team: "Scuderia Ferrari", Number: 16},
	{ID: "2", FirstName: "Carlos", LastName: "Sainz", Team: "Scuderia Ferrari", Number: 55},
}

// get all drivers, responds with json list
func getDrivers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, drivers)
}