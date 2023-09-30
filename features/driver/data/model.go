package data

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	GovermentID uint
	UserID uint
	Name uint
	StatusBertugas string
	VehicleID uint
	Latitude   float64
	Longitude  float64
}