package unit

import "time"

type UnitEntity struct {
	Id            uint
	CreateAt      time.Time
	UpdateAt      time.Time
	DeleteAt      time.Time
	EmergenciesID uint
	VehicleID     uint
}

type UnitHistoryEntity struct {
	Id              uint
	CreateAt        time.Time
	UpdateAt        time.Time
	DeleteAt        time.Time
	EmergenciesID   uint
	VehicleID       uint
	Status          string
	AlasanPenolakan string
}