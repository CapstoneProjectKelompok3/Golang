package router

import (
	"project-capston/app/middlewares"
	dE "project-capston/features/emergency/data"
	hE "project-capston/features/emergency/handler"
	sE "project-capston/features/emergency/service"

	dV "project-capston/features/vehicles/data"
	hV "project-capston/features/vehicles/handler"
	sV "project-capston/features/vehicles/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB,c *echo.Echo){
	dataE:=dE.New(db)
	serviceE:=sE.New(dataE)
	handlerE:=hE.New(serviceE)
	c.POST("users/:receiver_id/emergencies",handlerE.Add,middlewares.JWTMiddleware())
	c.DELETE("/emergencies/:emergency_id",handlerE.Delete,middlewares.JWTMiddleware())
	c.PUT("/emergencies/:emergency_id",handlerE.Edit,middlewares.JWTMiddleware())
	c.GET("/emergencies/:emergency_id",handlerE.GetById,middlewares.JWTMiddleware())
	c.GET("/emergencies",handlerE.GetAll,middlewares.JWTMiddleware())

	dataV:=dV.New(db)
	serviceV:=sV.New(dataV)
	handlerV:=hV.New(serviceV)

	c.POST("/vehicles",handlerV.Add,middlewares.JWTMiddleware())
	c.PUT("/vehicles/:vehicle_id",handlerV.Edit,middlewares.JWTMiddleware())
	c.GET("/vehicles/:vehicle_id",handlerV.GetById,middlewares.JWTMiddleware())
	c.GET("/vehicles",handlerV.GetAll,middlewares.JWTMiddleware())
	c.DELETE("/vehicles/:vehicle_id",handlerV.Delete,middlewares.JWTMiddleware())
}