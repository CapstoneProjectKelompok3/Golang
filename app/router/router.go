package router

import (
	dE "project-capston/features/emergency/data"
	hE "project-capston/features/emergency/handler"
	sE "project-capston/features/emergency/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB,c *echo.Echo){
	dataE:=dE.New(db)
	serviceE:=sE.New(dataE)
	handlerE:=hE.New(serviceE)
	c.POST("users/:receiver_id/emergencies",handlerE.Add)
	c.DELETE("/emergencies/:emergency_id",handlerE.Delete)
}