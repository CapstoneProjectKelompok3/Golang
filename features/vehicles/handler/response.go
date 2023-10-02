package handler

import "project-capston/features/vehicles"

type VehicleResponse struct {
	Id          uint   `json:"id,omitempty"`
	GovermentID uint   `json:"goverment_id,omitempty"`
	Plate       string `json:"plate,omitempty"`
	Status      bool   `json:"status,omitempty"`
}

func EntityToResponse(vehicle vehicles.VehicleEntity)VehicleResponse{
	return VehicleResponse{
		Id:          vehicle.Id,
		GovermentID: vehicle.GovermentID,
		Plate:       vehicle.Plate,
		Status:      vehicle.Status,
	}
}