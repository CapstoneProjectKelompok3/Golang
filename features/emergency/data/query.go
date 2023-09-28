package data

import (
	"errors"
	"project-capston/features/emergency"

	"gorm.io/gorm"
)

type EmergencyData struct {
	db *gorm.DB
}

// Delete implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) Delete(id uint) error {
	var inputModel Emergency
	tx:=repo.db.Delete(&inputModel,id)
	if tx.Error != nil{
		return errors.New("fail delete emergency")
	}
	if tx.RowsAffected==0{
		return errors.New("id not found")
	}
	return nil
}

// Insert implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) Insert(input emergency.EmergencyEntity) error {
	inputModel := EntityToModel(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return errors.New("failed create data emergency")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

func New(db *gorm.DB) emergency.EmergencyDataInterface {
	return &EmergencyData{
		db: db,
	}
}
