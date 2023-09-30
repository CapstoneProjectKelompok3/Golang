package data

import (
	usernodejs "project-capston/features/UserNodeJs"
	"project-capston/features/emergency"
	"time"

	"gorm.io/gorm"
)

type Emergency struct {
	gorm.Model
	CallerID   uint
	ReceiverID uint
	Latitude   float64
	Longitude  float64
}

type EmergencyUser struct{
	ID        		uint			
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		time.Time
	CallerID   uint
	ReceiverID uint
	Latitude   float64
	Longitude  float64
	Caller     User
	Receiver   User
}

type User struct{
	ID        		int
	Name 			string	
	Level           string
}
func UserToUserEntity(user User)emergency.UserEntity{
	return emergency.UserEntity{
		ID:    user.ID,
		Name:  user.Name,
		Level: user.Level,
	}
}

func UserNodeToUser(user usernodejs.User)User{
	return User{
		ID:    user.ID,
		Name:  user.Username,
		Level: user.Level,
	}
}

func UserEntityToEntity(user emergency.UserEntity)emergency.UserEntity{
	return emergency.UserEntity{
		ID:    user.ID,
		Name:  user.Name,
		Level: user.Level,
	}
}

func ModelToEmergencyUser(emergency Emergency)EmergencyUser{
	return EmergencyUser{
		ID:         emergency.ID,
		CallerID:   emergency.CallerID,
		ReceiverID: emergency.ReceiverID,
		Latitude:   emergency.Latitude,
		Longitude:  emergency.Longitude,
		CreatedAt:   emergency.CreatedAt,
		UpdatedAt:   emergency.UpdatedAt,
		DeletedAt:   emergency.DeletedAt.Time,
	}
}
func EmergencyUserToEntity(emergenci EmergencyUser)emergency.EmergencyEntity{
	return emergency.EmergencyEntity{
		Id:         emergenci.ID,
		CallerID:   emergenci.CallerID,
		ReceiverID: emergenci.ReceiverID,
		Latitude:   emergenci.Latitude,
		Longitude:  emergenci.Longitude,
		CreateAt:   emergenci.CreatedAt,
		UpdateAt:   emergenci.UpdatedAt,
		DeleteAt:   emergenci.DeletedAt,
		Caller: UserToUserEntity(emergenci.Caller),	
		Receiver: UserToUserEntity(emergenci.Receiver),	
	}
}

func ModelToEntity(emergenci Emergency)emergency.EmergencyEntity{
	return emergency.EmergencyEntity{
		Id:         emergenci.ID,
		CallerID:   emergenci.CallerID,
		ReceiverID: emergenci.ReceiverID,
		Latitude:   emergenci.Latitude,
		Longitude:  emergenci.Longitude,
		CreateAt:   emergenci.CreatedAt,
		UpdateAt:   emergenci.UpdatedAt,
		DeleteAt:   emergenci.DeletedAt.Time,
	}
}

func EntityToModel(emergenci emergency.EmergencyEntity)Emergency{
	return Emergency{
		CallerID:   emergenci.CallerID,
		ReceiverID: emergenci.ReceiverID,
		Latitude:   emergenci.Latitude,
		Longitude:  emergenci.Longitude,
	}
}