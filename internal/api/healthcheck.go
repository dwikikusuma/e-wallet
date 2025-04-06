package api

import (
	"github.com/gin-gonic/gin"
	"wallet/helpers"
	"wallet/internal/interfaces"
)

type HealthCheck struct {
	HealthCheckServices interfaces.HealthCheckServices
}

func (h *HealthCheck) HealthCheckHandlerHTTP(c *gin.Context) {
	msg, err := h.HealthCheckServices.HealthCheckServices()
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": err.Error()})
		return
	}
	helpers.SendResponseHTTP(c, 200, msg, nil)
}
