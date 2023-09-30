package data

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	GovermentID uint
	Plate string
	Status bool
}
