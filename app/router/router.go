package router

import (
	"project-capston/app/middlewares"
	dE "project-capston/features/emergency/data"
	hE "project-capston/features/emergency/handler"
	sE "project-capston/features/emergency/service"

	dV "project-capston/features/vehicles/data"
	hV "project-capston/features/vehicles/handler"
	sV "project-capston/features/vehicles/service"

	_governmentData "project-capston/features/goverment/data"
	_governmentHandler "project-capston/features/goverment/handler"
	_governmentService "project-capston/features/goverment/service"

	_driverData "project-capston/features/driver/data"
	_driverHandler "project-capston/features/driver/handler"
	_driverService "project-capston/features/driver/service"

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

	c.GET("/emergencies/action", handlerE.ActionLogic)

	dataV := dV.New(db)
	serviceV := sV.New(dataV)
	handlerV := hV.New(serviceV)

	c.POST("/vehicles", handlerV.Add, middlewares.JWTMiddleware())
	c.PUT("/vehicles/:vehicle_id", handlerV.Edit, middlewares.JWTMiddleware())
	c.GET("/vehicles/:vehicle_id", handlerV.GetById, middlewares.JWTMiddleware())
	c.GET("/vehicles", handlerV.GetAll, middlewares.JWTMiddleware())
	c.DELETE("/vehicles/:vehicle_id", handlerV.Delete, middlewares.JWTMiddleware())

	//Teguh Government
	governmentData := _governmentData.New(db)
	governmentService := _governmentService.New(governmentData)
	governmentHandlerAPI := _governmentHandler.New(governmentService)

	c.POST("/governments", governmentHandlerAPI.CreateGovernment, middlewares.JWTMiddleware())
	c.GET("/governments", governmentHandlerAPI.GetAllGovernment)
	c.GET("/governments/:government_id", governmentHandlerAPI.GetGovernmentById)
	c.PUT("/governments/:government_id", governmentHandlerAPI.UpdateGovernment, middlewares.JWTMiddleware())
	c.DELETE("/governments/:government_id", governmentHandlerAPI.DeleteGovernment, middlewares.JWTMiddleware())

	c.GET("/get-nearest-government", governmentHandlerAPI.GetNearestGovernment, middlewares.JWTMiddleware())

	//Teguh Government
	driverData := _driverData.New(db)
	driverService := _driverService.New(driverData)
	driverHandlerAPI := _driverHandler.New(driverService)

	c.POST("/drivers", driverHandlerAPI.CreateDriver)
	c.GET("/drivers", driverHandlerAPI.GetAllDriver)
	c.POST("/login-drivers", driverHandlerAPI.Login)
}
