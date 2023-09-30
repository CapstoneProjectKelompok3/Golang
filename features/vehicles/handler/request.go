package handler

type VehicleRequest struct {
	GovermentID uint   `json:"goverment_id" form:"goverment_id"`
	Plate       string `json:"plate" form:"plate"`
	Status      bool   `json:"status" form:"status"`
}