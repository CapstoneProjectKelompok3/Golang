package data

import (
	"errors"
	"project-capston/features/emergency"

	"gorm.io/gorm"
)

type EmergencyData struct {
	db *gorm.DB
}

// SelectAll implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) SelectAll(param emergency.QueryParams) (int64, []emergency.EmergencyEntity, error) {
	var inputModel []Emergency
	var totalEmergency int64

	query:=repo.db

	if param.IsClassDashboard{
		offset := (param.Page-1)*param.ItemsPerPage
		// if param.SearchName !=""{
		// 	query=query.Where("caller_id like ? or receiver_id=?","%"+param.SearchName+"%","%"+param.SearchName+"%")
		// }
		tx:=query.Find(&inputModel)
		if tx.Error != nil{
			return 0, nil, errors.New("failed get count emergency")
		}
		totalEmergency=tx.RowsAffected
		query=query.Offset(offset).Limit(param.ItemsPerPage)
	}
	// if param.SearchName !=""{
	// 	query=query.Where("caller_id like ? or receiver_id=?","%"+param.SearchName+"%","%"+param.SearchName+"%")
	// }
	tx:=query.Find(&inputModel)
	if tx.Error != nil{
		return 0,nil,errors.New("error get all emergency")
	}

	var output []emergency.EmergencyEntity
	for _,V:=range inputModel{
		output = append(output, ModelToEntity(V))
	}
	return totalEmergency,output,nil
}

// SelectById implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) SelectById(id uint) (emergency.EmergencyEntity, error) {
	var inputModel Emergency
	tx := repo.db.First(&inputModel, id)
	if tx.Error != nil {
		return emergency.EmergencyEntity{}, errors.New("fail emergency by id")
	}
	output := ModelToEntity(inputModel)
	return output, nil
}

// Update implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) Update(input emergency.EmergencyEntity, id uint) error {
	inputModel := EntityToModel(input)
	tx := repo.db.Model(&Emergency{}).Where("id=?", id).Updates(inputModel)
	if tx.Error != nil {
		return errors.New("update emergency fail")
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

// Delete implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) Delete(id uint) error {
	var inputModel Emergency
	tx := repo.db.Delete(&inputModel, id)
	if tx.Error != nil {
		return errors.New("fail delete emergency")
	}
	if tx.RowsAffected == 0 {
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
