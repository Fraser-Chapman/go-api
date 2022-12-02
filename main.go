package main

import (
    "fraser-chapman/go-api/error"
    "fraser-chapman/go-api/weather"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

    router.GET("/weather", func(c *gin.Context) {
        latitude, err := strconv.ParseFloat(c.Query("latitude"), 64)
        error.Handle(err)

        longitude, err := strconv.ParseFloat(c.Query("longitude"), 64)
        error.Handle(err)

        c.IndentedJSON(http.StatusOK, weather.Get(latitude, longitude))
    })

	router.Run("localhost:8080")
}
