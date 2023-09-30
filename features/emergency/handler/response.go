package handler

import (
	"project-capston/features/emergency"
	"time"
)

type EmergencyResponse struct {
	Id         uint `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	CreateAt   time.Time `json:"create_at,omitempty"`
	CallerID   uint    `json:"caller_id,omitempty"`
	ReceiverID uint    `json:"receiver_id,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
	Caller     UserResponse `json:"caller,omitempty"`
	Receiver   UserResponse	`json:"reciver,omitempty"`
	IsClose bool
}

type UserResponse struct{
	ID        		int `json:"id,omitempty"`
	Name 			string	`json:"name,omitempty"`
	Level           string `json:"level,omitempty"`
}

func EntityToResponse(data emergency.EmergencyEntity)EmergencyResponse{
	return EmergencyResponse{
		Id:         data.Id,
		CreateAt:   data.CreateAt,
		CallerID:   data.CallerID,
		ReceiverID: data.ReceiverID,
		Latitude:   data.Latitude,
		Longitude:  data.Longitude,
		Caller:     UserEntityToResponse(data.Caller),
		Receiver:   UserEntityToResponse(data.Receiver),
	}
}

func UserEntityToResponse(user emergency.UserEntity)UserResponse{
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Level: user.Level,
	}
}