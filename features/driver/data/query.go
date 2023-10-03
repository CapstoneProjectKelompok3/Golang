package data

import (
	"project-capston/app/middlewares"
	"project-capston/features/driver"
	"time"

	"gorm.io/gorm"
)

type driverQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) driver.DriverDataInterface {
	return &driverQuery{
		db: db,
	}
}

// Insert implements driver.DriverDataInterface.
func (repo *driverQuery) Insert(input driver.Core) error {
	hashedPassword, _ := middlewares.HashedPassword(input.Password)
	input.Password = hashedPassword

	userGorm := CoreToModel(input)

	tx := repo.db.Create(&userGorm) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAll implements driver.DriverDataInterface.
func (repo *driverQuery) SelectAll(pageNumber int, pageSize int) ([]driver.Core, error) {
	var driverData []Driver

	offset := (pageNumber - 1) * pageSize

	tx := repo.db.Offset(offset).Limit(pageSize).Find(&driverData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var driverCore []driver.Core

	for _, value := range driverData {
		driverCore = append(driverCore, driver.Core{
			Id:            value.ID,
			Fullname:      value.Fullname,
			Email:         value.Email,
			Password:      value.Password,
			Token:         value.Token,
			GovermentID:   value.GovermentID,
			Status:        value.Status,
			DrivingStatus: value.DrivingStatus,
			VehicleID:     value.VehicleID,
			Latitude:      value.Latitude,
			Longitude:     value.Longitude,
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
			DeletedAt:     time.Time{},
		})
	}

	return driverCore, nil

}
