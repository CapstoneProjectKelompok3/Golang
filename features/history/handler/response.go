package handler

import (
	"project-capston/features/history"

	"time"
)

type HistoryResponse struct {
	Id       uint      `json:"id,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
	UnitID   uint      `json:"unit_id,omitempty"`
	DriverID uint      `json:"driver_id,omitempty"`
	Status   string    `json:"status,omitempty"`
}

type UserResponse struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Level string `json:"level,omitempty"`
}

func EntityToResponse(data history.HistoryEntity) HistoryResponse {
	return HistoryResponse{
		Id:       data.Id,
		CreateAt: data.CreateAt,
		UnitID:   data.UnitID,
		DriverID: data.DriverID,
		Status:   data.Status,
	}
}

func UserEntityToResponse(user history.UserEntity) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Level: user.Level,
	}
}
