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

	dU "project-capston/features/unit/data"
	hU "project-capston/features/unit/handler"
	sU "project-capston/features/unit/service"

	dH "project-capston/features/history/data"
	hH "project-capston/features/history/handler"
	sH "project-capston/features/history/service"

	"github.com/go-redis/redis/v8"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo, redis *redis.Client) {
	dataE := dE.New(db, redis)
	serviceE := sE.New(dataE)
	handlerE := hE.New(serviceE)
	c.POST("/users/:receiver_id/emergencies", handlerE.Add, middlewares.JWTMiddleware())
	c.DELETE("/emergencies/:emergency_id", handlerE.Delete, middlewares.JWTMiddleware())
	c.PUT("/emergencies/:emergency_id", handlerE.Edit, middlewares.JWTMiddleware())
	c.GET("/emergencies/:emergency_id", handlerE.GetById, middlewares.JWTMiddleware())
	c.GET("/emergencies", handlerE.GetAll, middlewares.JWTMiddleware())

	c.GET("/emergencies/action", handlerE.ActionLogic)
	c.GET("/emergencies/count", handlerE.CountEmergency)

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
	c.GET("/governments", governmentHandlerAPI.GetAllGovernment, middlewares.JWTMiddleware())
	c.GET("/governments/:government_id", governmentHandlerAPI.GetGovernmentById, middlewares.JWTMiddleware())
	c.PUT("/governments/:government_id", governmentHandlerAPI.UpdateGovernment, middlewares.JWTMiddleware())
	c.DELETE("/governments/:government_id", governmentHandlerAPI.DeleteGovernment, middlewares.JWTMiddleware())

	c.GET("/get-nearest-government", governmentHandlerAPI.GetNearestGovernment, middlewares.JWTMiddleware())

	c.GET("/governments/count", governmentHandlerAPI.CountUnit, middlewares.JWTMiddleware())

	//Teguh Government
	driverData := _driverData.New(db)
	driverService := _driverService.New(driverData)
	driverHandlerAPI := _driverHandler.New(driverService)

	c.POST("/drivers", driverHandlerAPI.CreateDriver)
	c.GET("/drivers", driverHandlerAPI.GetAllDriver)
	c.POST("/drivers/login", driverHandlerAPI.Login)
	c.GET("/drivers/assign", driverHandlerAPI.KerahkanDriver)
	c.GET("/driver/profile", driverHandlerAPI.GetProfileDriver, middlewares.JWTMiddleware())
	c.GET("/driver/confirm", driverHandlerAPI.DriverAcceptOrRejectOrder, middlewares.JWTMiddleware())
	c.PUT("/driver/ontrip", driverHandlerAPI.DriverOnTrip, middlewares.JWTMiddleware())
	c.GET("/drivers/count", driverHandlerAPI.GetCountDriver)
	c.DELETE("/drivers/:driver_id",driverHandlerAPI.Delete)

	dataU := dU.New(db)
	serviceU := sU.New(dataU)
	handlerU := hU.New(serviceU)
	c.POST("units", handlerU.Add, middlewares.JWTMiddleware())
	c.DELETE("/units/:unit_id", handlerU.Delete, middlewares.JWTMiddleware())
	c.PUT("/units/:unit_id", handlerU.Edit, middlewares.JWTMiddleware())
	c.GET("/units/:unit_id", handlerU.GetById, middlewares.JWTMiddleware())
	c.GET("/units", handlerU.GetAll, middlewares.JWTMiddleware())

	dataH := dH.New(db)
	serviceH := sH.New(dataH)
	handlerH := hH.New(serviceH)
	c.POST("/histories", handlerH.Add, middlewares.JWTMiddleware())
	c.DELETE("/histories/:history_id", handlerH.Delete, middlewares.JWTMiddleware())
	c.PUT("/histories/:history_id", handlerH.Edit, middlewares.JWTMiddleware())
	c.GET("/histories/:history_id", handlerH.GetById, middlewares.JWTMiddleware())
	c.GET("/histories", handlerH.GetAll, middlewares.JWTMiddleware())

}
