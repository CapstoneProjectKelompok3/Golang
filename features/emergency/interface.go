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

type EmergencyDataInterface interface{
	Insert(input EmergencyEntity)(error)
	Delete(id uint)(error)
}

type EmergencyServiceInterface interface{
	Add(input EmergencyEntity)(error)
	Delete(id uint)(error)
}