package handler

type GovernmentResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GovernmentNearestResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	// Type      string  `json:"type"`
	// Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Distance  float64 `json:"distance"`
}
