package alerts_controller

import (
	find_alerts_usecase "alerts/modules/alerts/domain/usecases/find"
	"fmt"
	"net/http"
)

type AlertController struct {
	findAlertsUsecase find_alerts_usecase.IFindAlertsUsecase
}

func (controller *AlertController) FindAlerts(w http.ResponseWriter, r *http.Request) {

	inputs := find_alerts_usecase.FindAlertUseCaseInputDto{}

	_, err := controller.findAlertsUsecase.Run(inputs)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Welcome to the HomePage!")
	fmt.Fprintf(w, "Welcome to the HomePage!")

}

func CreateAlertController(
	findAlertsUseCase find_alerts_usecase.IFindAlertsUsecase,
) IAlertController {
	return &AlertController{
		findAlertsUsecase: findAlertsUseCase,
	}
}
