package data

import "gorm.io/gorm"

type Goverment struct {
	gorm.Model
	Name string
	Type string
	Address string
	Latitude   float64
	Longitude  float64
}