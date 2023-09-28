package data

import (
	"project-capston/features/emergency"

	"gorm.io/gorm"
)

type Emergency struct {
	gorm.Model
	CallerID   uint
	ReceiverID uint
	Latitude   float64
	Longitude  float64
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