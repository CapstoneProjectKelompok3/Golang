package data

import (
	"project-capston/features/vehicles"
	"time"

	"gorm.io/gorm"
)

type Vehicle struct {
	gorm.Model
	GovermentID uint
	Plate string
	Status bool `gorm:"default:true"`
	Goverments Government `gorm:"foreignKey:GovermentID"`
}
type Government struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;size:255"`
	Type      string `gorm:"type:enum('hospital','police','firestation','dishub','SAR');column:type;default:hospital"`
	Address   string
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
		Goverment: GovermentModelToEntity(vehicle.Goverments),
	}
}

func EntityToEntity(vehicle vehicles.VehicleEntity)Vehicle{
	return Vehicle{
		GovermentID: vehicle.GovermentID,
		Plate:       vehicle.Plate,
		Status:      vehicle.Status,
	}
}

func GovermentModelToEntity(goverment Government)vehicles.GovernmentEntity{
	return vehicles.GovernmentEntity{
		ID:        goverment.ID,
		Name:      goverment.Name,
		Type:      goverment.Type,
		Address:   goverment.Address,
		Latitude:  goverment.Latitude,
		Longitude: goverment.Latitude,
		CreatedAt: goverment.CreatedAt,
		UpdatedAt: goverment.UpdatedAt,
		DeletedAt: goverment.DeletedAt.Time,
	}
}

