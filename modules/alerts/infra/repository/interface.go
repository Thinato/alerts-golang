package alerts_repository

type IAlertsRepository interface {
	// GetById(id string) (Alert, error)
	GetAll() ([]Alert, error)
}
