package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()

    router.GET("/hello", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello world!")
    })

    router.Run("localhost:8080")
}