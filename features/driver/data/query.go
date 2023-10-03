package data

import (
	"fmt"
	"project-capston/app/middlewares"
	"project-capston/features/driver"
	goverment "project-capston/features/goverment/data"
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
func (repo *driverQuery) SelectAll(pageNumber int, pageSize int) ([]driver.DriverCore, error) {
	// var driverData []Driver

	// offset := (pageNumber - 1) * pageSize

	var driversWithGovernments []struct {
		Driver
		DriverID uint
		goverment.Government
	}

	tx := repo.db.Table("drivers").
		Select("drivers.* ,drivers.id AS DriverID, governments.name").
		Joins("INNER JOIN governments ON drivers.goverment_id=governments.id").
		Scan(&driversWithGovernments)

	for _, u := range driversWithGovernments {
		fmt.Printf("ID : %d,Nama: %s, Email: %s\n", u.DriverID, u.Name, u.Email)
		// for _, o := range u.D {
		// 	fmt.Printf("  Pesanan: %s\n", o.Product)
		// }
	}
	// tx := repo.db.Offset(offset).Limit(pageSize).Find(&driverData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var driverCore []driver.DriverCore

	for _, value := range driversWithGovernments {
		driverCore = append(driverCore, driver.DriverCore{
			Id:            value.DriverID,
			Fullname:      value.Fullname,
			Email:         value.Email,
			Password:      value.Password,
			Token:         value.Token,
			GovermentName: value.Government.Name,
			// GovermentID: value.GovermentID,
			// GovernmentName:value.Go
			Status:        value.Status,
			DrivingStatus: value.DrivingStatus,
			VehicleID:     value.VehicleID,
			// Latitude:      value.Latitude,
			// Longitude:     value.Longitude,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: time.Time{},
		})
	}

	return driverCore, nil

}
