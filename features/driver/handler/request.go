package handler

type DriverRequest struct {
	GovermentID    uint    `json:"goverment_id" form:"goverment_id"`
	UserID         uint    `json:"user_id" form:"user_id"`
	Name           uint    `json:"name" form:"name"`
	StatusBertugas string  `json:"status_berkendara" form:"status_berkendara"`
	VehicleID      uint    `json:"vehicle_id" form:"vehicle_id"`
	Latitude       float64 `json:"latitude" form:"latitude"`
	Longitude      float64 `json:"longitude" form:"longitude"`
}