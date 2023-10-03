package data

import (
	"errors"
	"fmt"
	"math"
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
		Select("drivers.* ,drivers.id AS DriverID, governments.name,governments.type").
		Joins("INNER JOIN governments ON drivers.goverment_id=governments.id").
		Scan(&driversWithGovernments)

	for _, u := range driversWithGovernments {
		fmt.Printf("ID : %d,Nama: %s, Type: %s\n", u.DriverID, u.Name, u.Government.Type)
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
			GovermentType: value.Government.Type,
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

// KerahkanDriver implements driver.DriverDataInterface.
func (repo *driverQuery) KerahkanDriver(police int, hospital int, firestation int, dishub int, SAR int) ([]driver.DriverCore, error) {
	var driversWithGovernments []struct {
		Driver
		DriverID uint
		goverment.Government
		Distance float64
	}

	fmt.Println("Jmlh Police ", police)
	fmt.Println("Jmlh Police ", hospital)

	if police >= 0 && hospital >= 0 {
		// tx := repo.db.Table("drivers").
		// 	Select("drivers.* ,drivers.id AS DriverID, governments.name,governments.type").
		// 	Joins("INNER JOIN governments ON drivers.goverment_id=governments.id").
		// 	Where("governments.type='police' LIMIT ?", police).
		// 	Where("governments.type='hospital' LIMIT ?", hospital).
		// 	Scan(&driversWithGovernments)

		// if tx.Error != nil {
		// 	return nil, tx.Error
		// }

		query1 := `
		SELECT
				(6371 * acos(cos(radians(drivers.latitude)) * cos(radians(-6.304990)) * cos(radians(106.820500) - radians(drivers.longitude)) + sin(radians(drivers.latitude)) * sin(radians(-6.304990)))) AS distance,
			drivers.*,governments.type
			FROM
				drivers
			INNER JOIN 
						governments ON governments.id = drivers.goverment_id
			where governments.type='police' AND status=true
			ORDER BY distance LIMIT
		`

		query2 := `
		SELECT
				(6371 * acos(cos(radians(drivers.latitude)) * cos(radians(-6.304990)) * cos(radians(106.820500) - radians(drivers.longitude)) + sin(radians(drivers.latitude)) * sin(radians(-6.304990)))) AS distance,
			drivers.*,governments.type
			FROM
				drivers
			INNER JOIN 
						governments ON governments.id = drivers.goverment_id
			where governments.type='hospital' AND status=true
			ORDER BY distance LIMIT 
		 `

		police_query := fmt.Sprintf("%s%d", query1, police)
		hospital_query := fmt.Sprintf("%s%d", query2, hospital)

		tx := repo.db.Raw(fmt.Sprintf("(%s) UNION ALL (%s)", police_query, hospital_query)).Scan(&driversWithGovernments)
		fmt.Println("adasdds", tx)
		if tx.Error != nil {
			return nil, tx.Error
		}

	} else if police >= 0 {
		tx := repo.db.Table("drivers").
			Select("drivers.* ,drivers.id AS DriverID, governments.name,governments.type").
			Joins("INNER JOIN governments ON drivers.goverment_id=governments.id").
			Where("governments.type='police' LIMIT ?", police).
			Scan(&driversWithGovernments)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	for _, u := range driversWithGovernments {
		fmt.Printf("ID : %d,Nama: %s, Email: %s\n", u.DriverID, u.Name, u.Email)
		repo.db.Exec("UPDATE drivers SET token = ? WHERE id = ? ", "Terima Kasus", u.DriverID)
		// for _, o := range u.D {
		// 	fmt.Printf("  Pesanan: %s\n", o.Product)
		// }
	}
	// tx := repo.db.Offset(offset).Limit(pageSize).Find(&driverData)

	var driverCore []driver.DriverCore

	for _, value := range driversWithGovernments {

		driverCore = append(driverCore, driver.DriverCore{
			Id:            value.DriverID,
			Fullname:      value.Fullname,
			Email:         value.Email,
			Password:      value.Password,
			Token:         value.Token,
			GovermentName: value.Government.Name,
			GovermentType: value.Government.Type,
			Status:        value.Status,
			DrivingStatus: value.DrivingStatus,
			VehicleID:     value.VehicleID,
			Latitude:      value.Driver.Latitude,
			Distance:      math.Floor(value.Distance*100) / 100,
			Longitude:     value.Driver.Longitude,
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
			DeletedAt:     time.Time{},
		})
	}

	return driverCore, nil
}
