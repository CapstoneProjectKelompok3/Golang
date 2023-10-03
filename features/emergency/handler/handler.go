package handler

import (
	"fmt"
	"net/http"
	"project-capston/app/middlewares"
	usernodejs "project-capston/features/UserNodeJs"
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
	idClaller,_:=middlewares.ExtractTokenUserId(c)
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
	return c.JSON(http.StatusCreated,"success create data emergency")

}

func (handler *EmergencyHandler)Delete(c echo.Context)error{
	id:=c.Param("emergency_id")
	idConv,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,"id not valid")
	}
	err:=handler.emergencyHandler.Delete(uint(idConv))
	if err!= nil{
		return c.JSON(http.StatusBadRequest,err.Error())
	}
	return c.JSON(http.StatusOK,"success delete emsergemcy")
}

func (handler *EmergencyHandler)Edit(c echo.Context)error{
	id:=c.Param("emergency_id")
	idConv,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,"id not valid")
	}
	var input EmergencyRequest
	errBind:=c.Bind(&input)
	if errBind != nil{
		return c.JSON(http.StatusBadRequest,"error bind data")
	}
	
	Entity:=RequestToEntity(input)
	err:=handler.emergencyHandler.Edit(Entity,uint(idConv))
	if err != nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusOK,"success update emergency")
}

func (handler *EmergencyHandler)GetById(c echo.Context)error{
	
	id:=c.Param("emergency_id")
	idConv,errConv:=strconv.Atoi(id)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest,"id not valid")
	}
	token,errToken:=usernodejs.GetTokenHandler(c)
	if errToken != nil{
		return c.JSON(http.StatusUnauthorized,"fail get token")
	}
	data,err:=handler.emergencyHandler.GetById(uint(idConv),token)
	if err!= nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	response:=EntityToResponse(data)
	return c.JSON(http.StatusOK,map[string]any{
		"message":"success get emergency by id",
		"data": response,
	})	
}

func (handler *EmergencyHandler) GetAll(c echo.Context)error{
	var qparams emergency.QueryParams
	page:= c.QueryParam("page")
	itemsPerPage:=c.QueryParam("itemsPerPage")

	if itemsPerPage ==""{
		qparams.IsClassDashboard=false
	}else{
		qparams.IsClassDashboard=true
		itemsConv, errItem := strconv.Atoi(itemsPerPage)
		if errItem != nil {
			return c.JSON(http.StatusBadRequest,"item per page not valid")
		}
		qparams.ItemsPerPage = itemsConv
	}

	if page ==""{
		qparams.Page=1
	}else{
		pageConv, errPage := strconv.Atoi(page)
		if errPage != nil {
			return c.JSON(http.StatusBadRequest,"page not valid")
		}
		qparams.Page = pageConv
	}
	// name:=c.QueryParam("searchName")
	// qparams.SearchName = name
	token,errToken:=usernodejs.GetTokenHandler(c)
	if errToken != nil{
		return c.JSON(http.StatusUnauthorized,"fail get token")
	}

	bol,data,err:=handler.emergencyHandler.GetAll(qparams,token)
	if err != nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	var response []EmergencyResponse
	for _,v:=range data{
		response = append(response, EntityToResponse(v))
	}

	return c.JSON(http.StatusOK,map[string]any{
		"message":"success get all emergency",
		"data": response,
		"next_page":bol,
	})
}

func (handler *EmergencyHandler)ActionLogic(c echo.Context)error{
	accept:=c.QueryParam("accept")
	fmt.Println("accept",accept)
	err:=handler.emergencyHandler.ActionGmail(accept)
	if err != nil{
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusOK,"action tersimpan")
}