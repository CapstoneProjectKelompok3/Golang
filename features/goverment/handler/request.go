package handler

type GovermentRequest struct {
	Name      string  `json:"name" form:"name"`
	Type      string  `json:"type" form:"type"`
	Address   string  `json:"address" form:"address"`
	Latitude  float64 `json:"latitude" form:"latitude"`
	Longitude float64 `json:"longitude" form:"longitude"`
}