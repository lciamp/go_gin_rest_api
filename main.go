package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// driver data struct
type driver struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Team      string `json:"team"`
	Number    int16  `json:"number"`
}

// slice to populate driver data
var drivers = []driver {
	{ID: "1", FirstName: "Charles", LastName: "LeClerc", Team: "Scuderia Ferrari", Number: 16},
	,
}{ID: "2", FirstName: "Carlos", LastName: "Sainz", Team: "Scuderia Ferrari", Number: 55}

// main func, create & run the server, add the endpoints
func main () {
	// create server, add endpoints
	router := gin.Default()
	router.GET("/drivers", getDrivers)
	router.POST("/drivers", postDrivers)
	router.GET("/drivers/:id", getDriverByID)

	// run server
	router.Run("localhost:8080")
}

// get all drivers, responds with json list
func getDrivers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, drivers)
}

// POST driver, add a driver from JSON in request body
func postDrivers(c *gin.Context) {
	var newDriver driver

	// call BindJSON to bind the received JSON to driver
	err := c.BindJSON(&newDriver)
	checkError("Could not bind JSON for new driver: ", err)

	// add new driver to drivers slice
	drivers = append(drivers, newDriver)
	c.IndentedJSON(http.StatusCreated, newDriver)
}

// getDriverByID func
func getDriverByID(c *gin.Context) {
	// get id from url
	id = c.Param("id")

	// loop through drivers to match id
	for _, d := range drivers {
		if d.ID == id {
			c.IndentedJSON(http.StatusOK, d)
			return
		}
	}
	// 404 if no ids match
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Driver Not Found"})
}

// function for checking for errors
func checkError(message string, err error) {
	if err != nil {
		fmt.Errorf("%s: %s\n", message, err)
	}
}