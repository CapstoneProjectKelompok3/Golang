package emergency

import (
	"project-capston/helper"
	"time"
)

type EmergencyEntity struct {
	Id         	uint
	Name string
	CallerID   	uint
	ReceiverID 	uint `validate:"required"`
	Latitude   	float64 `validate:"required"`
	Longitude  	float64 `validate:"required"`
	CreateAt   	time.Time
	UpdateAt 	time.Time
	DeleteAt 	time.Time
	Caller     UserEntity
	Receiver   UserEntity
	IsClose bool
}
type UnitEntity struct {
	Id         	uint
	CreateAt   	time.Time
	UpdateAt 	time.Time
	DeleteAt 	time.Time
	EmergenciesID uint
	VehicleID     uint
	GovermentType string
	SumOfUnit int 
}
type UserEntity struct{
	ID        		int
	Name 			string
	Email			string
	Level           string
	EmailActive     bool
}

type QueryParams struct {
	Page            int
	ItemsPerPage    int
	SearchName      string
	IsClassDashboard bool
}

type EmergencyDataInterface interface{
	Insert(input EmergencyEntity)(uint,error)
	Delete(id uint)(error)
	Update(input EmergencyEntity, id uint)(error)
	SelectById(id uint,token string)(EmergencyEntity,error)
	SelectAll(param QueryParams,token string,idCall uint,level string)(int64, []EmergencyEntity,error)
	SendNotification(input helper.MessageGomailE)(string,error)
	ActionGmail(input string)(error)
	SelectUser(id string,token string) (UserEntity, error)
	SumEmergency()(int64,error)
	IncloseEmergency(idEmergency uint)error
}

type EmergencyServiceInterface interface{
	Add(input EmergencyEntity,token string)(uint,error)
	Delete(id uint)(error)
	Edit(input EmergencyEntity,id uint,level string,idUser uint)error
	GetById(id uint,token string)(EmergencyEntity,error)
	GetAll(param QueryParams,token string,idCall uint,level string)(bool,[]EmergencyEntity,error)
	ActionGmail(input string)error
	SumEmergency()(int64,error)
	IncloseEmergency(idEmergency uint)error
}