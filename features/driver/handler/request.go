package handler

import (
	"project-capston/features/driver"
	"time"
)

type LoginDriverRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LatLonRequest struct {
	Lat       float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type AcceptOrRejectOrderRequest struct {
	IsAccepted bool `json:"is_accepted"`
}

type DriverRequest struct {
	GovernmentID  uint    `json:"government_id" form:"government_id"`
	Fullname      string  `json:"fullname" form:"fullname"`
	Email         string  `json:"email"`
	Password      string  `json:"password"`
	Token         string  `json:"token"`
	Status        bool    `json:"status" form:"status"`
	DrivingStatus string  `json:"driving_status" form:"driving_status"`
	VehicleID     uint    `json:"vehicle_id" form:"vehicle_id"`
	Latitude      float64 `json:"latitude" form:"latitude"`
	Longitude     float64 `json:"longitude" form:"longitude"`
}

func RequestToCore(input DriverRequest) driver.Core {
	return driver.Core{
		// Id:            0,
		Fullname:      input.Fullname,
		Email:         input.Email,
		Password:      input.Password,
		Token:         input.Token,
		GovermentID:   input.GovernmentID,
		Status:        input.Status,
		DrivingStatus: input.DrivingStatus,
		VehicleID:     input.VehicleID,
		Latitude:      input.Latitude,
		Longitude:     input.Longitude,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     time.Time{},
	}
}
