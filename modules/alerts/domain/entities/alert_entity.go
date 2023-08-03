package alert_entity

import "time"

type Alert struct {
	Id          string    `json:"id"`
	VehicleId   string    `json:"vehicle_id"`
	Origin      string    `json:"origin"`
	Type        string    `json:"type"`
	Muted       bool      `json:"muted"`
	TriggeredAt time.Time `json:"triggered_at"`
}

// tensao da bateria < 15
