package data

import (
	"errors"
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
