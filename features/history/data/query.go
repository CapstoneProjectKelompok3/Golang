package data

import (
	"errors"
	usernodejs "project-capston/features/UserNodeJs"
	"project-capston/features/history"
	"strconv"

	"gorm.io/gorm"
)

type HistoryData struct {
	db *gorm.DB
}

// SelectById implements history.HistoryDataInterface.
func (repo *HistoryData) SelectById(id uint, token string) (history.HistoryEntity, error) {
	var inputModel History
	tx := repo.db.First(&inputModel, id)
	if tx.Error != nil {
		return history.HistoryEntity{}, errors.New("fail history by id")
	}
	idUnit := strconv.Itoa(int(inputModel.UnitID))
	dataUnit, errUserU := usernodejs.GetByIdUser(idUnit, token)
	if errUserU != nil {
		return history.HistoryEntity{}, errUserU
	}

	idDriver := strconv.Itoa(int(inputModel.DriverID))
	dataDriver, errUserD := usernodejs.GetByIdUser(idDriver, token)
	if errUserD != nil {
		return history.HistoryEntity{}, errUserD
	}

	userUnit := UserNodeToUser(dataUnit)
	userEntityUnit := UserToUserEntity(userUnit)

	userDriver := UserNodeToUser(dataDriver)
	userEntityDriver := UserToUserEntity(userDriver)

	historyUser := ModelToUnitUser(inputModel)

	output := HistoryUserToEntity(historyUser)
	output.Unit = userEntityUnit
	output.Driver = userEntityDriver
	return output, nil
}

// SelectAll implements history.HistoryDataInterface.
func (repo *HistoryData) SelectAll(param history.QueryParams, token string) (int64, []history.HistoryEntity, error) {
	var inputModel []History
	var totalHistory int64

	query := repo.db

	if param.IsClassDashboard {
		offset := (param.Page - 1) * param.ItemsPerPage
		// if param.SearchName !=""{
		// 	query=query.Where("caller_id like ? or receiver_id=?","%"+param.SearchName+"%","%"+param.SearchName+"%")
		// }
		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get count history")
		}
		totalHistory = tx.RowsAffected
		query = query.Offset(offset).Limit(param.ItemsPerPage)
	}
	// if param.SearchName !=""{
	// 	query=query.Where("caller_id like ? or receiver_id=?","%"+param.SearchName+"%","%"+param.SearchName+"%")
	// }
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("error get all history")
	}

	var historiUser []HistoryUser
	for _, e := range inputModel {
		historiUser = append(historiUser, ModelToUnitUser(e))
	}

	var idUnit []string
	for _, v := range historiUser {
		id := strconv.Itoa(int(v.UnitID))
		idUnit = append(idUnit, id)
	}

	var idDriver []string
	for _, v := range historiUser {
		id := strconv.Itoa(int(v.DriverID))
		idDriver = append(idDriver, id)
	}
	var historiEntity []history.HistoryEntity
	for i := 0; i < len(historiUser); i++ {
		for j := 0; j < len(historiUser); j++ {
			data, _ := usernodejs.GetByIdUser(idUnit[j], token)
			idConv, _ := strconv.Atoi(idUnit[j])
			if uint(idConv) == historiUser[i].UnitID {
				user := UserNodeToUser(data)
				historiUser[i].Unit = user
			}
		}
		for k := 0; k < len(historiUser); k++ {
			data, _ := usernodejs.GetByIdUser(idDriver[k], token)
			idConv, _ := strconv.Atoi(idDriver[k])
			if uint(idConv) == historiUser[i].DriverID {
				user := UserNodeToUser(data)
				historiUser[i].Driver = user
			}
		}
		historiEntity = append(historiEntity, UnitUserToEntity(historiUser[i]))

	}

	return totalHistory, historiEntity, nil

}

// Update implements history.HistoryDataInterface.
func (repo *HistoryData) Update(input history.HistoryEntity, id uint) error {
	inputModel := EntityToModel(input)
	tx := repo.db.Model(&History{}).Where("id=?", id).Updates(inputModel)
	if tx.Error != nil {
		return errors.New("update history fail")
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

// Delete implements history.HistoryDataInterface.
func (repo *HistoryData) Delete(id uint) error {
	var inputModel History
	tx := repo.db.Delete(&inputModel, id)
	if tx.Error != nil {
		return errors.New("fail delete history")
	}
	if tx.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}

// Insert implements hisytor.HistoryDataInterface.
func (repo *HistoryData) Insert(input history.HistoryEntity) error {
	inputModel := EntityToModel(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return errors.New("failed create data history")
	}
	if tx.RowsAffected == 0 {
		return errors.New("row not affected")
	}
	return nil
}

func New(db *gorm.DB) history.HistoryDataInterface {
	return &HistoryData{
		db: db,
	}
}
