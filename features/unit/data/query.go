package data

import (
	"errors"
	usernodejs "project-capston/features/UserNodeJs"
	"project-capston/features/unit"
	"strconv"

	"gorm.io/gorm"
)

type UnitData struct {
	db *gorm.DB
}

// SelectById implements unit.UnitDataInterface.
func (repo *UnitData) SelectById(id uint, token string) (unit.UnitEntity, error) {
	var inputModel Unit
	tx := repo.db.First(&inputModel, id)
	if tx.Error != nil {
		return unit.UnitEntity{}, errors.New("fail unit by id")
	}
	idEmergencies := strconv.Itoa(int(inputModel.EmergenciesID))
	dataEmergencies, errUserE := usernodejs.GetByIdUser(idEmergencies, token)
	if errUserE != nil {
		return unit.UnitEntity{}, errUserE
	}

	idVehicle := strconv.Itoa(int(inputModel.VehicleID))
	dataVehicle, errUserV := usernodejs.GetByIdUser(idVehicle, token)
	if errUserV != nil {
		return unit.UnitEntity{}, errUserV
	}

	userEmergencies := UserNodeToUser(dataEmergencies)
	userEntityEmergencies := UserToUserEntity(userEmergencies)

	userVehicle := UserNodeToUser(dataVehicle)
	userEntityVehicle := UserToUserEntity(userVehicle)

	unitUser := ModelToUnitUser(inputModel)

	output := UnitUserToEntity(unitUser)
	output.Emergencies = userEntityEmergencies
	output.Vehicle = userEntityVehicle
	return output, nil
}

// SelectAll implements unit.UnitDataInterface.
func (repo *UnitData) SelectAll(param unit.QueryParams, token string) (int64, []unit.UnitEntity, error) {
	var inputModel []Unit
	var totalUnit int64

	query := repo.db

	if param.IsClassDashboard {
		offset := (param.Page - 1) * param.ItemsPerPage
		// if param.SearchName !=""{
		// 	query=query.Where("caller_id like ? or receiver_id=?","%"+param.SearchName+"%","%"+param.SearchName+"%")
		// }
		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get count unit")
		}
		totalUnit = tx.RowsAffected
		query = query.Offset(offset).Limit(param.ItemsPerPage)
	}
	// if param.SearchName !=""{
	// 	query=query.Where("caller_id like ? or receiver_id=?","%"+param.SearchName+"%","%"+param.SearchName+"%")
	// }
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("error get all unit")
	}

	var uniitUser []UnitUser
	for _, e := range inputModel {
		uniitUser = append(uniitUser, ModelToUnitUser(e))
	}

	var idEmergencies []string
	for _, v := range uniitUser {
		id := strconv.Itoa(int(v.EmergenciesID))
		idEmergencies = append(idEmergencies, id)
	}

	var idVehicle []string
	for _, v := range uniitUser {
		id := strconv.Itoa(int(v.VehicleID))
		idVehicle = append(idVehicle, id)
	}
	var uniitEntity []unit.UnitEntity
	for i := 0; i < len(uniitUser); i++ {
		for j := 0; j < len(uniitUser); j++ {
			data, _ := usernodejs.GetByIdUser(idEmergencies[j], token)
			idConv, _ := strconv.Atoi(idEmergencies[j])
			if uint(idConv) == uniitUser[i].EmergenciesID {
				user := UserNodeToUser(data)
				uniitUser[i].Emergencies = user
			}
		}
		for k := 0; k < len(uniitUser); k++ {
			data, _ := usernodejs.GetByIdUser(idVehicle[k], token)
			idConv, _ := strconv.Atoi(idVehicle[k])
			if uint(idConv) == uniitUser[i].VehicleID {
				user := UserNodeToUser(data)
				uniitUser[i].Vehicle = user
			}
		}
		uniitEntity = append(uniitEntity, UnitUserToEntity(uniitUser[i]))

	}

	return totalUnit, uniitEntity, nil

}

// Update implements unit.UnitDataInterface.
func (repo *UnitData) Update(input unit.UnitEntity, id uint) error {
	inputModel := EntityToModel(input)
	tx := repo.db.Model(&Unit{}).Where("id=?", id).Updates(inputModel)
	if tx.Error != nil {
		return errors.New("update unit fail")
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

// Delete implements unit.UnitDataInterface.
func (repo *UnitData) Delete(id uint) error {
	var inputModel Unit
	tx := repo.db.Delete(&inputModel, id)
	if tx.Error != nil {
		return errors.New("fail delete unit")
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

// Insert implements unut.UnitDataInterface.
func (repo *UnitData) Insert(input unit.UnitEntity) error {
	inputModel := EntityToModel(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return errors.New("failed create data unit")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

func New(db *gorm.DB) unit.UnitDataInterface {
	return &UnitData{
		db: db,
	}
}
