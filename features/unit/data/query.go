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

// CreateUnit implements unit.UnitDataInterface.
// func (repo *UnitData) CreateUnit() (uint, error) {
// 	// helper.InputUnit(input)
// 	// repo.db.Create()
// }

// // CreateUnitHistory implements unit.UnitDataInterface.
// func (*UnitData) CreateUnitHistory(id uint, input unit.UnitHistoryEntity) (uint, error) {
// 	panic("unimplemented")
// }

func (repo *UnitData) CreateUnit(idEmergency uint, tipe []string, count []int) error {

	for i := 0; i < len(tipe); i++ {
		if count[i] != 0 {
			unitData := Unit{
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

func (repo *UnitData) CreateUnitHistori(idEmergency uint) error {
	var inputUnit []Unit
	txx := repo.db.Where("emergencies_id=?", idEmergency).Find(&inputUnit)
	if txx.Error != nil {
		return errors.New("failed get unit history")
	}

	tx := repo.db.Begin()
	for _, v := range inputUnit {
		for j := 0; j < v.SumOfUnit; j++ {
			inputModel := UnitHistory{
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

	userEmergencies := UserNodeToUser(dataEmergencies)
	userEntityEmergencies := UserToUserEntity(userEmergencies)

	unitUser := ModelToUnitUser(inputModel)

	output := UnitUserToEntity(unitUser)
	output.Emergencies = userEntityEmergencies
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

// Insert implements unit.UnitDataInterface.
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
