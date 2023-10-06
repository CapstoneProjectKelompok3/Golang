package handler

import "project-capston/features/vehicles"

type VehicleResponse struct {
	Id          uint   `json:"id,omitempty"`
	GovermentID uint   `json:"goverment_id,omitempty"`
	Plate       string `json:"plate,omitempty"`
	Status      bool   `json:"status,omitempty"`
	Goverment  GovermentResponse `json:"goverment,omitempty"`
}

type GovermentResponse struct{
	Id uint `json:"id,omitempty"`
	NameGoverment string `json:"name_goverment"`
	Address string `json:"address"`
	Type string `json:"type"`
	Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

func EntityToResponse(vehicle vehicles.VehicleEntity)VehicleResponse{
	return VehicleResponse{
		Id:            vehicle.Id,
		GovermentID:   vehicle.GovermentID,
		Plate:         vehicle.Plate,
		Status:        vehicle.Status,
		Goverment: GovermentRresponse(vehicle.Goverment),
	}
}
func GovermentRresponse(goverment vehicles.GovernmentEntity)GovermentResponse{
	return GovermentResponse{
		Id:            goverment.ID,
		NameGoverment: goverment.Name,
		Address:       goverment.Address,
		Type:          goverment.Type,
		Latitude:      goverment.Latitude,
		Longitude:     goverment.Longitude,
	}
}