package alerts_repository

type AlertsRepository struct {
}

// Delete implements IAlertsRepository
func (*AlertsRepository) Delete() {
	panic("unimplemented")
}

// GetAll implements IAlertsRepository
func (r *AlertsRepository) GetAll() ([]Alert, error) {
	panic("unimplemented")
}

// GetById implements IAlertsRepository
func (*AlertsRepository) GetById(id string) (Alert, error) {
	panic("unimplemented")
}

// GetByPlate implements IAlertsRepository
func (*AlertsRepository) GetByPlate() {
	panic("unimplemented")
}

// GetByVehicleId implements IAlertsRepository
func (*AlertsRepository) GetByVehicleId() {
	panic("unimplemented")
}

// Insert implements IAlertsRepository
func (*AlertsRepository) Insert() {
	panic("unimplemented")
}

// Update implements IAlertsRepository
func (*AlertsRepository) Update() {
	panic("unimplemented")
}

func CreateAlertRepository() IAlertsRepository {
	return &AlertsRepository{}
}
