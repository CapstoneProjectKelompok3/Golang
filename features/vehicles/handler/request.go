package handler

import "project-capston/features/vehicles"

type VehicleRequest struct {
	GovermentID uint   `json:"goverment_id" form:"goverment_id"`
	Plate       string `json:"plate" form:"plate"`
	Status      bool   `json:"status" form:"status"`
}

func RequestToEntity(vehicle VehicleRequest) vehicles.VehicleEntity{
	return vehicles.VehicleEntity{
		GovermentID: vehicle.GovermentID,
		Plate:       vehicle.Plate,
		Status:      vehicle.Status,
	}
}