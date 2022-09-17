package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Service) Get(c *gin.Context) {
	c.String(http.StatusOK, "kek")
}
