package data

import (
	usernodejs "project-capston/features/UserNodeJs"
	"project-capston/features/history"
	"time"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	UnitID   uint
	DriverID uint
	Status   string
}

type HistoryUser struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	UnitID    uint
	DriverID  uint
	Status    string `gorm:"type:enum('Accepted','Rejected');default:'rejected';column:status"`
	Unit      User
	Driver    User
}

type User struct {
	ID    int
	Name  string
	Level string
}

func UserToUserEntity(user User) history.UserEntity {
	return history.UserEntity{
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

func ModelToUnitUser(history History) HistoryUser {
	return HistoryUser{
		ID:        history.ID,
		UnitID:    history.UnitID,
		DriverID:  history.DriverID,
		Status:    history.Status,
		CreatedAt: history.CreatedAt,
		UpdatedAt: history.UpdatedAt,
		DeletedAt: history.DeletedAt.Time,
	}
}
func UnitUserToEntity(histori HistoryUser) history.HistoryEntity {
	return history.HistoryEntity{
		Id:       histori.ID,
		UnitID:   histori.UnitID,
		DriverID: histori.DriverID,
		Status:   histori.Status,
		CreateAt: histori.CreatedAt,
		UpdateAt: histori.UpdatedAt,
		DeleteAt: histori.DeletedAt,
	}
}

func HistoryUserToEntity(histori HistoryUser) history.HistoryEntity {
	return history.HistoryEntity{
		Id:       histori.ID,
		UnitID:   histori.UnitID,
		DriverID: histori.DriverID,
		Status:   histori.Status,
		CreateAt: histori.CreatedAt,
		UpdateAt: histori.UpdatedAt,
		DeleteAt: histori.DeletedAt,
		Unit:     UserToUserEntity(histori.Unit),
		Driver:   UserToUserEntity(histori.Driver),
	}
}

func ModelToEntity(histori History) history.HistoryEntity {
	return history.HistoryEntity{
		Id:       histori.ID,
		UnitID:   histori.UnitID,
		DriverID: histori.DriverID,
		Status:   histori.Status,
		CreateAt: histori.CreatedAt,
		UpdateAt: histori.UpdatedAt,
		DeleteAt: histori.DeletedAt.Time,
	}
}

func EntityToModel(history history.HistoryEntity) History {
	return History{
		UnitID:   history.UnitID,
		DriverID: history.DriverID,
		Status:   history.Status,
	}
}
