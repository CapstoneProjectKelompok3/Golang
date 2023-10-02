package handler

import (
	"net/http"
	"project-capston/features/vehicles"
	"strconv"

	"github.com/labstack/echo/v4"
)

type VehicleHandler struct {
	vehicleHandler vehicles.VehicleServiceInterface
}

func (handler *VehicleHandler)Add(c echo.Context)error{

	var input VehicleRequest
	errBind:=c.Bind(&input)
	if errBind != nil{
		return c.JSON(http.StatusBadRequest,"error bind data")
	}
	entity:=RequestToEntity(input)
	err:=handler.vehicleHandler.Add(entity)
	if err != nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusCreated,"success create vehicle")
}

func (handler *VehicleHandler)Edit(c echo.Context)error{
	id:=c.Param("vehicle_id")
	idConv,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,"id not valid")
	}
	var input VehicleRequest
	errBind:=c.Bind(&input)
	if errBind != nil{
		return c.JSON(http.StatusBadRequest,"id not valid")
	}
	entity:=RequestToEntity(input)
	err:=handler.vehicleHandler.Edit(entity,uint(idConv))
	if err != nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusOK,"success update vehicle")	
}

func (handler *VehicleHandler)GetById(c echo.Context) error{
	id:=c.Param("vehicle_id")
	idConv,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,"id not valid")
	}
	data,err:=handler.vehicleHandler.GetById(uint(idConv))	
	if err != nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	response:=EntityToResponse(data)
	return c.JSON(http.StatusOK,map[string]any{
		"message":"success get vehicle by id",
		"data":response,
	})
}
func (handler *VehicleHandler)GetAll(c echo.Context)error{
	data,err:=handler.vehicleHandler.GetAll()
	if err!=nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	var response []VehicleResponse
	for _,v:=range data{
		response = append(response, EntityToResponse(v))
	}
	return c.JSON(http.StatusOK,map[string]any{
		"message":"success get all vehicle",
		"data":response,
	})
}

func (handler *VehicleHandler)Delete(c echo.Context)error{
	id:=c.Param("vehicle_id")
	idConv,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,"id not valid")
	}
	err:=handler.vehicleHandler.Delete(uint(idConv))
	if err != nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}	
	return c.JSON(http.StatusOK,"success delete vehicle")
}

func New(handler vehicles.VehicleServiceInterface)*VehicleHandler{
	return &VehicleHandler{
		vehicleHandler: handler,
	}
}