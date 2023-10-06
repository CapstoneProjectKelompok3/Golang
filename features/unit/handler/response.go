package handler

import (
	"project-capston/features/unit"
	"time"
)

type UnitResponse struct {
	Id            uint         `json:"id,omitempty"`
	CreateAt      time.Time    `json:"create_at,omitempty"`
	EmergenciesID uint         `json:"emergencies_id,omitempty"`
	GovermentType string       `json:"goverment_type,omitempty"`
	SumOfUnit     int          `json:"sum_of_unit,omitempty"`
	Emergencies   UserResponse `json:"emergencies,omitempty"`
}

type UnitHistoryResponse struct {
	EmergenciesID   uint   `json:"emergencies_id,omitempty"`
	VehicleID       uint   `json:"vehicle_id,omitempty"`
	Status          string `json:"status,omitempty"`
	AlasanPenolakan string `json:"alasan_penolakan,omitempty"`
}

type UserResponse struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Level string `json:"level,omitempty"`
}

func EntityToResponse(data unit.UnitEntity) UnitResponse {
	return UnitResponse{
		Id:            data.Id,
		CreateAt:      data.CreateAt,
		EmergenciesID: data.EmergenciesID,
		GovermentType: data.GovermentType,
		SumOfUnit:     data.SumOfUnit,
		Emergencies:   UserEntityToResponse(data.Emergencies),
	}
}

func UserEntityToResponse(user unit.UserEntity) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Level: user.Level,
	}
}
