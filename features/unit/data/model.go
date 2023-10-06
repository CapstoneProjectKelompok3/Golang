package data

import (
	usernodejs "project-capston/features/UserNodeJs"
	"project-capston/features/emergency"
	"project-capston/features/unit"
	"time"

	"gorm.io/gorm"
)

type Unit struct {
	gorm.Model
	EmergenciesID uint
	VehicleID     uint
	GovermentType string `gorm:"type:enum('hospital','police','firestation','dishub','SAR');column:type;default:hospital"`
	SumOfUnit int 
}

type UnitHistory struct {
	gorm.Model
	UnitID   uint
	DriverID uint
	Status  string
	Reason string
}

type UnitUser struct {
	ID            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
	EmergenciesID uint
	VehicleID     uint
	Emergencies   User
	Vehicle       User
}

type User struct {
	ID    int
	Name  string
	Level string
}

func UserToUserEntity(user User) unit.UserEntity {
	return unit.UserEntity{
		ID:    user.ID,
		Name:  user.Name,
		Level: user.Level,
	}
}

func UserNodeToUser(user usernodejs.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Username,
		Level: user.Level,
	}
}

func UserEntityToEntity(user emergency.UserEntity) emergency.UserEntity {
	return emergency.UserEntity{
		ID:    user.ID,
		Name:  user.Name,
		Level: user.Level,
	}
}

func ModelToUnitUser(unit Unit) UnitUser {
	return UnitUser{
		ID:            unit.ID,
		EmergenciesID: unit.EmergenciesID,
		VehicleID:     unit.VehicleID,
		CreatedAt:     unit.CreatedAt,
		UpdatedAt:     unit.UpdatedAt,
		DeletedAt:     unit.DeletedAt.Time,
	}
}
func UnitUserToEntity(uniit UnitUser) unit.UnitEntity {
	return unit.UnitEntity{
		Id:            uniit.ID,
		EmergenciesID: uniit.EmergenciesID,
		VehicleID:     uniit.VehicleID,
		CreateAt:      uniit.CreatedAt,
		UpdateAt:      uniit.UpdatedAt,
		DeleteAt:      uniit.DeletedAt,
		Emergencies:   UserToUserEntity(uniit.Emergencies),
		Vehicle:       UserToUserEntity(uniit.Vehicle),
	}
}

func ModelToEntity(uniit Unit) unit.UnitEntity {
	return unit.UnitEntity{
		Id:            uniit.ID,
		EmergenciesID: uniit.EmergenciesID,
		VehicleID:     uniit.VehicleID,
		CreateAt:      uniit.CreatedAt,
		UpdateAt:      uniit.UpdatedAt,
		DeleteAt:      uniit.DeletedAt.Time,
	}
}

func EntityToModel(unit unit.UnitEntity) Unit {
	return Unit{
		EmergenciesID: unit.EmergenciesID,
		VehicleID:     unit.VehicleID,
	}
}
