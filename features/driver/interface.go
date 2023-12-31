package driver

import (
	"time"
)

type Core struct {
	Id            uint
	Fullname      string `validate:"required"`
	Email         string `validate:"required,email"`
	Password      string `validate:"required"`
	Token         string
	GovermentID   uint `validate:"required"`
	Status        bool
	DrivingStatus string `gorm:"type:enum('on_ready','on_demand','on_trip','on_finished','on_cancle');column:driving_status;default:on_ready"`
	VehicleID     uint   `validate:"required"`
	EmergenciesID uint
	EmergencyName string
	Latitude      float64 `validate:"required"`
	Longitude     float64 `validate:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type DriverCore struct {
	Id            uint
	Fullname      string
	Email         string
	Password      string
	Token         string
	EmergencyID   uint
	EmergencyName string
	GovermentName string
	GovermentType string
	Status        bool
	DrivingStatus string
	VehicleID     uint
	EmergenciesID uint
	Latitude      float64
	Longitude     float64
	Distance      float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type DriverCoreStatus struct {
	Id            uint
	Fullname      string
	Status        bool
	DrivingStatus string
}

type DriverDataInterface interface {
	Insert(input Core) error
	SelectAll(pageNumber int, pageSize int) ([]DriverCore, error)
	Login(email string, password string) (dataLogin Core, err error)
	KerahkanDriver(lat string, long string, police int, hospital int, firestation int, dishub int, SAR int, emergency_id int) ([]DriverCore, error)
	// Logout(email string, password string) (dataLogin Core, err error)
	SelectProfile(id int) (DriverCore, error)
	AcceptOrRejectOrder(IsAccepted bool, idDriver int) error
	DriverOnTrip(id int, lat float64, long float64) (DriverCore, error)
	FinishTrip(id int) error
	Logout(id int) error
	SelectCountDriver() (int64, error)
	Delete(id uint) error
	CreateUnit(idEmergency uint, tipe []string, count []int) error
	CreateUnitHistori(idEmergency uint) error
	UpdateHistoryUnit(idDriver uint, idUnitHistori uint) error
	SelectUnit(idEmergenci uint) ([]uint, []string, error)
	SelectHistori(idUnit uint) (uint, error)
	UpdateFinish(id uint, idE uint) error
	// IsCloseEmergency(status bool,idEmergency uint)error
	// SelectAllEmergencyInUnit(idEmergency uint)(bool,error)

}

type DriverServiceInterface interface {
	Create(input Core) error
	GetAll(pageNumber int, pageSize int) ([]DriverCore, error)
	Login(email string, password string) (dataLogin Core, token string, err error)
	KerahkanDriver(id uint, lat string, long string, police int, hospital int, firestation int, dishub int, SAR int, emergency_id int) ([]DriverCore, error)
	// Logout(email string, password string) (dataLogin Core, err error)
	GetProfile(id int) (DriverCore, error)
	AcceptOrRejectOrder(idEmergenci uint, IsAccepted bool, idDriver int) error
	DriverOnTrip(id int, lat float64, long float64) (DriverCore, error)
	FinishTrip(id int, idE uint) error
	Logout(id int) error
	GetCountDriver() (int64, error)
	Delete(id uint) error
}
