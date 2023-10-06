package data

import (
	"errors"
	"project-capston/features/vehicles"

	"gorm.io/gorm"
)

type VehicleData struct {
	db *gorm.DB
}

// Delete implements vehicles.VehicleDataInterface.
func (repo *VehicleData) Delete(id uint) error {
	var inputModel Vehicle
	tx:=repo.db.Delete(&inputModel,id)
	if tx.Error != nil{
		return errors.New("fail delete")
	}
	if tx.RowsAffected==0{
		return errors.New("row not affected")
	}
	return nil
}

// SelectAll implements vehicles.VehicleDataInterface.
func (repo *VehicleData) SelectAll() ([]vehicles.VehicleEntity, error) {
	var inputModel []Vehicle
	tx := repo.db.Find(&inputModel)
	if tx.Error != nil {
		return nil, errors.New("get all failed")
	}
	var entity []vehicles.VehicleEntity
	for _, v := range inputModel {
		entity = append(entity, ModelToEntity(v))
	}
	return entity, nil
}

// SelectById implements vehicles.VehicleDataInterface.
func (repo *VehicleData) SelectById(id uint) (vehicles.VehicleEntity, error) {
	var inputModel Vehicle
	tx := repo.db.Preload("").First(&inputModel, id)
	if tx.Error != nil {
		return vehicles.VehicleEntity{}, errors.New("get vehicle by id failed")
	}
	entity := ModelToEntity(inputModel)
	return entity, nil
}

// Update implements vehicles.VehicleDataInterface.
func (repo *VehicleData) Update(input vehicles.VehicleEntity, id uint) error {
	inputModel := EntityToEntity(input)
	tx := repo.db.Model(&Vehicle{}).Where("id=?", id).Updates(inputModel)
	if tx.Error != nil {
		return errors.New("error update vehicle")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

// Insert implements vehicles.VehicleDataInterface.
func (repo *VehicleData) Insert(input vehicles.VehicleEntity) error {
	inputModel := EntityToEntity(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return errors.New("error create vehicle")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

func New(db *gorm.DB) vehicles.VehicleDataInterface {
	return &VehicleData{
		db: db,
	}
}
