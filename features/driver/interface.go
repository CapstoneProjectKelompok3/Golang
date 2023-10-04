package driver

import "time"

type Core struct {
	Id            uint
	Fullname      string `validate:"required"`
	Email         string `validate:"required,email"`
	Password      string `validate:"required"`
	Token         string
	GovermentID   uint `validate:"required"`
	Status        bool
	DrivingStatus string  `gorm:"type:enum('on_ready','on_demand','on_trip','on_finished','on_cancle');column:driving_status;default:on_ready"`
	VehicleID     uint    `validate:"required"`
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
	GovermentName string
	GovermentType string
	Status        bool
	DrivingStatus string
	VehicleID     uint
	Latitude      float64
	Longitude     float64
	Distance      float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type DriverDataInterface interface {
	Insert(input Core) error
	SelectAll(pageNumber int, pageSize int) ([]DriverCore, error)
	Login(email string, password string) (dataLogin Core, err error)
	KerahkanDriver(lat string, long string, police int, hospital int, firestation int, dishub int, SAR int) ([]DriverCore, error)
	// Logout(email string, password string) (dataLogin Core, err error)
	SelectProfile(id int) (DriverCore, error)
	AcceptOrRejectOrder(IsAccepted bool, idDriver int) error
	DriverOnTrip(id int, lat float64, long float64) (DriverCore, error)
}

type DriverServiceInterface interface {
	Create(input Core) error
	GetAll(pageNumber int, pageSize int) ([]DriverCore, error)
	Login(email string, password string) (dataLogin Core, token string, err error)
	KerahkanDriver(lat string, long string, police int, hospital int, firestation int, dishub int, SAR int) ([]DriverCore, error)
	// Logout(email string, password string) (dataLogin Core, err error)
	GetProfile(id int) (DriverCore, error)
	AcceptOrRejectOrder(IsAccepted bool, idDriver int) error
	DriverOnTrip(id int, lat float64, long float64) (DriverCore, error)
}
