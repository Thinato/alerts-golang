package find_alerts_usecase

import (
	alert_entity "alerts/modules/alerts/domain/entities"
	alerts_repository "alerts/modules/alerts/infra/repository"
)

type FindAlertsUsecase struct {
	repository alerts_repository.IAlertsRepository
}

func (usecase *FindAlertsUsecase) Run(input FindAlertUseCaseInputDto) ([]alert_entity.Alert, error) {

	var res []alert_entity.Alert

	repoAlerts, err := usecase.repository.GetAll()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(repoAlerts); i++ {
		res[i] = alert_entity.Alert{
			Id:        repoAlerts[i].Id,
			VehicleId: repoAlerts[i].VehicleId,
		}
	}

	return res, nil
}

func CreateFindAlertsUsecase(
	repository alerts_repository.IAlertsRepository,
) IFindAlertsUsecase {
	return &FindAlertsUsecase{
		repository: repository,
	}
}
