package emergency

import "time"

type EmergencyEntity struct {
	Id         	uint
	Name string
	CallerID   	uint `validate:"required"`
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

type UserEntity struct{
	ID        		int
	Name 			string	
	Level           string
}

type QueryParams struct {
	Page            int
	ItemsPerPage    int
	SearchName      string
	IsClassDashboard bool
}

type EmergencyDataInterface interface{
	Insert(input EmergencyEntity)(error)
	Delete(id uint)(error)
	Update(input EmergencyEntity, id uint)(error)
	SelectById(id uint,token string)(EmergencyEntity,error)
	SelectAll(param QueryParams,token string)(int64, []EmergencyEntity,error)
	ActionGmail(input string)error
}

type EmergencyServiceInterface interface{
	Add(input EmergencyEntity)(error)
	Delete(id uint)(error)
	Edit(input EmergencyEntity,id uint)error
	GetById(id uint,token string)(EmergencyEntity,error)
	GetAll(param QueryParams,token string)(bool,[]EmergencyEntity,error)
	ActionGmail(input string)error
}