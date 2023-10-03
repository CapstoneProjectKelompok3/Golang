package data

import (
	"errors"
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

// Login implements driver.DriverDataInterface.
func (repo *driverQuery) Login(email string, password string) (dataLogin driver.Core, err error) {
	var data Driver

	repo.db.Raw("SELECT * FROM drivers WHERE email=?", email).Scan(&data)
	samePassword := middlewares.CheckPassword(password, data.Password)

	fmt.Println("is same", samePassword)
	fmt.Println("is same", password)
	fmt.Println("data password", data.Password)

	if samePassword {
		fmt.Println("isi data", data.Password)
		// repo.db.Raw("SELECT * FROM users WHERE email=?", email).Scan(&data)
		// fmt.Println("data", data)
		query := `
		SELECT *FROM drivers WHERE email=? AND password=?
		`
		fmt.Println("Query", query)

		tx := repo.db.Where("email = ? and password = ?", email, data.Password).Find(&data)
		repo.db.Exec("UPDATE drivers SET status=1 WHERE email=?", email)

		if tx.Error != nil {
			return driver.Core{}, tx.Error
		}

		if tx.RowsAffected == 0 {
			return driver.Core{}, errors.New("data not found")
		}
	} else {
		return driver.Core{}, errors.New("data not found")
	}

	dataLogin = ModelToCore(data)

	return dataLogin, nil
}
