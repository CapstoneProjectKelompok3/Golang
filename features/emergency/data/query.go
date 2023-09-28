package data

import (
	"errors"
	"project-capston/features/emergency"

	"gorm.io/gorm"
)

type EmergencyData struct {
	db *gorm.DB
}

// Insert implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) Insert(input emergency.EmergencyEntity) error {
	inputModel:= EntityToModel(input)
	tx:=repo.db.Create(&inputModel)
	if tx.Error != nil{
		return errors.New("failed create data emergency")
	}
	if tx.RowsAffected==0{
		return errors.New("row not affected")
	}
	return nil
}

func New(db *gorm.DB) emergency.EmergencyDataInterface {
	return &EmergencyData{
		db: db,
	}
}
