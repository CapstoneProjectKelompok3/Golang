package handler

type UnitRequest struct {
	EmergenciesID uint `json:"emergencies_id" form:"emergencies_id"`
	VehicleID     uint `json:"vehicle_id" form:"vehicle_id"`
}

type UnitHistoryRequest struct {
	EmergenciesID   uint   `json:"emergencies_id" form:"emergencies_id"`
	VehicleID       uint   `json:"vehicle_id" form:"vehicle_id"`
	Status          string `json:"status" form:"status"`
	AlasanPenolakan string `json:"alasan_penolakan" form:"alasan_penolakan"`
}