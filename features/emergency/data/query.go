package data

import (
	"errors"
	usernodejs "project-capston/features/UserNodeJs"
	"project-capston/features/emergency"

	"strconv"

	"gorm.io/gorm"
)

type EmergencyData struct {
	db *gorm.DB
}

// ActionGmail implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) ActionGmail(input string) error {
	var inputModel HistoryAdmin
	inputModel.Status=input
	tx:=repo.db.Create(&inputModel)
	if tx.Error!=nil{
		return tx.Error
	}
	return nil
}

// SelectAll implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) SelectAll(param emergency.QueryParams, token string) (int64, []emergency.EmergencyEntity, error) {
	var inputModel []Emergency
	var totalEmergency int64

	query := repo.db

	if param.IsClassDashboard {
		offset := (param.Page - 1) * param.ItemsPerPage
		// if param.SearchName !=""{
		// 	query=query.Where("caller_id like ? or receiver_id=?","%"+param.SearchName+"%","%"+param.SearchName+"%")
		// }
		tx := query.Find(&inputModel)
		if tx.Error != nil {
			return 0, nil, errors.New("failed get count emergency")
		}
		totalEmergency = tx.RowsAffected
		query = query.Offset(offset).Limit(param.ItemsPerPage)
	}
	// if param.SearchName !=""{
	// 	query=query.Where("caller_id like ? or receiver_id=?","%"+param.SearchName+"%","%"+param.SearchName+"%")
	// }
	tx := query.Find(&inputModel)
	if tx.Error != nil {
		return 0, nil, errors.New("error get all emergency")
	}

	var emergensiUser []EmergencyUser
	for _, e := range inputModel {
		emergensiUser = append(emergensiUser, ModelToEmergencyUser(e))
	}

	var idReceiver []string
	for _, v := range emergensiUser {
		id := strconv.Itoa(int(v.ReceiverID))
		idReceiver = append(idReceiver, id)
	}

	var idCaller []string
	for _, v := range emergensiUser {
		id := strconv.Itoa(int(v.CallerID))
		idCaller = append(idCaller, id)
	}
	var emergenciEntity []emergency.EmergencyEntity
	for i := 0; i < len(emergensiUser); i++ {
		for j := 0; j < len(emergensiUser); j++ {
			data, _ := usernodejs.GetByIdUser(idCaller[j], token)
			idConv, _ := strconv.Atoi(idCaller[j])
			if uint(idConv) == emergensiUser[i].CallerID {
				user := UserNodeToUser(data)
				emergensiUser[i].Caller = user
			}
		}
		for k := 0; k < len(emergensiUser); k++ {
			data, _ := usernodejs.GetByIdUser(idReceiver[k], token)
			idConv, _ := strconv.Atoi(idReceiver[k])
			if uint(idConv) == emergensiUser[i].ReceiverID {
				user := UserNodeToUser(data)
				emergensiUser[i].Receiver = user
			}
		}
		emergenciEntity = append(emergenciEntity, EmergencyUserToEntity(emergensiUser[i]))

	}

	return totalEmergency, emergenciEntity, nil

}

// SelectById implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) SelectById(id uint, token string) (emergency.EmergencyEntity, error) {
	var inputModel Emergency
	tx := repo.db.First(&inputModel, id)
	if tx.Error != nil {
		return emergency.EmergencyEntity{}, errors.New("fail emergency by id")
	}

	idReceiver := strconv.Itoa(int(inputModel.ReceiverID))
	dataReceiver, errUserR := usernodejs.GetByIdUser(idReceiver, token)
	if errUserR != nil {
		return emergency.EmergencyEntity{}, errUserR
	}

	idCaller := strconv.Itoa(int(inputModel.CallerID))
	dataCaller, errUserC := usernodejs.GetByIdUser(idCaller, token)
	if errUserC != nil {
		return emergency.EmergencyEntity{}, errUserC
	}

	userCaller := UserNodeToUser(dataCaller)
	userEntityCaller := UserToUserEntity(userCaller)

	userReceiver := UserNodeToUser(dataReceiver)
	userEntityReceiver := UserToUserEntity(userReceiver)

	emergensyUser := ModelToEmergencyUser(inputModel)

	output := EmergencyUserToEntity(emergensyUser)
	output.Caller = userEntityCaller
	output.Receiver = userEntityReceiver
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
