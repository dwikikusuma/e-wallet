package services

import "wallet/internal/interfaces"

type HealthCheck struct {
	HealthCheckRepository interfaces.HealthCheckRepository
}

func (s *HealthCheck) HealthCheckServices() (string, error) {
	return "OK", nil
}
