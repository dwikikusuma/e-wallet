package interfaces

type HealthCheckServices interface {
	HealthCheckServices() (string, error)
}

type HealthCheckRepository interface {
}
