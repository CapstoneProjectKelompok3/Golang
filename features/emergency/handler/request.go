package handler

import "project-capston/features/emergency"

type EmergencyRequest struct {
	Name string `json:"name" form:"name"`
	CallerID   uint    `json:"caller_id" form:"caller_id"`
	ReceiverID uint    `json:"receiver_id" form:"receiver_id"`
	Latitude   float64 `json:"latitude" form:"latitude"`
	Longitude  float64 `json:"longitude" form:"longitude"`
	IsClose bool
}

func RequestToEntity(data EmergencyRequest) emergency.EmergencyEntity{
	return emergency.EmergencyEntity{
		CallerID:   data.CallerID,
		ReceiverID: data.ReceiverID,
		Latitude:   data.Latitude,
		Longitude:  data.Longitude,
	}
}