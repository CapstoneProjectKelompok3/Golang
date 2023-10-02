package handler

import (
	"project-capston/features/goverment"
	"time"
)

type GovernmentRequest struct {
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Address   string    `json:"address"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func RequestToCore(input GovernmentRequest) goverment.Core {
	return goverment.Core{
		Name:      input.Name,
		Type:      input.Type,
		Address:   input.Address,
		Latitude:  input.Latitude,
		Longitude: input.Longitude,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: time.Time{},
	}
}
