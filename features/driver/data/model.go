package data

import (
	"project-capston/features/driver"
	"time"

	"gorm.io/gorm"
)

type Driver struct {
	ID            uint   `gorm:"primaryKey"`
	Fullname      string `gorm:"type:varchar(100)"`
	Email         string `gorm:"unique;size:255"`
	Password      string `gorm:"type:varchar(255);unique_index"`
	Token         string `gorm:"type:varchar(255);unique_index"`
	GovermentID   uint
	Status        bool
	DrivingStatus string `gorm:"type:enum('on_ready','on_demand','on_trip','on_finished','on_cancel');column:driving_status;default:on_ready"`
	VehicleID     uint
	Latitude      float64
	Longitude     float64
	EmergencyId   uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func CoreToModel(dataCore driver.Core) Driver {
	return Driver{
		// ID:            0,
		Fullname:      dataCore.Fullname,
		Email:         dataCore.Email,
		Password:      dataCore.Password,
		GovermentID:   dataCore.GovermentID,
		Status:        dataCore.Status,
		EmergencyId:   dataCore.EmergenciesID,
		DrivingStatus: dataCore.DrivingStatus,
		VehicleID:     dataCore.VehicleID,
		Latitude:      dataCore.Latitude,
		Longitude:     dataCore.Longitude,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}

func ModelToCore(dataModel Driver) driver.Core {
	return driver.Core{
		Id:            dataModel.ID,
		Fullname:      dataModel.Fullname,
		Email:         dataModel.Email,
		Password:      dataModel.Password,
		GovermentID:   dataModel.GovermentID,
		Status:        dataModel.Status,
		DrivingStatus: dataModel.DrivingStatus,
		EmergenciesID: dataModel.EmergencyId,
		VehicleID:     dataModel.VehicleID,
		Latitude:      dataModel.Latitude,
		Longitude:     dataModel.Longitude,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}
}
