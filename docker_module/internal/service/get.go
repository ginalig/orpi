package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) Get(c *gin.Context) {
	message := c.Param("message")
	c.String(http.StatusOK, message)
}
