package handler

type UnitResponse struct {
	EmergenciesID uint `json:"emergencies_id,omitempty"`
	VehicleID     uint `json:"vehicle_id,omitempty"`
}

type UnitHistoryResponse struct {
	EmergenciesID   uint   `json:"emergencies_id,omitempty"`
	VehicleID       uint   `json:"vehicle_id,omitempty"`
	Status          string `json:"status,omitempty"`
	AlasanPenolakan string `json:"alasan_penolakan,omitempty"`
}