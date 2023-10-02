package data

import (
	"project-capston/features/goverment"
	"time"

	"gorm.io/gorm"
)

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

func CoreToModel(dataCore goverment.Core) Government {
	return Government{
		ID:        dataCore.ID,
		Name:      dataCore.Name,
		Type:      dataCore.Type,
		Address:   dataCore.Address,
		Latitude:  dataCore.Latitude,
		Longitude: dataCore.Longitude,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
}

func ModelToCore(dataModel Government) goverment.Core {
	return goverment.Core{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Type:      dataModel.Type,
		Address:   dataModel.Address,
		Latitude:  dataModel.Latitude,
		Longitude: dataModel.Longitude,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
	}
}
