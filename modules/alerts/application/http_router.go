package alerts_application

import (
	alerts_controller "alerts/modules/alerts/application/controller"
	"fmt"
	"net/http"
)

type AlertHttpRouter struct {
	alertController alerts_controller.IAlertController
}

func (router *AlertHttpRouter) Initialize() {
	http.HandleFunc("/alerts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("sasasasasas")
	})
	return
}

func CreateAlertHttpRouter(alertController alerts_controller.IAlertController) IAlertHttpRouter {
	return &AlertHttpRouter{
		alertController: alertController,
	}
}
