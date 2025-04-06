package cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	"wallet/helpers"
	"wallet/internal/api"
	"wallet/internal/services"
)

func bootstrapEngine() {

}

func ServeHTTP() {

	healthCheckServices := &services.HealthCheck{}
	healthCheckApi := api.HealthCheck{
		HealthCheckServices: healthCheckServices,
	}

	r := gin.Default()
	r.GET("/health", healthCheckApi.HealthCheckHandlerHTTP)

	err := r.Run(":" + helpers.GetEnv("HTTP_PORT", "8080"))
	if err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
