package handler

type GovermentResponse struct {
	Id        uint    `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Type      string  `json:"type,omitempty"`
	Address   string  `json:"address,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}