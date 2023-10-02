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

type GovernmentDataInterface interface {
	Insert(input Core) error
	SelectAll(pageNumber int, pageSize int) ([]Core, error)
	Select(id uint) (Core, error)
}

type GovernmentServiceInterface interface {
	Create(input Core) error
	GetAll(pageNumber int, pageSize int) ([]Core, error)
	GetById(id uint) (Core, error)
}
