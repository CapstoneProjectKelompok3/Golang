package database

import (
	"fmt"
	"project-capston/app/config"
	driver "project-capston/features/driver/data"
	emergency "project-capston/features/emergency/data"
	government "project-capston/features/goverment/data"
	unit "project-capston/features/unit/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&emergency.Emergency{})
	// db.AutoMigrate(&emergency.Unit{})
	db.AutoMigrate(&government.Government{})
	db.AutoMigrate(&driver.Driver{})
	// db.AutoMigrate(&data.Vehicle{})
	db.AutoMigrate(&unit.Unit{})
	db.AutoMigrate(&unit.UnitHistory{})
	// db.AutoMigrate(&history.History{})
	

	//db.AutoMigrate(&emergency.Emergency{},&driver.Driver{},&goverment.Goverment{},&vehicles.Vehicle{},&unit.Unit{},&unit.UnitHistory{})
}
