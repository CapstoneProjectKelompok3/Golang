package handler

type VehicleResponse struct {
	Id          uint   `json:"id,omitempty"`
	GovermentID uint   `json:"goverment_id,omitempty"`
	Plate       string `json:"plate,omitempty"`
	Status      bool   `json:"status,omitempty"`
}