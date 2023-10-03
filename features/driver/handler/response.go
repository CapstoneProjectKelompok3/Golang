package handler

type DriverResponse struct {
	Id            uint    `json:"id,omitempty"`
	GovermentName string  `json:"goverment_name,omitempty"`
	Fullname      string  `json:"fullname,omitempty"`
	Toker         string  `json:"toker,omitempty"`
	Status        string  `json:"status,omitempty"`
	DrivingStatus string  `json:"driving_status,omitempty"`
	VehicleID     uint    `json:"vehicle_id,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
}
