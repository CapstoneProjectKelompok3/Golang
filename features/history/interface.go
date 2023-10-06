package history

import "time"

type HistoryEntity struct {
	Id       uint
	UnitID   uint
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	DriverID uint
	Status   string `gorm:"type:enum('Accepted','Rejected');default:'rejected';column:status"`
	Reason   string
	Unit     UserEntity
	Driver   UserEntity
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

type HistoryDataInterface interface {
	Insert(input HistoryEntity) error
	Delete(id uint) error
	Update(input HistoryEntity, id uint) error
	SelectById(id uint, token string) (HistoryEntity, error)
	SelectAll(param QueryParams, token string) (int64, []HistoryEntity, error)
}

type HistoryServiceInterface interface {
	Add(input HistoryEntity) error
	Delete(id uint) error
	Edit(input HistoryEntity, id uint) error
	GetById(id uint, token string) (HistoryEntity, error)
	GetAll(param QueryParams, token string) (bool, []HistoryEntity, error)
}
