package data

import (
	"errors"
	"math"
	"project-capston/features/goverment"
	"time"

	"gorm.io/gorm"
)

type governmentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) goverment.GovernmentDataInterface {
	return &governmentQuery{
		db: db,
	}
}

func (repo *governmentQuery) Insert(input goverment.Core) error {
	governmentGorm := CoreToModel(input)
	tx := repo.db.Create(&governmentGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectAll implements goverment.GovernmentDataInterface.
func (repo *governmentQuery) SelectAll(pageNumber int, pageSize int) ([]goverment.Core, error) {
	var governmentData []Government

	offset := (pageNumber - 1) * pageSize

	tx := repo.db.Offset(offset).Limit(pageSize).Find(&governmentData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var governmentCore []goverment.Core

	for _, value := range governmentData {
		governmentCore = append(governmentCore, goverment.Core{
			ID:        value.ID,
			Name:      value.Name,
			Type:      value.Type,
			Address:   value.Address,
			Latitude:  value.Latitude,
			Longitude: value.Longitude,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: time.Time{},
		})
	}

	return governmentCore, nil
}

// SelectNearestLocation implements goverment.GovernmentDataInterface.
func (repo *governmentQuery) SelectNearestLocation(latitude float64, longitude float64, radius float64) ([]goverment.Location, error) {
	var governmentData []Government

	// offset := (pageNumber - 1) * pageSize
	// radius := 10

	// tx := repo.db.Find(&governmentData)
	tx := repo.db.Where("6371 * ACOS(SIN(RADIANS(?)) * SIN(RADIANS(latitude)) + COS(RADIANS(?)) * COS(RADIANS(latitude)) * COS(RADIANS(? - longitude))) <= ?", latitude, latitude, longitude, radius).Find(&governmentData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	type Location struct {
		ID        uint
		Name      string
		Latitude  float64
		Longitude float64
		Jarak     float64
	}

	var locations []Location
	err := repo.db.Raw("SELECT id, name, latitude, longitude, (6371 * ACOS(COS(RADIANS(?)) * COS(RADIANS(latitude)) * COS(RADIANS(longitude) - RADIANS(?)) + SIN(RADIANS(?)) * SIN(RADIANS(latitude)))) AS jarak FROM governments WHERE(6371 * ACOS(SIN(RADIANS(?)) * SIN(RADIANS(latitude)) + COS(RADIANS(?)) * COS(RADIANS(latitude)) * COS(RADIANS(? - longitude))) <= ? AND deleted_at IS NULL) ORDER BY jarak ", latitude, longitude, latitude, latitude, latitude, longitude, radius).Scan(&locations).Error

	// for _, location := range locations {
	// 	fmt.Printf("ID: %d, Nama Lokasi: %s, Latitude: %f, Longitude: %f, Jarak: %f km\n", location.ID, location.Name, location.Latitude, location.Longitude, math.Round(location.Jarak*100)/100)
	// }

	if err != nil {
		panic(err)
	}

	var governmentCore []goverment.Location

	for _, value := range locations {
		governmentCore = append(governmentCore, goverment.Location{
			ID:        value.ID,
			Name:      value.Name,
			Latitude:  value.Latitude,
			Longitude: value.Longitude,
			Distance:  math.Round(value.Jarak*100) / 100,
		})
	}

	return governmentCore, nil
}

// Select implements goverment.GovernmentDataInterface.
func (repo *governmentQuery) Select(id uint) (goverment.Core, error) {
	var governmentData Government
	tx := repo.db.Where("id = ?", id).First(&governmentData)
	if tx.Error != nil {
		return goverment.Core{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return goverment.Core{}, errors.New("data not found")
	}

	return ModelToCore(governmentData), nil
}

// Update implements goverment.GovernmentDataInterface.
func (repo *governmentQuery) Update(id uint, input goverment.Core) error {
	governmentGorm := CoreToModel(input)
	tx := repo.db.Model(&Government{}).Where("id = ?", id).Updates(&governmentGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements goverment.GovernmentDataInterface.
func (repo *governmentQuery) Delete(id uint) error {
	var governmentGorm Government
	tx := repo.db.Where("id = ?", id).Delete(&governmentGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
