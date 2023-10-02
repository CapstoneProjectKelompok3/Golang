package goverment

import "time"

type Core struct {
	ID        uint
	Name      string `validate:"required"`
	Type      string `validate:"required"`
	Address   string
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Location struct {
	ID        uint
	Name      string
	Latitude  float64
	Longitude float64
	Distance  float64
}

type GovernmentDataInterface interface {
	Insert(input Core) error
	SelectAll(pageNumber int, pageSize int) ([]Core, error)
	Select(id uint) (Core, error)
	Update(id uint, input Core) error
	Delete(id uint) error

	//get nearest location
	SelectNearestLocation(latitude float64, longitude float64) ([]Location, error)
}

type GovernmentServiceInterface interface {
	Create(input Core) error
	GetAll(pageNumber int, pageSize int) ([]Core, error)
	GetById(id uint) (Core, error)
	EditById(id uint, input Core) error
	DeleteById(id uint) error

	//get nearest location
	GetNearestLocation(latitude float64, longitude float64) ([]Location, error)
}
