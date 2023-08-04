package find_alerts_usecase

import (
	alerts_repository "alerts/modules/alerts/infra/repository"
	"testing"
	"time"
)

type MockAlertRepository struct {
}

func (r *MockAlertRepository) GetAll() ([]alerts_repository.Alert, error) {

	res := []alerts_repository.Alert{}
	res = append(res, alerts_repository.Alert{
		Id: "dsadsad",
		VehicleId: "",
		Origin: "",
		Type: "",
		Muted: false,
		TriggeredAt: time.Now(),
	})
	// res = append(res, alerts_repository.Alert{
	// 	Id: "sdsda",
	// })
	return res, nil
}

func CreateMockRepository() alerts_repository.IAlertsRepository {
	return &MockAlertRepository{}
}

func TestFindAlertsUseCaseNoFilter(t *testing.T) {
	mockAlertRepository := CreateMockRepository()
	findAlertUsecase := CreateFindAlertsUsecase(mockAlertRepository)

	input := FindAlertUseCaseInputDto{}

	res, _ := findAlertUsecase.Run(input)

	if len(res) != 1 {
		t.Errorf("Expected 1 alert. Got %d", len(res))
	}
}
