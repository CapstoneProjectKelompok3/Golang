package driver

import "time"

type DriverEntity struct {
	Id             uint
	CreateAt       time.Time
	UpdateAt       time.Time
	DeleteAt       time.Time
	GovermentID    uint
	UserID         uint
	Name           uint
	StatusBertugas string
	VehicleID      uint
	Latitude       float64
	Longitude      float64
}