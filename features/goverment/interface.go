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
}

type GovernmentServiceInterface interface {
	Create(input Core) error
}
