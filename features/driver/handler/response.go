package handler

type DriverResponse struct {
	Id            uint    `json:"id,omitempty"`
	GovermentID   uint    `json:"goverment_id,omitempty"`
	Fullname      string  `json:"fullname,omitempty"`
	Status        string  `json:"status,omitempty"`
	DrivingStatus string  `json:"driving_status,omitempty"`
	VehicleID     uint    `json:"vehicle_id,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
}
