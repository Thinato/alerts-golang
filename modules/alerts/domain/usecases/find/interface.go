package find_alerts_usecase

import (
	alert_entity "alerts/modules/alerts/domain/entities"
)

type IFindAlertsUsecase interface {
	Run(input FindAlertUseCaseInputDto) ([]alert_entity.Alert, error)
}
