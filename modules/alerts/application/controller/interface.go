package alerts_controller

import (
	"net/http"
)

type IAlertController interface {
	FindAlerts(w http.ResponseWriter, r *http.Request)
}
