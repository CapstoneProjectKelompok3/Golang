package data

import (
	"project-capston/features/vehicles"

	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	GovermentID uint
	Plate string
	Status bool `gorm:"default:true"`
}

func ModelToEntity(vehicle Vehicle)vehicles.VehicleEntity{
	return vehicles.VehicleEntity{
		Id:          vehicle.ID,
		CreateAt:    vehicle.CreatedAt,
		UpdateAt:    vehicle.UpdatedAt,
		DeleteAt:    vehicle.DeletedAt.Time,
		GovermentID: vehicle.GovermentID,
		Plate:       vehicle.Plate,
		Status:      vehicle.Status,
	}
}

func EntityToEntity(vehicle vehicles.VehicleEntity)Vehicle{
	return Vehicle{
		GovermentID: vehicle.GovermentID,
		Plate:       vehicle.Plate,
		Status:      vehicle.Status,
	}
}
