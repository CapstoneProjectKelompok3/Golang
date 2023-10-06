package handler

import "project-capston/features/unit"

type UnitRequest struct {
	EmergenciesID uint   `json:"emergencies_id" form:"emergencies_id"`
	GovermentType string `json:"goverment_type" form:"goverment_type"`
	SumOfUnit     int    `json:"sum_of_unit" form:"sum_of_unit"`
}

type UnitHistoryRequest struct {
	EmergenciesID   uint   `json:"emergencies_id" form:"emergencies_id"`
	VehicleID       uint   `json:"vehicle_id" form:"vehicle_id"`
	Status          string `json:"status" form:"status"`
	AlasanPenolakan string `json:"alasan_penolakan" form:"alasan_penolakan"`
}

func RequestToEntity(data UnitRequest) unit.UnitEntity {
	return unit.UnitEntity{
		EmergenciesID: data.EmergenciesID,
		GovermentType: data.GovermentType,
		SumOfUnit:     data.SumOfUnit,
	}
}
