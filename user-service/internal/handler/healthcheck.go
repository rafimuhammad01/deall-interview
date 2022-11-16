package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (h *HealthCheck) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"test": "success",
	})
}
