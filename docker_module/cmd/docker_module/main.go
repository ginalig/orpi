package main

import (
	"docker_module/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	s := service.Service{}
	router.GET(":message", s.Get)

	err := router.Run("localhost:3080")
	if err != nil {
		return
	}
}
