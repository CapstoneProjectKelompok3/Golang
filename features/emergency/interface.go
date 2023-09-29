package emergency

import "time"

type EmergencyEntity struct {
	Id         	uint
	CallerID   	uint `validate:"required"`
	ReceiverID 	uint `validate:"required"`
	Latitude   	float64 `validate:"required"`
	Longitude  	float64 `validate:"required"`
	CreateAt   	time.Time
	UpdateAt 	time.Time
	DeleteAt 	time.Time
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
	SelectById(id uint)(EmergencyEntity,error)
	SelectAll(param QueryParams)(int64, []EmergencyEntity,error)
}

type EmergencyServiceInterface interface{
	Add(input EmergencyEntity)(error)
	Delete(id uint)(error)
	Edit(input EmergencyEntity,id uint)error
	GetById(id uint)(EmergencyEntity,error)
	GetAll(param QueryParams)(bool,[]EmergencyEntity,error)
}