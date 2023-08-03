package find_alerts_usecase

import "time"

type FindAlertUseCaseInputDto struct {
	Muted           bool      `json:"muted"`
	Origin          string    `json:"origin"`
	Type            string    `json:"type"`
	TriggeredAtFrom time.Time `json:"triggered_at_from"`
	TriggeredAtTill time.Time `json:"triggered_at_till"`
	VehicleId       string    `json:"vehicle_id"`
	Plate           string    `json:"plate"`
}
