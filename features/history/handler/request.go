package handler

import "project-capston/features/history"

type HistoryRequest struct {
	UnitID   uint   `json:"unit_id" form:"unit_id"`
	DriverID uint   `json:"driver_id" form:"driver_id"`
	Status   string `json:"status" form:"status"`
	Reason   string `json:"reason" form:"reason"`
}

func RequestToEntity(data HistoryRequest) history.HistoryEntity {
	return history.HistoryEntity{
		UnitID:   data.UnitID,
		DriverID: data.DriverID,
		Status:   data.Status,
		Reason:   data.Reason,
	}
}
