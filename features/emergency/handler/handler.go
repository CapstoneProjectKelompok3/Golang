package handler

import (
	"net/http"
	"project-capston/features/emergency"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type EmergencyHandler struct {
	emergencyHandler emergency.EmergencyServiceInterface
}

func New(handler emergency.EmergencyServiceInterface) *EmergencyHandler{
	return &EmergencyHandler{
		emergencyHandler: handler,
	}
}

func (handler *EmergencyHandler) Add(c echo.Context)error{
	idClaller:=5
	idReceiver:=c.Param("receiver_id")
	idConv,errConv:=strconv.Atoi(idReceiver)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,"id not valid")
	}
	var input EmergencyRequest
	errBind:=c.Bind(&input)
	if errBind != nil{
		return c.JSON(http.StatusBadRequest,"error bind data")
	}
	entity:=RequestToEntity(input)
	entity.CallerID=uint(idClaller)
	entity.ReceiverID=uint(idConv)

	err:=handler.emergencyHandler.Add(entity)
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return c.JSON(http.StatusBadRequest,err.Error())
		}else{
			return c.JSON(http.StatusInternalServerError,err.Error())
		}
	}
	return c.JSON(http.StatusOK,"success create data emergency")

}