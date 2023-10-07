package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"project-capston/app/middlewares"
	"project-capston/features/driver"
	goverment "project-capston/features/goverment/data"
	unites "project-capston/features/unit/data"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var ctx = context.Background()

type driverQuery struct {
	db *gorm.DB
}

// SelectHistori implements driver.DriverDataInterface.
func (repo *driverQuery) SelectHistori(idUnit uint) (uint, error) {
	
	var unit unites.Unit
	tx:=repo.db.First(&unit,idUnit)
	if tx.Error != nil{
		return 0,errors.New("error select unit")
	}

	var history unites.UnitHistory
	txx:=repo.db.Where("unit_id=? and status=? and driver_id=?",unit.ID,"-",uint(0)).First(&history)
	if txx.Error != nil{
		return 0,errors.New("error select histori")
	}	
	return history.ID,nil
}

// SelectUnit implements driver.DriverDataInterface.
func (repo *driverQuery) SelectUnit(idEmergenci uint) ([]uint, []string, error) {
	var inputModel []unites.Unit
	tx := repo.db.Where("emergencies_id=?", idEmergenci).Find(&inputModel)
	if tx.Error != nil {
		return nil, nil, errors.New("failed get type unit")
	}
	var tipe []string
	for _, v := range inputModel {
		tipe = append(tipe, v.Type)
	}

	var id []uint
	for _, v := range inputModel {
		id = append(id, v.ID)
	}
	return id, tipe, nil
}

func (repo *driverQuery) UpdateHistoryUnit(idDriver uint, idUnitHistori uint) error {
	var units unites.UnitHistory
	units.DriverID = idDriver
	tx := repo.db.Model(&unites.UnitHistory{}).Where("id=?", idUnitHistori).Updates(units)
	if tx.Error != nil {
		return errors.New("failed update histori")
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

// CreateUnit implements driver.DriverDataInterface.
func (repo *driverQuery) CreateUnit(idEmergency uint, tipe []string, count []int) error {

	for i := 0; i < len(tipe); i++ {
		if count[i] != 0 {
			unitData := unites.Unit{
				EmergenciesID: idEmergency,
				Type:          tipe[i],
				SumOfUnit:     count[i],
			}
			tx := repo.db.Create(&unitData)
			if tx.Error != nil {
				return errors.New("failed create unit")
			}
			if tx.RowsAffected == 0 {
				return errors.New("row not affected")
			}
		}
	}
	return nil
}

func (repo *driverQuery) CreateUnitHistori(idEmergency uint) error {
	var inputUnit []unites.Unit
	txx := repo.db.Where("emergencies_id=?", idEmergency).Find(&inputUnit)
	if txx.Error != nil {
		return errors.New("failed get unit history")
	}

	tx := repo.db.Begin()
	for _, v := range inputUnit {
		for j := 0; j < v.SumOfUnit; j++ {
			inputModel := unites.UnitHistory{
				UnitID: v.ID,
			}
			if err := tx.Create(&inputModel).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// Delete implements driver.DriverDataInterface.
func (repo *driverQuery) Delete(id uint) error {
	var inputModel Driver
	tx := repo.db.Delete(&inputModel, id)
	if tx.Error != nil {
		return errors.New("delete driver failed")
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

// SelectCountDriver implements driver.DriverDataInterface.
func (repo *driverQuery) SelectCountDriver() (int64, error) {
	var input []Driver
	tx := repo.db.Find(&input)
	if tx.Error != nil {
		return 0, errors.New("fail get driver")
	}
	count := tx.RowsAffected
	return count, nil
}

// Logout implements driver.DriverDataInterface.
func (repo *driverQuery) Logout(id int) error {
	tx := repo.db.Exec("UPDATE drivers SET status=0 WHERE id=?", id) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// FinishTrip implements driver.DriverDataInterface.
func (repo *driverQuery) FinishTrip(id int) error {
	tx := repo.db.Exec("UPDATE drivers SET driving_status='on_finished' WHERE id=?", id) // proses query insert
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// DriverOnTrip implements driver.DriverDataInterface.
func (repo *driverQuery) DriverOnTrip(id int, lat float64, long float64) (driver.DriverCore, error) {
	redisClient := middlewares.CreateRedisClient()

	start := int64(0)
	end := int64(-1) // Membaca semua elemen dalam array
	resultRedis, errRedis := redisClient.LRange(ctx, "data_array", start, end).Result()
	if errRedis != nil {
		fmt.Println("Error reading data from Redis:", errRedis)
	}

	//2. dapatkan data lat long dari redis
	fmt.Println("Data dari Redis:", resultRedis)
	fmt.Println("Data dari Redis:", resultRedis[0])
	fmt.Println("Data dari Redis:", resultRedis[1])

	// latString := strconv.FormatFloat(lat, 'f', -1, 64)
	// lonString := strconv.FormatFloat(long, 'f', -1, 64)

	result := repo.db.Exec("UPDATE drivers SET latitude=?,longitude=? WHERE id=?", lat, long, id)

	fmt.Println("Result Udate", result)

	// Periksa error
	if result.Error != nil {
		fmt.Println("Error executing raw SQL:", result.Error)
	}

	// Periksa jumlah baris yang terpengaruh
	fmt.Printf("Jumlah baris yang terpengaruh: %d\n", result.RowsAffected)

	var driversWithGovernments struct {
		Driver
		Distance float64
		DriverID uint
		goverment.Government
	}

	sub_query1a := `
		SELECT
				(6371 * acos(cos(radians(`

	sub_query1ab := `)) * cos(radians(`

	sub_query1b := `)) * 
				cos(radians(`

	sub_query1c := `) - radians(`

	sub_query1cd := `drivers.longitude)) + sin(radians(`

	sub_query1cd1 := `)) *
	            sin(radians(`

	sub_query1d := `)))) AS Distance,
			    drivers.*,governments.type,governments.name,drivers.id AS DriverID
		FROM
				drivers
		INNER JOIN 
				governments ON governments.id = drivers.goverment_id
		where drivers.id =
		`
	query := sub_query1a + "drivers.latitude" + sub_query1ab + resultRedis[1] + sub_query1b + resultRedis[0] + sub_query1c + sub_query1cd + "drivers.latitude" + sub_query1cd1 + resultRedis[1] + sub_query1d

	sql := fmt.Sprintf("%s%d", query, id)

	fmt.Println("Query On Trip", sql)

	tx := repo.db.Raw(sql).Scan(&driversWithGovernments)

	if tx.Error != nil {
		return driver.DriverCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return driver.DriverCore{}, errors.New("data not found")
	}

	fmt.Println("Driver Government Distance", driversWithGovernments.Distance)

	// var driversCore = ModelToCore(driversWithGovernments)

	var driverCore driver.DriverCore

	driverCore.Id = driversWithGovernments.DriverID
	driverCore.Fullname = driversWithGovernments.Driver.Fullname
	driverCore.Status = driversWithGovernments.Status
	driverCore.Email = driversWithGovernments.Driver.Email
	driverCore.Token = driversWithGovernments.Token
	driverCore.GovermentName = driversWithGovernments.Government.Name
	driverCore.GovermentType = driversWithGovernments.Government.Type
	driverCore.DrivingStatus = driversWithGovernments.DrivingStatus
	driverCore.Latitude = driversWithGovernments.Driver.Latitude
	driverCore.Longitude = driversWithGovernments.Driver.Longitude
	driverCore.Distance = driversWithGovernments.Distance

	fmt.Println("Driver Core", driverCore.Distance)

	// // Mengubah []float64 menjadi string
	// dataStr := make([]string, len(data))
	// for i, v := range data {
	// 	dataStr[i] = strconv.FormatFloat(v, 'f', -1, 64)
	// }

	// // Marshal data ke format JSON (opsional)
	// dataJSON, err := json.Marshal(dataStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Key Redis untuk menyimpan nilai float
	key := "myFloatValue"

	// Nilai float yang ingin Anda simpan
	floatValue := driverCore.Distance

	// Konversi float menjadi string
	stringValue := strconv.FormatFloat(floatValue, 'f', -1, 64)

	// Simpan nilai string dalam Redis
	err := redisClient.Set(ctx, key, stringValue, 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Nilai float %f berhasil disimpan di Redis dengan key: %s\n", floatValue, key)

	// Mengambil nilai dari Redis dan mengonversinya kembali ke float
	storedValue, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}

	// Mengonversi nilai string kembali ke float
	convertedValue, err := strconv.ParseFloat(storedValue, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Nilai float yang diambil dari Redis: %f\n", convertedValue)

	return driverCore, nil
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
	driverCore.Status = driversWithGovernments.Status
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
			fmt.Println("kwrjjerkekerwkew")
			return driver.Core{}, errors.New("Wrong email or password")
		}
	} else {
		return driver.Core{}, errors.New("data not found")
	}

	dataLogin = ModelToCore(data)

	return dataLogin, nil
}

// KerahkanDriver implements driver.DriverDataInterface.
func (repo *driverQuery) KerahkanDriver(lat string, lon string, police int, hospital int, firestation int, dishub int, SAR int) ([]driver.DriverCore, error) {
	//1. Simpan Lat Long di dalam redis
	redisClient := middlewares.CreateRedisClient()

	if redisClient == nil {
		fmt.Println("Gagal terhubung ke Redis")
		// return c.String(http.StatusInternalServerError, "Gagal terhubung ke Redis")
	}

	data := []string{lat, lon}

	// Simpan array dalam Redis
	// for _, item := range data {
	errRedis := redisClient.LPush(ctx, "data_array", data).Err()
	if errRedis != nil {
		fmt.Println("Gagal menyimpan array di Redis", errRedis.Error())
		// return c.String(http.StatusInternalServerError, "Gagal menyimpan array di Redis")
	} else {
		fmt.Println("Berhasil menyimpan array di Redis", data)
	}

	var driversWithGovernments []struct {
		Driver
		DriverID uint
		goverment.Government
		Distance float64
	}

	var sql string

	if police > 0 && hospital > 0 && firestation > 0 {
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
		ORDER BY distance ASC
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
		ORDER BY distance ASC
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
		ORDER BY distance ASC
	   LIMIT
		`

		query3 := sub_query3a + lat + sub_query3b + lon + sub_query3c + lat + sub_query3d

		police_query := fmt.Sprintf("%s%d", query1, police)
		hospital_query := fmt.Sprintf("%s%d", query2, hospital)
		firestation_query := fmt.Sprintf("%s%d", query3, firestation)

		sql = fmt.Sprintf("(%s) UNION ALL (%s) UNION ALL (%s) %s", police_query, hospital_query, firestation_query, "ORDER BY distance")

		tx := repo.db.Raw(sql).Scan(&driversWithGovernments)

		fmt.Println("adasdds", tx)
		if tx.Error != nil {
			return nil, tx.Error
		}

	} else if police > 0 && hospital > 0 {
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
		ORDER BY distance ASC
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
		ORDER BY distance ASC
	   LIMIT
		`

		query2 := sub_query2a + lat + sub_query2b + lon + sub_query2c + lat + sub_query2d

		police_query := fmt.Sprintf("%s%d", query1, police)
		hospital_query := fmt.Sprintf("%s%d", query2, hospital)

		sql := fmt.Sprintf("(%s) UNION ALL (%s) %s", police_query, hospital_query, "ORDER BY distance")
		fmt.Println("Query Police Hospital", sql)
		tx := repo.db.Raw(sql).Scan(&driversWithGovernments)
		fmt.Println("adasdds", tx)
		if tx.Error != nil {
			return nil, tx.Error
		}

	} else if police > 0 {

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
		ORDER BY distance ASC
	   LIMIT
		`
		query := sub_query1a + lat + sub_query1b + lon + sub_query1c + lat + sub_query1d

		sql = fmt.Sprintf("%s%d", query, police)

		fmt.Println("Query Police", sql)

		tx := repo.db.Raw(sql).Scan(&driversWithGovernments)

		// fmt.Println("adasdds", tx)
		if tx.Error != nil {
			return nil, tx.Error
		}

		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if hospital > 0 {
		fmt.Println("Hospital", hospital)
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
		where governments.type='hospital' AND status=true AND driving_status='on_ready'
		ORDER BY distance ASC
	   LIMIT
		`
		query := sub_query1a + lat + sub_query1b + lon + sub_query1c + lat + sub_query1d

		sql := fmt.Sprintf("%s%d", query, hospital)

		tx := repo.db.Raw(sql).Scan(&driversWithGovernments)

		// fmt.Println("adasdds", tx)
		if tx.Error != nil {
			return nil, tx.Error
		}

		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	//2. Generate Token Kasus
	tokenKasus := uuid.New()

	fmt.Println("Sql Execute", sql)

	for _, u := range driversWithGovernments {
		fmt.Printf("ID : %d,Nama: %s, Email: %s\n", u.DriverID, u.Name, u.Email)

		//3 Update Token
		repo.db.Exec("UPDATE drivers SET token = ? WHERE id = ? ", tokenKasus, u.DriverID)

		//4. Set driving status menjadi on_demand
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

		dataIdUser := value.DriverID
		jsonData, err := json.Marshal(dataIdUser)
		if err != nil {
			fmt.Println("Error marshaling data:", err)

		}
		// Simpan array dalam Redis
		// for _, item := range data {
		errRedisIdUser := redisClient.LPush(ctx, "id_user", jsonData).Err()

		if errRedisIdUser != nil {
			fmt.Println("Gagal menyimpan array id user di Redis", errRedisIdUser.Error())
			// return c.String(http.StatusInternalServerError, "Gagal menyimpan array di Redis")
		} else {
			fmt.Println("Berhasil menyimpan array id user di Redis", jsonData)
		}

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

// AcceptOrRejectOrder implements driver.DriverDataInterface.
func (repo *driverQuery) AcceptOrRejectOrder(IsAccepted bool, idDriver int) error {
	//1. Konek Ke redis
	redisClient := middlewares.CreateRedisClient()

	if redisClient == nil {
		fmt.Println("Gagal terhubung ke Redis")
		// return c.String(http.StatusInternalServerError, "Gagal terhubung ke Redis")
	}

	//2. Membaca data array latitude longitude didalm array redis
	start := int64(0)
	end := int64(-1) // Membaca semua elemen dalam array
	result, errRedis := redisClient.LRange(ctx, "data_array", start, end).Result()
	if errRedis != nil {
		fmt.Println("Error reading data from Redis:", errRedis)
	}

	//2. dapatkan data lat long dari redis
	fmt.Println("Data dari Redis:", result)
	fmt.Println("Data dari Redis:", result[0])
	fmt.Println("Data dari Redis:", result[1])

	var driversWithGovernments struct {
		Driver
		DriverID uint
		goverment.Government
	}

	//3 Tampilkan data dari driver nya berdasarkan id driver dari JWT TOKEN
	repo.db.Table("drivers").
		Select("drivers.* ,drivers.id AS DriverID, governments.name,governments.type").
		Joins("INNER JOIN governments ON drivers.goverment_id=governments.id").
		Where("drivers.id=?", idDriver).
		Scan(&driversWithGovernments)

	//4. Cek apakah drivers punya order atau tidak dari token kasus didalam tabel
	if driversWithGovernments.Token != "" {
		if IsAccepted {
			if driversWithGovernments.DrivingStatus != "on_trip" {
				tx := repo.db.Exec("UPDATE drivers SET status=false,driving_status='on_trip' WHERE id=?", idDriver) // proses query insert
				if tx.Error != nil {
					return tx.Error
				}

				if tx.Error != nil {
					panic("Failed to update user")
				}

				rowsAffected := tx.RowsAffected
				fmt.Printf("Rows affected: %d\n", rowsAffected)
			} else {
				err := errors.New("Sorry But Now you are on the way you cannot receive order again")

				// Melemparkan error
				if err != nil {
					fmt.Println(err.Error())
				}
				return err
			}
		} else {
			//5 Ubah status = true driving_status=on_cancel
			tx := repo.db.Exec("UPDATE drivers SET status=true,driving_status='on_cancel' WHERE id=?", idDriver) // proses query insert

			if tx.Error != nil {
				return tx.Error
			}

			//6. Dapatkan Value dari government type sesuai dengan government type user
			fmt.Printf(" Type: %s \n", driversWithGovernments.Government.Type)
			governmentType := "'" + driversWithGovernments.Government.Type + "'"

			//7. Tampilkan /dapatkan data driver lain yang telah di assigned dengan filter
			//query government_status yang sama berdasarkan lokasi terdekat dari history lat long
			//didalam redis

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
			where governments.type=`

			sub_query1e := ` AND drivers.status=true AND drivers.driving_status='on_ready'  AND drivers.id !=1
		   LIMIT
			`
			queryOther := sub_query1a + result[0] + sub_query1b + result[1] + sub_query1c + result[0] + sub_query1d + governmentType + sub_query1e

			sqlOther := fmt.Sprintf("%s%d", queryOther, 1)

			var otherDriverWithGovernments struct {
				Driver
				DriverID uint
				goverment.Government
			}

			//8 Dapatkan serta lempar token dari user login ke user other yang di assigned
			fmt.Printf(" Type: %s \n", driversWithGovernments.Driver.Token)
			token := driversWithGovernments.Driver.Token
			fmt.Println("Token", token)

			otherDriver := &otherDriverWithGovernments

			repo.db.Raw(sqlOther).Scan(otherDriver)

			//9. Tampilkan Driver lain yang di dapatkan
			fmt.Println("Driver Id lain", otherDriver.Driver.ID)

			sqlAssignedTokenToOTherDriver := "UPDATE drivers SET token=?,driving_status='on_demand' WHERE ID=?"
			repo.db.Exec(sqlAssignedTokenToOTherDriver, token, otherDriver.Driver.ID)

			sqlRemoveMyToken := "UPDATE drivers SET token=? WHERE ID=?"
			repo.db.Exec(sqlRemoveMyToken, "", driversWithGovernments.DriverID)

		}

		return nil

	} else {
		err := errors.New("Sorry But Now you don't have any order")

		// Melemparkan error
		if err != nil {
			fmt.Println(err.Error())
		}
		return err
	}

}
