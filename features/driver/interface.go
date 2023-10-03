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
	DrivingStatus string  `gorm:"type:enum('on_ready','on_demand','on_trip','on_finished');column:driving_status;default:on_ready"`
	VehicleID     uint    `validate:"required"`
	Latitude      float64 `validate:"required"`
	Longitude     float64 `validate:"required"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
}

type DriverDataInterface interface {
	Insert(input Core) error
}

type DriverServiceInterface interface {
	Create(input Core) error
}
