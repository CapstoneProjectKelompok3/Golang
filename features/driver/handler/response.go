package handler

type DriverResponse struct {
	Id          uint    `json:"id,omitempty"`
	GovermentID uint    `json:"goverment_id,omitempty"`
	Name        uint    `json:"name,omitempty"`
	Status      string  `json:"status,omitempty"`
	VehicleID   uint    `json:"vehicle_id,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
}
