package data

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	usernodejs "project-capston/features/UserNodeJs"
	"project-capston/features/emergency"
	"project-capston/helper"
	"time"

	"strconv"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type EmergencyData struct {
	db    *gorm.DB
	redis *redis.Client
}

// CreateUnit implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) CreateUnit(input emergency.UnitEntity) (uint, error) {
	inputModel:=UnitEntityToModel(input)
	tx:=repo.db.Create(&inputModel)
	if tx.Error != nil{
		return 0,errors.New("failed insert unit")
	}
	if tx.RowsAffected ==0{
		return 0, errors.New("row not affected")
	}
	return inputModel.ID,nil
}

// SumEmergency implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) SumEmergency() (int64, error) {
	var inputModel []Emergency
	tx := repo.db.Find(&inputModel)
	if tx.Error != nil {
		return 0, errors.New("error get all data")
	}
	count := tx.RowsAffected
	return count, nil
}

// SelectUser implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) SelectUser(id string, token string) (emergency.UserEntity, error) {

	data, err := usernodejs.GetByIdUser(id, token)
	if err != nil {
		return emergency.UserEntity{}, errors.New("user tidak ditemukan")
	}
	fmt.Println("data",data)
	dataUser := UserNodeToUser(data)
	dataEntity := UserToUserEntity(dataUser)
	return dataEntity, nil

}

// SendNotification implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) SendNotification(input helper.MessageGomailE) (string, error) {

	data, err := helper.SendGomailMessageE(input)
	if err != nil {
		return "", err
	}
	return data, nil
}

// ActionGmail implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) ActionGmail(input string) error {
	var inputModel HistoryAdmin
	inputModel.Status = input
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAll implements emergency.EmergencyDataInterface.
func (repo *EmergencyData) SelectAll(param emergency.QueryParams, token string,idCall uint,level string) (int64, []emergency.EmergencyEntity, error) {
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
	if level =="admin" ||level =="superadmin"{
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

	ctx := context.Background()
	var emergenciEntity []emergency.EmergencyEntity
	var userCallRedis User
	var userReceiverRedis User
	for i := 0; i < len(emergensiUser); i++ {
		for j := 0; j < len(emergensiUser); j++ {
			redisKey := fmt.Sprintf("user:%s", idCaller[j])

			cachedData, err := repo.redis.Get(ctx, redisKey).Result()
			if err == nil {
				var userRedis User
				if err := json.Unmarshal([]byte(cachedData), &userRedis); err != nil {
					return 0, nil, err
				}
				log.Println("Data ditemukan di Redis cache")
				userCallRedis = userRedis
			} else if err != redis.Nil {
				return 0, nil, err
			} else {
				data, _ := usernodejs.GetByIdUser(idCaller[j], token)
				user := UserNodeToUser(data)
				jsonData, err := json.Marshal(user)
				if err != nil {
					return 0, nil, err
				}
				errSet := repo.redis.Set(ctx, redisKey, jsonData, 24*time.Hour).Err()
				if errSet != nil {
					log.Println("Gagal menyimpan data ke Redis:", errSet)
				} else {
					log.Println("Data disimpan di Redis cache")
				}
				userCallRedis = user
			}

			idConv, _ := strconv.Atoi(idCaller[j])
			if uint(idConv) == emergensiUser[i].CallerID {
				emergensiUser[i].Caller = userCallRedis
			}
		}
		for k := 0; k < len(emergensiUser); k++ {
			redisKey := fmt.Sprintf("user:%s", idReceiver[k])

			cachedData, err := repo.redis.Get(ctx, redisKey).Result()
			if err == nil {
				var userRedis User
				if err := json.Unmarshal([]byte(cachedData), &userRedis); err != nil {
					return 0, nil, err
				}
				log.Println("Data ditemukan di Redis cache")
				userReceiverRedis = userRedis
			} else if err != redis.Nil {
				return 0, nil, err
			} else {
				data, _ := usernodejs.GetByIdUser(idReceiver[k], token)
				user := UserNodeToUser(data)
				jsonData, err := json.Marshal(user)
				if err != nil {
					return 0, nil, err
				}
				errSet := repo.redis.Set(ctx, redisKey, jsonData, 24*time.Hour).Err()
				if errSet != nil {
					log.Println("Gagal menyimpan data ke Redis:", errSet)
				} else {
					log.Println("Data disimpan di Redis cache")
				}
				userReceiverRedis = user
			}
			idConv, _ := strconv.Atoi(idReceiver[k])
			if uint(idConv) == emergensiUser[i].ReceiverID {
				emergensiUser[i].Receiver = userReceiverRedis
			}
		}
		emergenciEntity = append(emergenciEntity, EmergencyUserToEntity(emergensiUser[i]))

	}
	return totalEmergency, emergenciEntity, nil

	}
	if level == "user"{
		tx := query.Where("caller_id=?",idCall).Find(&inputModel)
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

	ctx := context.Background()
	var emergenciEntity []emergency.EmergencyEntity
	var userCallRedis User
	var userReceiverRedis User
	for i := 0; i < len(emergensiUser); i++ {
		for j := 0; j < len(emergensiUser); j++ {
			redisKey := fmt.Sprintf("user:%s", idCaller[j])

			cachedData, err := repo.redis.Get(ctx, redisKey).Result()
			if err == nil {
				var userRedis User
				if err := json.Unmarshal([]byte(cachedData), &userRedis); err != nil {
					return 0, nil, err
				}
				log.Println("Data ditemukan di Redis cache")
				userCallRedis = userRedis
			} else if err != redis.Nil {
				return 0, nil, err
			} else {
				data, _ := usernodejs.GetByIdUser(idCaller[j], token)
				user := UserNodeToUser(data)
				jsonData, err := json.Marshal(user)
				if err != nil {
					return 0, nil, err
				}
				errSet := repo.redis.Set(ctx, redisKey, jsonData, 24*time.Hour).Err()
				if errSet != nil {
					log.Println("Gagal menyimpan data ke Redis:", errSet)
				} else {
					log.Println("Data disimpan di Redis cache")
				}
				userCallRedis = user
			}

			idConv, _ := strconv.Atoi(idCaller[j])
			if uint(idConv) == emergensiUser[i].CallerID {
				emergensiUser[i].Caller = userCallRedis
			}
		}
		for k := 0; k < len(emergensiUser); k++ {
			redisKey := fmt.Sprintf("user:%s", idReceiver[k])

			cachedData, err := repo.redis.Get(ctx, redisKey).Result()
			if err == nil {
				var userRedis User
				if err := json.Unmarshal([]byte(cachedData), &userRedis); err != nil {
					return 0, nil, err
				}
				log.Println("Data ditemukan di Redis cache")
				userReceiverRedis = userRedis
			} else if err != redis.Nil {
				return 0, nil, err
			} else {
				data, _ := usernodejs.GetByIdUser(idReceiver[k], token)
				user := UserNodeToUser(data)
				jsonData, err := json.Marshal(user)
				if err != nil {
					return 0, nil, err
				}
				errSet := repo.redis.Set(ctx, redisKey, jsonData, 24*time.Hour).Err()
				if errSet != nil {
					log.Println("Gagal menyimpan data ke Redis:", errSet)
				} else {
					log.Println("Data disimpan di Redis cache")
				}
				userReceiverRedis = user
			}
			idConv, _ := strconv.Atoi(idReceiver[k])
			if uint(idConv) == emergensiUser[i].ReceiverID {
				emergensiUser[i].Receiver = userReceiverRedis
			}
		}
		emergenciEntity = append(emergenciEntity, EmergencyUserToEntity(emergensiUser[i]))

	}
	return totalEmergency, emergenciEntity, nil
	}

	return 0, nil, errors.New("failed get by id")

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
func (repo *EmergencyData) Insert(input emergency.EmergencyEntity) (uint, error) {
	inputModel := EntityToModel(input)
	tx := repo.db.Create(&inputModel)
	if tx.Error != nil {
		return 0, errors.New("failed create data emergency")
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("row not affected")
	}
	return inputModel.ID, nil
}

func New(db *gorm.DB, redis *redis.Client) emergency.EmergencyDataInterface {
	return &EmergencyData{
		db:    db,
		redis: redis,
	}
}
