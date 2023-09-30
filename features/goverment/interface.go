package goverment

import "time"

type GovermentEntity struct {
	Id        uint
	CreateAt  time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
	Name      string
	Type      string
	Address   string
	Latitude  float64
	Longitude float64
}