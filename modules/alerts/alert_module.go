// responsavel por receber e injetar as dependencias relacionadas à alertas
package alerts

import (
	alerts_application "alerts/modules/alerts/application"
	alerts_controller "alerts/modules/alerts/application/controller"
	find_alerts_usecase "alerts/modules/alerts/domain/usecases/find"
	alerts_repository "alerts/modules/alerts/infra/repository"
	"fmt"
)

type AlertModule struct{}

func (m *AlertModule) Initialize() error {

	alertsRepository := alerts_repository.CreateAlertRepository()

	findAlertsUseCase := find_alerts_usecase.CreateFindAlertsUsecase(alertsRepository)

	alertController := alerts_controller.CreateAlertController(findAlertsUseCase)

	alerts_application.CreateAlertHttpRouter(alertController)

	fmt.Println("AlertModule initialized!")
	return nil
}

func CreateModule() IAlertModuler {
	return &AlertModule{}
}
