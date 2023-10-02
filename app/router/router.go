package router

import (
	"project-capston/app/middlewares"
	dE "project-capston/features/emergency/data"
	hE "project-capston/features/emergency/handler"
	sE "project-capston/features/emergency/service"

	_governmentData "project-capston/features/goverment/data"
	_governmentHandler "project-capston/features/goverment/handler"
	_governmentService "project-capston/features/goverment/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	dataE := dE.New(db)
	serviceE := sE.New(dataE)
	handlerE := hE.New(serviceE)
	c.POST("users/:receiver_id/emergencies", handlerE.Add, middlewares.JWTMiddleware())
	c.DELETE("/emergencies/:emergency_id", handlerE.Delete, middlewares.JWTMiddleware())
	c.PUT("/emergencies/:emergency_id", handlerE.Edit, middlewares.JWTMiddleware())
	c.GET("/emergencies/:emergency_id", handlerE.GetById, middlewares.JWTMiddleware())
	c.GET("/emergencies", handlerE.GetAll, middlewares.JWTMiddleware())

	//Government
	governmentData := _governmentData.New(db)
	governmentService := _governmentService.New(governmentData)
	governmentHandlerAPI := _governmentHandler.New(governmentService)

	c.POST("/governments", governmentHandlerAPI.CreateGovernment, middlewares.JWTMiddleware())
}
