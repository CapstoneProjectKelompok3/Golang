package data

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	EmergenciesID uint
	VehicleID uint
}

type UnitHistory struct{
	gorm.Model
	EmergenciesID uint
	VehicleID uint
	Status string
	AlasanPenolakan string
}