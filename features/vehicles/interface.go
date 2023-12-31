package vehicles

import "time"

type VehicleEntity struct {
	Id          uint
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAt    time.Time
	GovermentID uint `validate:"required"`
	Plate       string `validate:"required"`
	Status      bool
	Goverment GovernmentEntity
}
type GovernmentEntity struct {
	ID        uint   
	Name      string 
	Type      string 
	Address   string
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
type VehicleDataInterface interface{
	Insert(input VehicleEntity)(error)
	Update(input VehicleEntity,id uint)error
	SelectById(id uint)(VehicleEntity,error)
	SelectAll()([]VehicleEntity,error)
	Delete(id uint)error
}

type VehicleServiceInterface interface{
	Add(input VehicleEntity,level string)(error)
	Edit(input VehicleEntity,id uint,level string)error
	GetById(id uint)(VehicleEntity,error)
	GetAll()([]VehicleEntity,error)
	Delete(id uint)error
}