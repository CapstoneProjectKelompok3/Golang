package vehicles

import "time"

type VehicleEntity struct {
	Id          uint
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    time.Time
	GovermentID uint
	Plate       string
	Status      bool
}
