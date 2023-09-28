package handler

import (
	"project-capston/features/emergency"
	"time"
)

type EmergencyResponse struct {
	Id         uint `json:"id,omitempty"`
	CreateAt   time.Time `json:"create_at,omitempty"`
	CallerID   uint    `json:"caller_id,omitempty"`
	ReceiverID uint    `json:"receiver_id,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
}

func EntityToResponse(data emergency.EmergencyEntity)EmergencyResponse{
	return EmergencyResponse{
		Id:         data.Id,
		CreateAt:   data.CreateAt,
		CallerID:   data.CallerID,
		ReceiverID: data.ReceiverID,
		Latitude:   data.Latitude,
		Longitude:  data.Longitude,
	}
}