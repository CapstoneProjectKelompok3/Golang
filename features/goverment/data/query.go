package data

import (
	"project-capston/features/goverment"

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
