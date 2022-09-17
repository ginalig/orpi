package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET(":message", get)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func get(c *gin.Context) {
	message := c.Param("message")
	c.String(http.StatusOK, message)
}
