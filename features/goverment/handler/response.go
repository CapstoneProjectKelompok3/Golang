package handler

import "project-capston/features/goverment"

type GovernmentResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GovernmentNearestResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	// Type      string  `json:"type"`
	// Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance"`
}

type UnitCountR struct {
	RumahSakit int64 `json:"unit_rumah_sakit"`
	Pemadam    int64 `json:"unit_pemadam"`
	Kepolisian int64 `json:"unit_kepolisian"`
	SAR        int64 `json:"unit_SAR"`
	Dishub     int64 `json:"unit_dishub"`
}

func MappingCountUnit(unit goverment.UnitCount)UnitCountR{
	return UnitCountR{
		RumahSakit: unit.RumahSakit,
		Pemadam:    unit.Pemadam,
		Kepolisian: unit.Kepolisian,
		SAR:        unit.SAR,
		Dishub:     unit.Dishub,
	}
}