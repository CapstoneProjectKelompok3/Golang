package data

import (
	"errors"
	"fmt"
	"math"
	"project-capston/app/middlewares"
	"project-capston/features/driver"
	goverment "project-capston/features/goverment/data"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type driverQuery struct {
	db *gorm.DB
}

// SelectProfile implements driver.DriverDataInterface.
func (repo *driverQuery) SelectProfile(id int) (driver.DriverCore, error) {
	var driversWithGovernments struct {
		Driver
		DriverID uint
		goverment.Government
	}

	tx := repo.db.Table("drivers").
		Select("drivers.* ,drivers.id AS DriverID, governments.name,governments.type").
		Joins("INNER JOIN governments ON drivers.goverment_id=governments.id").
		Where("drivers.id=?", id).
		Scan(&driversWithGovernments)

	// tx := repo.db.First(&driverData, id).Scan(&driverData) //

	if tx.Error != nil {
		return driver.DriverCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return driver.DriverCore{}, errors.New("data not found")
	}

	// var driversCore = ModelToCore(driversWithGovernments)

	var driverCore driver.DriverCore

	driverCore.Id = driversWithGovernments.DriverID
	driverCore.Fullname = driversWithGovernments.Driver.Fullname
	driverCore.Email = driversWithGovernments.Driver.Email
	driverCore.Token = driversWithGovernments.Token
	driverCore.GovermentName = driversWithGovernments.Government.Name
	driverCore.GovermentType = driversWithGovernments.Government.Type
	driverCore.DrivingStatus = driversWithGovernments.DrivingStatus
	driverCore.Latitude = driversWithGovernments.Driver.Latitude
	driverCore.Longitude = driversWithGovernments.Driver.Longitude

	return driverCore, nil
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
		fmt.Printf("ID : %d,Nama: %s, Type: %s \n", u.DriverID, u.Name, u.Government.Type)
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

// AcceptOrRejectOrder implements driver.DriverDataInterface.
func (repo *driverQuery) AcceptOrRejectOrder(IsAccepted bool, idDriver int) error {
	if IsAccepted {
		tx := repo.db.Exec("UPDATE drivers SET status=false,driving_status=on_trip WHERE id=?", idDriver) // proses query insert
		if tx.Error != nil {
			return tx.Error
		}

		if tx.Error != nil {
			panic("Failed to update user")
		}

		rowsAffected := tx.RowsAffected
		fmt.Printf("Rows affected: %d\n", rowsAffected)
	} else {
		tx := repo.db.Exec("UPDATE drivers SET status=true,driving_status=on_cancel WHERE id=?", idDriver) // proses query insert
		if tx.Error != nil {
			return tx.Error
		}

		if tx.Error != nil {
			panic("Failed to update user")
		}

		rowsAffected := tx.RowsAffected
		fmt.Printf("Rows affected: %d\n", rowsAffected)
	}

	return nil
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
func (repo *driverQuery) KerahkanDriver(lat string, lon string, police int, hospital int, firestation int, dishub int, SAR int) ([]driver.DriverCore, error) {
	var driversWithGovernments []struct {
		Driver
		DriverID uint
		goverment.Government
		Distance float64
	}

	if police >= 0 && hospital >= 0 && firestation >= 0 {
		sub_query1a := `
		SELECT
				(6371 * acos(cos(radians(drivers.latitude)) * cos(radians(`

		sub_query1b := `)) * 
				cos(radians(`

		sub_query1c := `) - radians(drivers.longitude)) + sin(radians(drivers.latitude)) *
	            sin(radians(`

		sub_query1d := `)))) AS distance,
			    drivers.*,governments.type,governments.name,drivers.id AS DriverID
		FROM
				drivers
		INNER JOIN 
				governments ON governments.id = drivers.goverment_id
		where governments.type='police' AND status=true AND driving_status='on_ready'
	   LIMIT
		`

		query1 := sub_query1a + lat + sub_query1b + lon + sub_query1c + lat + sub_query1d

		sub_query2a := `
		SELECT
				(6371 * acos(cos(radians(drivers.latitude)) * cos(radians(`

		sub_query2b := `)) * 
				cos(radians(`

		sub_query2c := `) - radians(drivers.longitude)) + sin(radians(drivers.latitude)) *
	            sin(radians(`

		sub_query2d := `)))) AS distance,
			    drivers.*,governments.type,governments.name,drivers.id AS DriverID
		FROM
				drivers
		INNER JOIN 
				governments ON governments.id = drivers.goverment_id
		where governments.type='hospital' AND status=true AND driving_status='on_ready'
	   LIMIT
		`

		query2 := sub_query2a + lat + sub_query2b + lon + sub_query2c + lat + sub_query2d

		sub_query3a := `
		SELECT
				(6371 * acos(cos(radians(drivers.latitude)) * cos(radians(`

		sub_query3b := `)) * 
				cos(radians(`

		sub_query3c := `) - radians(drivers.longitude)) + sin(radians(drivers.latitude)) *
	            sin(radians(`

		sub_query3d := `)))) AS distance,
			    drivers.*,governments.type,governments.name,drivers.id AS DriverID
		FROM
				drivers
		INNER JOIN 
				governments ON governments.id = drivers.goverment_id
		where governments.type='firestation' AND status=true AND driving_status='on_ready'
	   LIMIT
		`

		query3 := sub_query3a + lat + sub_query3b + lon + sub_query3c + lat + sub_query3d

		police_query := fmt.Sprintf("%s%d", query1, police)
		hospital_query := fmt.Sprintf("%s%d", query2, hospital)
		firestation_query := fmt.Sprintf("%s%d", query3, firestation)

		sql := fmt.Sprintf("(%s) UNION ALL (%s) UNION ALL (%s) %s", police_query, hospital_query, firestation_query, "ORDER BY distance")

		tx := repo.db.Raw(sql).Scan(&driversWithGovernments)
		fmt.Println("adasdds", tx)
		if tx.Error != nil {
			return nil, tx.Error
		}

	} else if police >= 0 && hospital >= 0 {
		sub_query1a := `
		SELECT
				(6371 * acos(cos(radians(drivers.latitude)) * cos(radians(`

		sub_query1b := `)) * 
				cos(radians(`

		sub_query1c := `) - radians(drivers.longitude)) + sin(radians(drivers.latitude)) *
	            sin(radians(`

		sub_query1d := `)))) AS distance,
			    drivers.*,governments.type,governments.name,drivers.id AS DriverID
		FROM
				drivers
		INNER JOIN 
				governments ON governments.id = drivers.goverment_id
		where governments.type='police' AND status=true AND driving_status='on_ready'
	   LIMIT
		`

		query1 := sub_query1a + lat + sub_query1b + lon + sub_query1c + lat + sub_query1d

		sub_query2a := `
		SELECT
				(6371 * acos(cos(radians(drivers.latitude)) * cos(radians(`

		sub_query2b := `)) * 
				cos(radians(`

		sub_query2c := `) - radians(drivers.longitude)) + sin(radians(drivers.latitude)) *
	            sin(radians(`

		sub_query2d := `)))) AS distance,
			    drivers.*,governments.type,governments.name,drivers.id AS DriverID
		FROM
				drivers
		INNER JOIN 
				governments ON governments.id = drivers.goverment_id
		where governments.type='hospital' AND status=true AND driving_status='on_ready'
	   LIMIT
		`

		query2 := sub_query2a + lat + sub_query2b + lon + sub_query2c + lat + sub_query2d

		police_query := fmt.Sprintf("%s%d", query1, police)
		hospital_query := fmt.Sprintf("%s%d", query2, hospital)

		sql := fmt.Sprintf("(%s) UNION ALL (%s) %s", police_query, hospital_query, "ORDER BY distance")

		tx := repo.db.Raw(sql).Scan(&driversWithGovernments)
		fmt.Println("adasdds", tx)
		if tx.Error != nil {
			return nil, tx.Error
		}

	} else if police >= 0 {

		sub_query1a := `
		SELECT
				(6371 * acos(cos(radians(drivers.latitude)) * cos(radians(`

		sub_query1b := `)) * 
				cos(radians(`

		sub_query1c := `) - radians(drivers.longitude)) + sin(radians(drivers.latitude)) *
	            sin(radians(`

		sub_query1d := `)))) AS distance,
			    drivers.*,governments.type,governments.name,drivers.id AS DriverID
		FROM
				drivers
		INNER JOIN 
				governments ON governments.id = drivers.goverment_id
		where governments.type='police' AND status=true AND driving_status='on_ready'
	   LIMIT
		`
		query := sub_query1a + lat + sub_query1b + lon + sub_query1c + lat + sub_query1d

		sql := fmt.Sprintf("%s%d", query, police)

		tx := repo.db.Raw(sql).Scan(&driversWithGovernments)

		fmt.Println("adasdds", tx)
		if tx.Error != nil {
			return nil, tx.Error
		}

		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	tokenKasus := uuid.New()
	fmt.Println(tokenKasus)
	for _, u := range driversWithGovernments {
		fmt.Printf("ID : %d,Nama: %s, Email: %s\n", u.DriverID, u.Name, u.Email)

		repo.db.Exec("UPDATE drivers SET token = ? WHERE id = ? ", tokenKasus, u.DriverID)
		// Update kolom bertipe ENUM
		// result := repo.db.Exec("UPDATE drivers SET driving_status='on_demand' WHERE id=?", u.DriverID).Scan(&Driver{})
		result := repo.db.Model(&Driver{}).Where("id = ?", u.DriverID).Update("driving_status", "on_demand")

		if result.Error != nil {
			panic("Gagal melakukan UPDATE")
		}

		// Tampilkan jumlah baris yang terpengaruh
		rowsAffected := result.RowsAffected
		fmt.Printf("Jumlah baris yang terpengaruh: %d\n", rowsAffected)
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
