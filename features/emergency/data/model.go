package data

import (
	usernodejs "project-capston/features/UserNodeJs"
	"project-capston/features/emergency"
	"project-capston/helper"
	"time"

	"gorm.io/gorm"
)

type Emergency struct {
	gorm.Model
	Name string
	CallerID   uint
	ReceiverID uint
	Latitude   float64
	Longitude  float64
	IsClose bool
}

type HistoryAdmin struct{
	gorm.Model
	AdminId uint
	Status string
}

type EmergencyUser struct{
	ID        		uint			
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		time.Time
	Name		string
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
	Email 			string
	EmailActive		bool
}

type MessageGomail struct {
    EmailReceiver string
    Sucject       string
    Content       string
    Name          string
    Email         string
}

func MessageGomailToMessage(gmail helper.MessageGomailE)MessageGomail{
	return MessageGomail{
		EmailReceiver: gmail.EmailReceiver,
		Sucject:       gmail.Sucject,
		Content:       gmail.Content,
		Name:          gmail.Name,
		Email:         gmail.Name,
	}
}
func UserToUserEntity(user User)emergency.UserEntity{
	return emergency.UserEntity{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Level: user.Level,
		EmailActive: user.EmailActive,
	}
}

func UserNodeToUser(user usernodejs.User)User{
	return User{
		ID:    user.ID,
		Name:  user.Username,
		Level: user.Level,
		Email: user.Email,
		EmailActive: user.EmailActive,
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
		CreatedAt:  emergency.CreatedAt,
		UpdatedAt:  emergency.UpdatedAt,
		DeletedAt:  emergency.DeletedAt.Time,
		Name:       emergency.Name,
		CallerID:   emergency.CallerID,
		ReceiverID: emergency.ReceiverID,
		Latitude:   emergency.Latitude,
		Longitude:  emergency.Longitude,
	}
}
func EmergencyUserToEntity(emergenci EmergencyUser)emergency.EmergencyEntity{
	return emergency.EmergencyEntity{
		Id:         emergenci.ID,
		CallerID:   emergenci.CallerID,
		Name:       emergenci.Name,	
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
		Name: emergenci.Name,
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
		Name: emergenci.Name,
		Latitude:   emergenci.Latitude,
		Longitude:  emergenci.Longitude,
	}
}