package unit

import "time"

type UnitEntity struct {
	Id            uint
	CreateAt      time.Time
	UpdateAt      time.Time
	DeleteAt      time.Time
	EmergenciesID uint   `validate:"required"`
	GovermentType string `gorm:"type:enum('Polisi','Rumah Sakit','DISHUB','SAR','Damkar');column:goverment_type"`
	SumOfUnit     int    `validate:"required"`
	Emergencies   UserEntity
}

type UserEntity struct {
	ID    int
	Name  string
	Level string
}

type QueryParams struct {
	Page             int
	ItemsPerPage     int
	SearchName       string
	IsClassDashboard bool
}

type UnitDataInterface interface {
	Insert(input UnitEntity) error
	Delete(id uint) error
	Update(input UnitEntity, id uint) error
	SelectById(id uint, token string) (UnitEntity, error)
	SelectAll(param QueryParams, token string) (int64, []UnitEntity, error)
}

type UnitServiceInterface interface {
	Add(input UnitEntity) error
	Delete(id uint) error
	Edit(input UnitEntity, id uint) error
	GetById(id uint, token string) (UnitEntity, error)
	GetAll(param QueryParams, token string) (bool, []UnitEntity, error)
}
